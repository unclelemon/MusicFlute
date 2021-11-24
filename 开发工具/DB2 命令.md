 # DB2 学习指南

## 常用开发命令

### 连接数据库

~~~sql
db2 connect to iaca42db
~~~

### 执行sql 脚本

~~~sql
db2 -tvf *.sql，此命令执行*.sql脚本中间出现错误不断开；

db2 -stvf *.sql，此命令执行*.sql脚本中间出现错误会断开，并提示错误；

db2 -stvf *.sql -z script.log
-z 表示其后的信息记录文件用于记录屏幕的输出
-v 打印执行的每一条sql
-l 记录执行的sql命令道文件中去
-s 参数s的意思是每次遇到报错的命令就停止执行操作，后面的sql就不再执行了
-t 表示使用分号语句终结符
-f 表示其后就是读取sql的脚本文件 从输入文件读内容
-x 省略列标题的打印
~~~

### 从备份文件导入数据到表中

~~~sql
导出：export；导入：import,load；
db2 import from yourFile.del of del replace into yourTable

db2 export to /home/xxxx.del of del select * from tablename
db2 import from /home/xxxx.del of del insert into tablename

db2 export to /home/xxxx.IXF of IXF select * from tablename
db2 import from /home/xxxx.IXF of IXF insert into tablename

数据移动格式：ASC\DEL文本文件；WSF工作表格式，主要用于LOTUS软件；IXF集成交换格式；
   ASC,DEL,WSF在跨平台时可能会导致数据丢失，跨平台用IXF；DEL文件可视，IXF文件不可视。
   export命令：`db2 export to test.ixf of ixf select * from table;`
   import命令：`db2 import from test.ixf of ixf insert_update into table;`
其中插入方式有以下几种：insert、insert_update、replace、replace_create（只支持ixf格式）、create（只支持ixf格式）；

   db2move可以实现整个数据库的转移，语法格式为
   db2move <dbname> <action> <option>;
   其中action包括export、import、copy和load；命令：db2move mydatabase

~~~

### 备份
~~~sql
db2  +c  -tvsf db_yuupdate.sql -z db_yuupdate.log 
db2 rollback
~~~

~~~sql
* 8 on system error
* 4 db2 error (constraint violation, object not found etc)
* 2 db2 warning 
* 1 no rows found

-- 如果需要记录使用-z 指令 ，不能用tee 来打印，会提前提交sql数据。
-- db2 +c -stvf test1.sql | tee table_dd1_1.log 
db2 -l migration.log +c -vstf migration.sql
# -ge 4
if [ $? -ne 0 ]; then
    db2 rollback
    tail -10 migration.log
else
    db2 commit
fi
~~~

~~~sql
#!/bin/bash
 
# 开始连接数据库
db2 connect to xxx
 
# 执行sql
sql="select aa from table"
aas=`db2 ${sql}`
 
# 返回值判断
if [ $? -ne 0 ]
then
#显示db2返回的错误信息
echo "${aas}"
exit 1
fi
 
# 对取得的数据进行处理,循环。
echo "$aas" | sed -e '4,/^$/!d;/^$/d' |
while read aa
do
  echo "当前值：${aa}"
  cp file.sql file_tmp.sql
  #利用替换来达到给sql文件传递参数的效果
  perl -i -pe 's/\#aa/'${aa}'/g' file_tmp.sql
  #执行sql文件。sql文件需已英文分号结尾
  db2 -tf file_tmp.sql
  rm file_tmp.sql
done
 
# 断开数据库连接
db2 terminate

~~~
其中如果select 1 from table fetch  first 1 rows only,如果表中没有数据，则返回失败，不等于0
~~~shell
-eq //equals等于
-ne //no equals不等于
-gt //greater than 大于
-lt //less than小于
-ge //greater equals大于等于
-le //less equals小于等于
~~~

### 重启

~~~sql
-- 查看是否有活动链接
db2 list applications for db db_name
db2 force application all
db2stop
db2start
~~~
## 查询端口
~~~
db2 get dbm cfg |more    ----查询db2数据库配置信息
# 查看端口
db2 get dbm cfg | grep SVCENAME
~~~

### DB2 SCHEMA
1. DB2数据库中用户的概念
Oracle和MySQL数据库的用户都是数据库内部的用户，由数据库来管理。
但DB2不一样，DB2没有数据库用户的概念，它依赖于操作系统用户，也就是操作系统上有这个用户，这个用户才可能连接到DB2数据库。
db2这里模式名字和用户名字可以不一样，db2数据库下创建模式(也可以不创建。不创建的情况下，DB2会使用你连接的用户名作为默认的模式名字)。
- 总结： 一个instance下可以有多个database，一个database下可以有多个schema，1个schema可以被多个user使用。
连接数据库的命令：
~~~
db2 connect to dbname user <username> using <password>
~~~
这里的username是一个操作系统的用户，password是指这个系统用户的密码。

2. DB2中Schema的概念
每个DB2的表的完整名子都是由两部分组成 SchemaName.Tablename，没有例外。
如果访问表的时候，不加Schema的名子，就有一个默认的Schema, 比如运行了以下的SQL语句：
~~~
db2 "create table a.t1 (id int)"
db2 "create table b.t1 (id int)"
~~~
那么就创建了两个表，schema分别是a和b，那么问题来了，如果创建表的时候不指定schema呢？
默认的Schema就是当前用户的用户名，比如当前连接用户是 db2inst1，发出如下SQL
db2 "create table t1(id int)"
那么这个表的schema就是db2inst1

当然，可以使用db2 set current schema来改变当前会话默认的Schema
~~~
db2 "set current schema c"
db2 "create table t1(id int)"
~~~
这个新建的表的完整表名就是c.t1

使用db2 list tables for all可以看到所有表，如下：
 
~~~shell
db2inst1@node01:~> db2 "list tables for all" | grep -iw t1
Table/View                      Schema          Type  Creation time             
------------------------------- --------------- ----- --------------------------
T1                              A               T     2019-08-23-06.53.18.807936
T1                              B               T     2019-08-23-06.53.22.671670
T1                              C               T     2019-08-23-06.53.39.154763
T1                              DB2INST1        T     2019-08-23-06.53.29.872110
可以看到有4条表 A.T1, B.T1, C.T1和DB2INST1.T1，如果现在发出命令db2 "select * from t1"，那么实际访问的是C.T1，因为前面有发出过命令db2 set current schema c
~~~

[DB2新建用户](https://www.cnblogs.com/OliverQin/p/8428019.html)
由于DB2用户必须是系统用户，所以新建系统用户。