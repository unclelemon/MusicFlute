## 遍历文件
~~~shell
#!/bin/bash
# get all filename in specified path

path=$1
files=$(ls $path)
for filename in $files
do
   echo $filename >> filename.txt
done
~~~


## Shell脚本中判断字符串是否被包含在内
1. 字段 grep：
案例：
~~~
str1="abcdefgh"
str2="def"
result=$(echo $str1 | grep "${str2}")
if [[ "$result" != "" ]];then
    echo "包含"
else
    echo "不包含"
fi
~~~

2. 字符串运算符 =~:
案例：
~~~
str1="abcdefgh"
str2="def"
if [[ $str1 =~ $str2 ]];then
    echo "包含"
else
    echo "不包含"
fi
~~~

3. 正则表达式中的通配符 *:
案例：
~~~shell
str1="abcdefgh"
str2="def"
if [[ $str1 == *$str2* ]];then
    echo "包含"
else
    echo "不包含"
fi
~~~