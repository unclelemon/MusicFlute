# 正则表达式
https://www.runoob.com/regexp/regexp-rule.html
## PHP的正则表达式有一些内置的通用字符簇：
字符簇	描述
- [[:alpha:]]	任何字母
- [[:digit:]]	任何数字
- [[:alnum:]]	任何字母和数字
- [[:space:]]	任何空白字符
- [[:upper:]]	任何大写字母
- [[:lower:]]	任何小写字母
- [[:punct:]]	任何标点符号
- [[:xdigit:]]	任何16进制的数字，相当于[0-9a-fA-F]

# 大小写匹配
这里主要将一下，使用正则表达式进行文本的修改和替换，替换时我们需要使用到捕获组，使用()表示，然后在IDEA中可以获取$n拿到捕获组中的值。

如: (\w+)-(\w+)   可以使用 $1 $2 $3 引用分组

主要注意的是，每一个()都代表一个捕获组，使用$n时下标不能乱。

大小写转换

- \l 将字符更改为小写，直到字符串中的下一个字符，例如，BAR 变成 bAR
- \u 将字符更改为大写，直到字符串中的下一个字符，例如，bar 变成 Bar
- \L 将字符更改为小写，直到文字字符串的末尾，例如，BAR 变成 bar
- \U 将字符更改为大写，直到文字字符串的末尾，例如，bar 变成 BAR
例如:

1. 案例1：
XXX-yyy
替换为:
xXXxxx_yyyYYY



2. 案例2：
VLR_OTHER           
VLR_VOLKSWAGEN      
VLR_BUICK           
VLR_BMW             
替换成：
String MPC_VEHICLE_TYPE_VLR_OTHER = "mpc.vehicle.type.vlr.other";
String MPC_VEHICLE_TYPE_VLR_VOLKSWAGEN = "mpc.vehicle.type.vlr.volkswagen";
String MPC_VEHICLE_TYPE_VLR_BUICK = "mpc.vehicle.type.vlr.buick";

只列举一部分，其实有几百个，需要定义成字符串常量，这个重复的工作枯燥还容易错，主要牵扯到大小写，还要将“_”替换成“.”



查找：([A-Za-z0-9]+)_([A-Za-z0-9]+)

替换：String MPC_VEHICLE_TYPE_$1_$2 = "mpc.vehicle.type.\L$1.\L$2" 

3. 案例3：
VLR_VOLKSWAGEN      = 1,    //大众
VLR_BUICK           = 2,    //别克
VLR_BMW             = 3,    //宝马
替换成：
VLR_VOLKSWAGEN(1, "大众"),
VLR_BUICK(2, "别克"),
VLR_BMW(3, "宝马"),

只列举一部分，其实有几百个，主要将给定的车辆品牌定义为枚举类，这几百个如果人工修改那就那难，而且还要保证不能出错，数字要和品牌进行对应



查找：VLR_([A-Za-z0-9]+) *= *(\d+), *//([\u4e00-\u9fa5A-Za-z0-9]+)

替换：VLR_$1($2, "$3"),

# 贪婪匹配和非贪婪匹配
- 贪婪匹配:正则表达式一般趋向于最大长度匹配，也就是所谓的贪婪匹配。
- 非贪婪匹配：就是匹配到结果就好，就少的匹配字符。
比如在对html文件进行正则匹配时：
~~~
<span>a</span><span>b</span>
~~~
如果我们使用.*匹配
~~~
<span>.*</span>
~~~
匹配结果是
~~~
<span>a</span><span>b</span>
~~~
很明显不能满足我们的需求。
这时候我们就需要使用非贪婪式匹配，也可以称为懒惰匹配,懒惰匹配与贪婪匹配的差别就是在贪婪匹配的后面加个?号，同样对上面的html进行匹配，我们使用.*?，匹配结果是
~~~
<span>a</span>
~~~
和我们预期效果一样。
除了.*?这种匹配方式,非贪婪式匹配还包括：
~~~
.+?　　　　匹配一次
.??　　　　不匹配
.{m,n}?   匹配m次
~~~