## grep
grep 命令非常常用，经常用于匹配文本字符。

一般情况下，grep 命令只能匹配一个关键字，怎么实现匹配多个关键字呢？

下面看几种情况：

- grep ‘字符串’
这是最普通的使用方式：

匹配文件中包含 MANPATH 的那一行：
cat manpath.config | grep 'MANPATH'

- grep -v ‘字符串’
与上例子相反，反向匹配

匹配文件中不包含 MANPATH 的那一行：
cat manpath.config| grep -v 'MANPATH'

- grep -E 同时匹配多个关键字–或关系  (正则 等价于egrep)
grep -E "word1|word2|word3" file.txt
等价于：
egrep  "word1|word2|word3" file.txt  其中"" 双引号必须加上
匹配 file.txt 中包含 word1 或 word2 或 word3 的行。
满足其中任意条件（word1、word2和word3之一）就会匹配。

- 同时匹配多个关键字–与关系
使用管道符连接多个 grep ，间接实现多个关键字的与关系匹配：
grep word1 file.txt | grep word2 |grep word3
必须同时满足三个条件（word1、word2和word3）才匹配。

### 常见参数：
- grep -i main haison.c #忽略大小写
- grep –n main haison.c #输出行号
- grep -v main haison.c #反检索，只显示不匹配的行



## xargs
- xargs 的一个选项 -I，使用 -I 指定一个替换字符串 {}，这个字符串在 xargs 扩展时会被替换掉，当 -I 与 xargs 结合使用，每一个参数命令都会被执行一次：
find ./ | xargs -I {} grep -n word1 {} | grep word2 

## 清除日志

错误方法：

rm -f logfile
1
原因：
应用已经打开文件句柄，直接删除会造成：

应用无法正确释放日志文件和写入
显示磁盘空间未释放，最后磁盘空间占用100%
正确方法：

- 第一种：cat /dev/null > filename

- 第二种：: > filename

- 第三种：> filename

- 第四种：echo “” > filename

- 第五种：echo > filename

rm -f 之后解决办法：

找到删除文件所在的分区，查看当前系统句柄未释放情况
lsof -n /opt | grep deleted
或
lsof | grep deleted 其中出现的第一个数字为进程号

输出里面有进程号

kill 进程号

运行

bash lsof -n /opt | grep delete
查看是否还存在删除了但未释放空间的文件，应该没了。

再次运行df -lh查看空间是否已经释放了。
.
还有一种解决方式：（假设文件名为a.log）

清空文件内容
> a.log

删除该文件

rm -rf a.log

### 建议

删除文件使用mv

### 打包命令tar
~~~shell
tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
~~~
选项与参数
- -c :建立打包文件，可搭配-v 查看过程中被打包的文件名（filename）；
- -t :查看打包文件的内容包含有哪些文件名，重点在查看【文件名】
- -x :解包或解压缩的功能，可以搭配-C(大写)在特定目录解压，特别留意的是，-c,-t,-x不可同时出现在一串命令行中
- -z :通过gzip的支持进行压缩/解压缩：此时的文件名最好为*.tar.gz;
- -j :通过bzips2的支持进行压缩/解压缩：此时的文件名最好为*.tar.bz2;
- -J :通过xz的支持进行压缩/解压缩：此时的文件名最好为*.tar.xz;特别留意，-z,-j,-J不可同时出现在一串命令行中
- -v :在压缩/解压缩的过程中，将正在处理的文件名显示出来；
- -f :后面要立刻接要被处理的文件名，建议-f单独写一个选项，（比较不会忘记）
- -C :目录  这个选项用在解压缩
- -p 小写 :保留备份文件的原本权限和属性
- -P 大写 :保留绝对路径
- -exclude=File : 在压缩的过程中，不要将file打包

### vim删除swp
  这是因为，在用vim打开一个文件时，其会产生一个filename.swap文件，用于保存数据，当文件非正常关闭时，可用此文件来恢复，当正常关闭时，此文件会被删除，非正常关闭时，不会被删除，所以提示存在.swap文件，此时你可以恢复文件： vim -r filename.c 恢复以后把.swap文件删掉，在打开时就不会用提示良，注意.swap文件是个隐藏文件。可用：la查看。以.开头的是隐藏文件。
- ls -a

- vm -r 可查看交换文件