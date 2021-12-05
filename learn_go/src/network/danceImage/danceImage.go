package danceImage

import (
	"fmt"
	"os"
	"bytes"
	"github.com/nfnt/resize"
  "image"
  "image/png"
  "io/ioutil"
  "log"

)

/*先定义一个函数
参数：
  imgPath: 图片路径
  size: 生成文本后的尺寸(这个不是真实的尺寸，1代表1个像素，1个像素会被替换成1个字符，所以是字符的个数，高度是自动换算的，所以这里的size指的是“宽度”被压缩成多少像素)
  txts: 将像素处理成的字符列表
  rowend: 换行字符（因为windows和linux不同）
  output: 生成文本文件保存路径
*/
func img2txt(imgPath string, size uint, txts []string, rowend string, output string) {
	//获取图片文件 
	file, err := os.Open(imgPath)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
	defer file.Close()
	
	//用图片文件获取图片对象
	img, err := png.Decode(file)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
  
	//用将宽度设置为size，然后换算出等比例的高度
	var width = size
	var height = (size * uint(img.Bounds().Dy())) / (uint(img.Bounds().Dx()))
	height = height * 6 / 10 //这里6/10是大致字符的宽高比
	newimg := resize.Resize(width, height, img, resize.Lanczos3)  //根据高宽resize图片，并得到新图片的像素值
	dx := newimg.Bounds().Dx()
	dy := newimg.Bounds().Dy()
  
	//创建一个字节buffer，一会用来保存字符
	textBuffer := bytes.Buffer{}
  
	//遍历图片每一行每一列像素
	for y := 0; y < dy; y++ {
	  for x := 0; x < dx; x++ {
		colorRgb := newimg.At(x, y)
		r, g, b, _ := colorRgb.RGBA()
  
		//获得三原色的值，算一个平均数出来
		avg := uint8((r + g + b) / 3 >> 8)
		//有多少个用来替换的字符就将256分为多少个等分，然后计算这个像素的平均值趋紧与哪个字符，最后，将这个字符添加到字符buffer里
		num := avg / uint8(256/len(txts))
		textBuffer.WriteString(txts[num])
		fmt.Print(txts[num]) //打印出来
	  }
  
	  textBuffer.WriteString(rowend)  //一行结束，换行
	  fmt.Print(rowend)
	}
  
	//将字符buffer的数据写入到文本文件里，结束。
	f, err := os.Create(output + ".txt")
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
	defer f.Close()
  
	f.WriteString(textBuffer.String())
  }

// GetImage 根据路径获取一个图片
func GetImage(imagePath string) *image.Image {
    // 读取image文件
    imageFile, _ := ioutil.ReadFile(imagePath)
    // bytes.buffer 是一个缓冲byte类型的缓冲器存放着都是byte
    imageBytes := bytes.NewBuffer(imageFile)
    // 解码
    im, _, err := image.Decode(imageBytes)
    if err != nil {
        log.Printf("解码图像失败%#v", err)
        return nil
    }
    return &im
}

// equalScaleImageFromWidth 根据宽度，等比例缩放图片, 返回新的像素值（等比例处理后的）
// maxWidth 图片最大宽度 如果超过此值就按照宽度 targetWidth 等比例缩放图片
func equalScaleImageFromWidth(img image.Image, maxWidth int, targetWidth int) *image.Image {
    // 原图像界限范围
    bounds := img.Bounds()
    dx := bounds.Dx()
    dy := bounds.Dy()
    log.Printf("新用户新来了，展示图宽=%d,高=%d", dx, dy)
    m := img
    if dx > maxWidth {
        // 按照宽度缩放
        m = resize.Resize(uint(targetWidth), 0, img, resize.Lanczos3)
    }
    return &m
}


// ImageToChars 图像转成字符展示
// RGBA(R,G,B,A) 像素值
// R：红色值 G：绿色值 B：蓝色值 A：Alpha透明度
func ImageToChars(img *image.Image, inputChars string) string {
    sourceImage := *img
    bounds := sourceImage.Bounds()
    dx := bounds.Dx()
    dy := bounds.Dy()
    imgChars := "**  "
    if inputChars != "" {
        imgChars = inputChars
    }
    resultString := ""
    intSliceRGB := []int{}
    maxIntRGB := 0
    minIntRGB := 255 * 3
    for i := 0; i < dy; i++ {
        for j := 0; j < dx; j++ {
            colorRgb := sourceImage.At(j, i)
            // 获取 uint32 像素值 >> 8 转换为 255 值
            r, g, b, _ := colorRgb.RGBA()
            // sumRGB 越大越趋近于白色，越小越趋近于黑色
            sumRGB := int(uint8(r>>8)) + int(uint8(g>>8)) + int(uint8(b>>8))
            // 找到最大值和最小值，方便将像素值划分不同区间段
            if maxIntRGB < sumRGB {
                maxIntRGB = sumRGB
            }
            if minIntRGB > sumRGB {
                minIntRGB = sumRGB
            }
            intSliceRGB = append(intSliceRGB, sumRGB)
        }
    }
    for index, val := range intSliceRGB {
        // partLen为区间跨度 我们按照传入字符串元素个数平均分配像素值，将像素值分成几个区间
        //  +1 防止下标溢出
        // 例如 像素值 0~500 用两个字符替换 0~250 251~500 两个区间分别与其对应即可
        partLen := (maxIntRGB-minIntRGB)/(len(imgChars)) + 1
        // 根据像素值取不同的字符替换像素
        str := string(imgChars[(val-minIntRGB)/partLen])
        resultString += str
        // 判断换行
        if (index+1)%dx == 0 {
            resultString += "\n"
        }
    }
    return resultString
}

// func main()  {
// 	// imgPath := "/Users/linhaizeng/LhzDocuments/hz-project/learn_go/src/network/charImage/lsy.png"
// 	// var size uint = 70
// 	// txts := []string{"I","L","U","lsy","0","3","2","8","text","sd","*","2","hzl"}
// 	// rowend := "\r\n"
// 	// output := "/Users/linhaizeng/LhzDocuments/hz-project/learn_go/src/network/charImage/lsy_ouput"
// 	// img2txt(imgPath,size,txts,rowend,output)
// 	str := ImageToChars(equalScaleImageFromWidth(*(GetImage("/Users/linhaizeng/LhzDocuments/hz-project/learn_go/src/network/charImage/2.png")), 150, 100), "#$*@123+-.?,    ")
// 	fmt.Println(str)
// }

func AddHead() string {
  return ImageToChars(equalScaleImageFromWidth(*(GetImage("/Users/linhaizeng/LhzDocuments/hz-project/learn_go/src/network/charImage/lsy.png")), 150, 150), "#$*@123+-.?,    ")
}