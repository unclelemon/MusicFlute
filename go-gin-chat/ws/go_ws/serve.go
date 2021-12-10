package go_ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-gin-chat/models"
	"go-gin-chat/services/helper"
	"go-gin-chat/ws"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)


// 客户端连接详情
type wsClients struct {
	Conn *websocket.Conn `json:"conn"`

	RemoteAddr string `json:"remote_addr"`

	Uid float64 `json:"uid"`

	Username string `json:"username"`

	RoomId string `json:"room_id"`

	AvatarId string `json:"avatar_id"`
}

// client & serve 的消息体
type msg struct {
	Status int             `json:"status"`
	Data   interface{}     `json:"data"`
	Conn   *websocket.Conn `json:"conn"`
}

// 变量定义初始化
var (
	wsUpgrader = websocket.Upgrader{}

	clientMsg = msg{}

	mutex = sync.Mutex{}

	//rooms = [roomCount + 1][]wsClients{}
	rooms = make(map[int][]wsClients)

	enterRooms = make(chan wsClients)

	sMsg = make(chan msg)

	offline = make(chan *websocket.Conn)

	chNotify = make(chan int ,1)
)

// 定义消息类型
const msgTypeOnline = 1        // 上线
const msgTypeOffline = 2       // 离线
const msgTypeSend = 3          // 消息发送
const msgTypeGetOnlineUser = 4 // 获取用户列表
const msgTypePrivateChat = 5   // 私聊

const roomCount = 6 // 房间总数

type GoServe struct {
	ws.ServeInterface
}

func (goServe *GoServe) RunWs(gin *gin.Context)  {
	// 使用 channel goroutine
	Run(gin)
}

func (goServe *GoServe) GetOnlineUserCount() int {
	return GetOnlineUserCount()
}

func (goServe *GoServe) GetOnlineRoomUserCount(roomId int) int  {
	return GetOnlineRoomUserCount(roomId)
}


func Run(gin *gin.Context) {

	// @see https://github.com/gorilla/websocket/issues/523
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, _ := wsUpgrader.Upgrade(gin.Writer, gin.Request, nil)

	defer c.Close()

	go read(c)
	go write()

	select {}

}

func read(c *websocket.Conn) {

	defer func() {
		//捕获read抛出的panic
		if err := recover();err!=nil{
			log.Println("read发生错误",err)
		}
	}()

	for {
		_, message, err := c.ReadMessage()
		// log.Println("client message", string(message),c.RemoteAddr())
		if err != nil { // 离线通知
			offline <- c
			log.Println("ReadMessage error1", err)
			return
		}

		serveMsgStr := message

		// 处理心跳响应 , heartbeat为与客户端约定的值
		if string(serveMsgStr) == `heartbeat` {
			c.WriteMessage(websocket.TextMessage, []byte(`{"status":0,"data":"heartbeat ok"}`))
			continue
		}

		json.Unmarshal(message, &clientMsg)
		// log.Println("来自客户端的消息", clientMsg,c.RemoteAddr())
		if clientMsg.Data != nil {
			if clientMsg.Status == msgTypeOnline { // 进入房间，建立连接
				roomId, _ := getRoomId()

				enterRooms <- wsClients{
					Conn:       c,
					RemoteAddr: c.RemoteAddr().String(),
					Uid:        clientMsg.Data.(map[string]interface{})["uid"].(float64),
					Username:   clientMsg.Data.(map[string]interface{})["username"].(string),
					RoomId:     roomId,
					AvatarId:   clientMsg.Data.(map[string]interface{})["avatar_id"].(string),
				}
			}

			_, serveMsg := formatServeMsgStr(clientMsg.Status, c)
			sMsg <- serveMsg
		}
	}
}

func write() {

	defer func() {
		//捕获write抛出的panic
		if err := recover();err!=nil{
			log.Println("write发生错误",err)
		}
	}()

	for {
		select {
		case r := <-enterRooms:
			handleConnClients(r.Conn)
		case cl := <-sMsg:
			serveMsgStr, _ := json.Marshal(cl)
			switch cl.Status {
			case msgTypeOnline, msgTypeSend:
				notify(cl.Conn, string(serveMsgStr))
			case msgTypeGetOnlineUser:
				chNotify <- 1
				cl.Conn.WriteMessage(websocket.TextMessage, serveMsgStr)
				<-chNotify
			case msgTypePrivateChat:
				chNotify <- 1
				toC := findToUserCoonClient()
				if toC != nil {
					toC.(wsClients).Conn.WriteMessage(websocket.TextMessage, serveMsgStr)
				}
				<-chNotify
			}
		case o := <-offline:
			disconnect(o)
		}
	}
}

