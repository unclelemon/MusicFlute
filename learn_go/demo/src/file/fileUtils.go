package file

import (
	"bufio"
    "log"
    "io"
    "fmt"
    "os"
)


func CreateFile()  {
	newFile,err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func CreateFileByPath(filePath string)  {
	newFile,err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func ReadFile(filePath string)  {
	file,err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败")
		log.Fatal(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break;
		}
		if err != nil {
			fmt.Println("文件打开失败")
		}
		fmt.Printf("%s",str)
	}
}