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

## 当一个脚本需要传入的参数较多时，可以使用for循环进行参数遍历
~~~
#!/bin/bash
number=65       #定义一个退出值
index=1          #定义一个计数器
if [ -z "$1" ];then              #对用户输入的参数做判断，如果未输入参数则返回脚本的用法并退出，退出值65
  echo "Usage:$0 + canshu"
  exit $number
fi
echo "listing args with \$*:"         #在屏幕输入，在$*中遍历参数
for arg in $*                     
do
  echo "arg: $index = $arg"         
  let index+=1
done
echo
index=1                       #将计数器重新设置为1
echo "listing args with \"\$@\":"    #在"$@"中遍历参数
for arg in "$@"
do
  echo "arg: $index = $arg"
  let index+=1
done
~~~