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
 
#开始连接数据库
db2 connect to xxx
 
#执行sql
sql="select aa from table"
aas=`db2 ${sql}`
 
#返回值判断
if [ $? -ne 0 ]
then
#显示db2返回的错误信息
echo "${aas}"
exit 1
fi
 
#对取得的数据进行处理,循环。
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
 
#断开数据库连接
db2 terminate

~~~

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
