# 变量
## 查看包变量
go env

## 修改变量
go env -w GOPATH=/Users/linhaizeng/LhzDocuments/hz-project/learn_go

### GO111MODULE=“off”
需要将写的代码放在$GOPATH/src
### GO111MODULE=“on”
需要创建go.mod，内容如下
~~~
go mod init test #其中test可以是任意的名称
~~~
之后需要导入包
~~~
import "test/module/moduleA"
~~~


- 值类型：int float bool string 数组 结构体stuct
- 引用类型： 指针 slice切片 map 管道chan interface 等都是引用类型

值类型：栈内存，通常是在栈区
引用类型：变量存储的是一个地址，内存通常在堆区分配。当没有任何变量引用这个地址时候，该地址对应的数据空间就变成一个垃圾，有GC来回收

# 标识符的命令
1. 有26个英文祖母大小写，0——9， _组成
2. 数字不可以开头
3. 严格区分大小写
4. 标识符不能包含空格
5. 下划线本身在golang中本身就是一个特殊的标识符，成为特殊标识符，可以代表任何其他的标识符，但是它对应的值会被忽略，所以仅能作为占位符使用，不能作为标识符使用
6. 不能使用保留关键字（25个）作为标识符
int float32 不算关键字，但是也不建议用作标识符

# 保持package尽量和目录名保持一致

# 变量名 函数名 常量名 尽量使用驼峰法

# 首字母大写，则可以被其他的包访问，首字母小写只能在本包中使用
可以简单理解成为首字母大写是public ，而首字母小写是private，在golang中没有public和private等关键字

# 交叉编译
~~~
# mac 编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build client.go
~~~
# 解决从github下载过慢
~~~
git config --global --unset http.proxy

git config --global --unset https.proxy

~~~