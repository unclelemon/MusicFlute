package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"file"
)

var pathSeparator = string(os.PathSeparator)

/**
 * 文件重命名
 */
func rename(path string, old string, new string) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			err = rename(path+pathSeparator+fileInfo.Name(), old, new)
			if err != nil {
				return err
			}
			err = os.Rename(path+pathSeparator+fileInfo.Name(), path+pathSeparator+strings.Replace(fileInfo.Name(), old, new, -1))
			if err != nil {
				return err
			}
		} else {
			err = os.Rename(path+pathSeparator+fileInfo.Name(), path+pathSeparator+strings.Replace(fileInfo.Name(), old, new, -1))
			if err != nil {
				return err
			}
		}
	}
	return err
}

func renameTest()  {
	//1.获取要被重命名目录(文件)的绝对路径
	fmt.Print("请输入文件的绝对路径：")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.Replace(filePath, "\n", "", -1)

	//2.获取要被替换掉的名称
	fmt.Print("请输入要被替换为空的名称：")
	reader = bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	//3.递归调用重命名
	err := rename(filePath, name, "")
	if err != nil {
		log.Fatalf("发生错误，错误为：%v\n", err)
	}
	err = os.Rename(filePath, strings.Replace(filePath, name, "", -1))
	if err != nil {
		log.Fatalf("发生错误，错误为：%v\n", err)
	}

	fmt.Println("success")
}

func main() {

	// file.ReadFile("/Users/linhaizeng/LhzDocuments/hz-project/learn_go/demo/src/main/test.txt")
	// file.CreateFile()
	file.CreateFileByPath("/Users/linhaizeng/LhzDocuments/hz-project/learn_go/demo/src/main/test1.txt")
}
