# Redis学习

[Redisdoc.com](redisdoc.com)

## Docker安装Redis

~~~
# 在虚拟机中
mkdir -p /mydata/redis/conf
touch /mydata/redis/conf/redis.conf

docker pull redis

docker run -p 6379:6379 --name redis \
-v /mydata/redis/data:/data \
-v /mydata/redis/conf/redis.conf:/etc/redis/redis.conf \
-d redis redis-server /etc/redis/redis.conf

# 直接进去redis客户端。
docker exec -it redis redis-cli
~~~

默认是不持久化的。在配置文件中输入appendonly yes，就可以aof持久化了。修改完docker restart redis，docker -it redis redis-cli

~~~
vim /mydata/redis/conf/redis.conf
# 插入下面内容
appendonly yes
# 保存

docker restart redis
~~~

设置redis容器在docker启动的时候启动

~~~
docker update redis --restart=always
~~~

## Redis五大数据类型：

String （字符串） Hash（哈希） List(列表) Set（集合） zset(有序集合)

默认安装16个数据库，编号0~15

~~~redis
# 切换数据库1号数据
select 1

# 查看当前的key数量
dbsize

# 清空当前数据库
flushdb

# 清空所有数据库
flushall
~~~

## Redis数据类型和CRUD

### String

~~~
# 添加 和 修改
set key1 hello

# 获取数据
get key1

# 删除数据
del key1

#设置生存时间setex(set with expire)
# 设置mess01 变量为hello,you 生存10秒
setex mess01 10 hello,you

# mset 同时设置一个或多个value
mset worker01 tom worker02 scott

# mget worker01  worker02 
~~~

### Hash

~~~
# 添加
hset user1 age 30
hset user1 name "smith"

hget user1 age
hget user1 name
 
hdel user1
~~~

### List

~~~
lpush/ rpush /lrange /lpop /rpop /del
lpush city beijing shanghai tianjing

lrange city 0 -1
~~~

## Go操作redis

~~~go
package main
import （
	"fmt"
	"github.com/garybuard/redigo/redis"
）

func main() {
  // 通过go 向redis 写入输入
  conn,err := redis.Dial("tcp", "localhost:6379")
  if err != nil {
    fmt.Println("connect to redis error:",err)
    return
  }
  defer conn.Close()
}
~~~







