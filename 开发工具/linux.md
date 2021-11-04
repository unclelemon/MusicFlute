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
egrep -E "word1|word2|word3" file.txt  其中"" 双引号必须加上
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
