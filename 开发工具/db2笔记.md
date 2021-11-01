当命令行处理器完成处理命令或SQL语句时，它返回一个返回（或退出）代码。这些代码对于从命令行执行CLP函数的用户是透明的，但是当从shell脚本执行这些函数时，可以检索它们。
CLP:DB2 命令行处理器
例如，以下shell脚本执行
~~~
GET DATABASE MANAGER CONFIGURATION
~~~
命令，然后检查CLP返回码：

返回码描述如下

- 0       DB2命令或SQL语句已成功执行
- 1       SELECT或FETCH语句没有返回任何行
- 2       DB2命令或SQL语句警告 
- 4       DB2命令或SQL语句错误 
- 8       命令行处理器系统错误