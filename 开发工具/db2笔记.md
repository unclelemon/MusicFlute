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

## 表空间
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

## 创建表、新增字段、插入数据 失败回滚
~~~
# 也可以从输入参数，使用$* 获取全部输入参数
tables=('table1' 'table2' 'table3')

db2 connect to schema
# -v 打印执行的每一条sql view
# -s 参数s的意思是每次遇到报错的命令就停止执行操作，后面的sql就不再执行了 stop
# -t 表示使用分号语句终结符 termin
# -f 表示其后就是读取sql的脚本文件 从输入文件读内容 file
db2 -stvf bak.txt -z table_ddl_1.log
if [[ $? -ne 0 ]];then
    echo "error"
    db2 rollbak
    for arg in ${tables[@]}
    do
        # 设置单独一列但临时空间，该列值都为1
        sql="select 1 from ${arg} fetch first 1 rows only"
        STATE=`db2 ${sql}`
        # =~ 表示${STATE} 字符串中是否包含字符串 “SQLSTATE”
        if [[ ${STATE} =~ "SQLSTATE" ]]
        then
            echo "${arg} NOT EXISTS"
        else 
            db2 "drop table ${arg}"
            echo "drop ${arg} success"
        fi
    done
else 
    db2 commit
fi
db2 disconnect current
# 或者db2 connect  reset
~~~
### decimal()
知道decimal(p[ , s])的意思就能明白了，decimal(10)中默认的s为0，就是说没有小数位，提取10位的整数，舍去小数后的小数位；decimal(10,2)的意思是提取10位的浮点数（包括整数位和小数位，一共10位），取小数点后两位，四舍五入。在DB2中decimal最大精度是31位，小数的范围从-10^31+1到10^31-1。