func handleConnClients(c *websocket.Conn) {
	roomId, roomIdInt := getRoomId()
	assignRoom := rooms[roomIdInt]
	for cKey, wcl := range assignRoom {
		if wcl.Uid == clientMsg.Data.(map[string]interface{})["uid"].(float64) {
			// 通知当前用户下线
			wcl.Conn.WriteMessage(websocket.TextMessage, []byte(`{"status":-1,"data":[]}`))
			mutex.Lock()
			rooms[roomIdInt] = append(rooms[roomIdInt][:cKey], rooms[roomIdInt][cKey+1:]...)
			mutex.Unlock()
			wcl.Conn.Close()
		}
	}

	mutex.Lock()
	rooms[roomIdInt] = append(rooms[roomIdInt], wsClients{
		Conn:       c,
		RemoteAddr: c.RemoteAddr().String(),
		Uid:        clientMsg.Data.(map[string]interface{})["uid"].(float64),
		Username:   clientMsg.Data.(map[string]interface{})["username"].(string),
		RoomId:     roomId,
		AvatarId:   clientMsg.Data.(map[string]interface{})["avatar_id"].(string),
	})
	mutex.Unlock()
}

// 获取私聊的用户连接
func findToUserCoonClient() interface{} {
	_, roomIdInt := getRoomId()

	toUserUid := clientMsg.Data.(map[string]interface{})["to_uid"].(string)
	assignRoom := rooms[roomIdInt]
	for _, c := range assignRoom {
		stringUid := strconv.FormatFloat(c.Uid, 'f', -1, 64)
		if stringUid == toUserUid {
			return c
		}
	}

	return nil
}

// 统一消息发放
func notify(conn *websocket.Conn, msg string) {
	chNotify <- 1 // 利用channel阻塞 避免并发去对同一个连接发送消息出现panic: concurrent write to websocket connection这样的异常
	_, roomIdInt := getRoomId()
	assignRoom := rooms[roomIdInt]
	for _, con := range assignRoom {
		if con.RemoteAddr != conn.RemoteAddr().String() {
			con.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
	<-chNotify
}

// 离线通知
func disconnect(conn *websocket.Conn) {
	_, roomIdInt := getRoomId()
	assignRoom := rooms[roomIdInt]
	for index, con := range assignRoom {
		if con.RemoteAddr == conn.RemoteAddr().String() {
			data := map[string]interface{}{
				"username": con.Username,
				"uid":      con.Uid,
				"time":     time.Now().UnixNano() / 1e6, // 13位  10位 => now.Unix()
			}

			jsonStrServeMsg := msg{
				Status: msgTypeOffline,
				Data:   data,
			}
			serveMsgStr, _ := json.Marshal(jsonStrServeMsg)

			disMsg := string(serveMsgStr)

			mutex.Lock()
			rooms[roomIdInt] = append(rooms[roomIdInt][:index], rooms[roomIdInt][index+1:]...)
			mutex.Unlock()
			con.Conn.Close()
			notify(conn, disMsg)
		}
	}
}

// 格式化传送给客户端的消息数据
func formatServeMsgStr(status int, conn *websocket.Conn) ([]byte, msg) {

	roomId, roomIdInt := getRoomId()

	data := map[string]interface{}{
		"username": clientMsg.Data.(map[string]interface{})["username"].(string),
		"uid":      clientMsg.Data.(map[string]interface{})["uid"].(float64),
		"room_id":  roomId,
		"time":     time.Now().UnixNano() / 1e6, // 13位  10位 => now.Unix()
	}

	if status == msgTypeSend || status == msgTypePrivateChat {
		data["avatar_id"] = clientMsg.Data.(map[string]interface{})["avatar_id"].(string)
		content := clientMsg.Data.(map[string]interface{})["content"].(string)

		data["content"] = content
		if helper.MbStrLen(content) > 800 {
			// 直接截断
			data["content"] = string([]rune(content)[:800])
		}

		toUidStr := clientMsg.Data.(map[string]interface{})["to_uid"].(string)
		toUid, _ := strconv.Atoi(toUidStr)

		// 保存消息
		stringUid := strconv.FormatFloat(data["uid"].(float64), 'f', -1, 64)
		intUid, _ := strconv.Atoi(stringUid)

		if _, ok := clientMsg.Data.(map[string]interface{})["image_url"]; ok {
			// 存在图片
			models.SaveContent(map[string]interface{}{
				"user_id":    intUid,
				"to_user_id": toUid,
				"content":    data["content"],
				"room_id":    data["room_id"],
				"image_url":  clientMsg.Data.(map[string]interface{})["image_url"].(string),
			})
		} else {
			models.SaveContent(map[string]interface{}{
				"user_id":    intUid,
				"to_user_id": toUid,
				"room_id":    data["room_id"],
				"content":    data["content"],
			})
		}

	}

	if status == msgTypeGetOnlineUser {
		ro := rooms[roomIdInt]
		data["count"] = len(ro)
		data["list"] = ro
	}

	jsonStrServeMsg := msg{
		Status: status,
		Data:   data,
		Conn:   conn,
	}
	serveMsgStr, _ := json.Marshal(jsonStrServeMsg)

	return serveMsgStr, jsonStrServeMsg
}

func getRoomId() (string, int) {
	roomId := clientMsg.Data.(map[string]interface{})["room_id"].(string)

	roomIdInt, _ := strconv.Atoi(roomId)
	return roomId, roomIdInt
}

// =======================对外方法=====================================

func GetOnlineUserCount() int {
	num := 0
	for i := 1; i <= roomCount; i++ {
		num = num + GetOnlineRoomUserCount(i)
	}
	return num
}

func GetOnlineRoomUserCount(roomId int) int {
	return len(rooms[roomId])
}
