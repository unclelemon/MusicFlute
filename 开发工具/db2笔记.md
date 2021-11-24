当命令行处理器完成处理命令或SQL语句时，它返回一个返回（或退出）代码。这些代码对于从命令行执行CLP函数的用户是透明的，但是当从shell脚本执行这些函数时，可以检索它们。
CLP:DB2 命令行处理器
例如，以下shell脚本执行
~~~
GET DATABASE MANAGER CONFIGURATION
~~~
命令，然后检查CLP返回码：

返回码描述如下

- 0 ：      DB2命令或SQL语句已成功执行
- 1 ：      SELECT或FETCH语句没有返回任何行
- 2 ：      DB2命令或SQL语句警告 
- 4 ：      DB2命令或SQL语句错误 
- 8 ：      命令行处理器系统错误

# 表空间
1. 查看表所在的表空间：
~~~shell
–-表名大写
select tabname,tbspace from syscat.tables where tabname = ‘表名’;
~~~
2. 查看表的索引：
~~~sql
–-表名大写
select * from syscat.indexes where tabname = ‘表名’;
~~~

## 迁移数据库
- db2move schema export
- db2move schema import

select 1 from 中的1是一常量（可以为任意数值），查到的所有行的值都是它，但从效率上来说，1>xxx>*，因为不用查字典表。

1：select  1 from table       增加临时列，每行的列值是写在select后的数，这条sql语句中是1

2：select count(1)  from table   管count(a)的a值如何变化，得出的值总是table表的行数

3：select sum(1) from table   计算临时列的和