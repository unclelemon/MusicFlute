# 变量
## 查看包变量
go env

## 修改变量
go env -w GOPATH=/Users/linhaizeng/LhzDocuments/hz-project/learn_go

###### GO111MODULE=“off”
需要将写的代码放在$GOPATH/src
###### GO111MODULE=“on”

需要创建go.mod，内容如下
~~~
go mod init test #其中test可以是任意的名称
~~~
之后需要导入包
~~~
import "test/module/moduleA"
~~~

使用 go env 查看go的环境变量

go path 显示go的路径，一般go的模块会安装到该路径下

go111module 是否使用模块支持的变量，如果设为off ，代表无模块支持，import的包会从gopath下寻找。

如果设为on，代表模块支持，会忽略gopath，在go.mod中寻找依赖。

所以如果go111module = 'off',要将项目放在gopath的路径下，并使用go get 安装需要的第三方模块

如果 go111module = 'on' ，可以go mod init 初始化go.mod文件，再使用go build，会自动下载需要的第三方模块。

ps: 配置go环境变量的方法

go env -w GO111MODULE="on"
ps again: go path 与 go mod方式不能共存，如果开启 go111module = 'on'，需要删除project下设置的gopath,（global gopath保留）

### 内存


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

###### 保持package尽量和目录名保持一致

###### 变量名 函数名 常量名 尽量使用驼峰法

###### 首字母大写，则可以被其他的包访问，首字母小写只能在本包中使用
可以简单理解成为首字母大写是public ，而首字母小写是private，在golang中没有public和private等关键字

# 交叉编译
~~~
# mac 编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build client.go
# Mac 编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build filename.go
~~~
# 解决从github下载过慢
~~~
git config --global --unset http.proxy

git config --global --unset https.proxy

~~~

# 查看go版本

~~~
go version
~~~
# 创建可执行文件
需要在项目根目录下，src文件夹外层。编译后可生成库文件 .a
~~~
go build -o bin\my.exe ****\main 
~~~

