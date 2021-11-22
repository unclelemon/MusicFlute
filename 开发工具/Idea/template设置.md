[参考文章](https://mp.weixin.qq.com/s/klijSAr3SR-12toIduOygw)
1. 类注释
打开 IDEA 的 Settings，点击 Editor-->File and Code Templates，点击右边 File 选项卡下面的Class，在其中添加图中红框内的内容：
/**
 * @author jitwxs
 * @date ${YEAR}年${MONTH}月${DAY}日 ${TIME}
 */

![类注释](https://images.gitee.com/uploads/images/2021/1122/093746_0121e064_7472878.png "屏幕截图.png")

2. 方法注释
不同于目前网络上互相复制粘贴的方法注释教程，本文将实现以下功能：
```
根据形参数目自动生成 @param 注解
根据方法是否有返回值智能生成 @Return 注解
```
相较于类模板，为方法添加注释模板就较为复杂，首先在 Settings 中点击 Editor-->Live Templates。
点击最右边的 +，首先选择 2. Template Group... 来创建一个模板分组：
![模板分组](https://images.gitee.com/uploads/images/2021/1122/094051_b83215b2_7472878.png "屏幕截图.png")
在弹出的对话框中填写分组名，本文这里取名 userDefine：
![分组名](https://images.gitee.com/uploads/images/2021/1122/094103_de064032_7472878.png "屏幕截图.png")
然后选中刚刚创建的模板分组 userDefine，然后点击 +，选择 1. Live Template：
![live Template](https://images.gitee.com/uploads/images/2021/1122/094114_1764afd6_7472878.png "屏幕截图.png")
此时就会创建了一个空的模板，我们修改该模板的 Abbreviation、Description 和 Template text。需要注意的是，Abbreviation 必须为 *，最后检查下 Expand with 的值是否为 Enter 键。
![Expand with](https://images.gitee.com/uploads/images/2021/1122/094130_c95a2069_7472878.png "屏幕截图.png")
上图中· Template text 内容如下，请直接复制进去，需要注意首行没有 /，且 \* 是顶格的。

```
*
 * 
 * @author jitwxs
 * @date $date$ $time$$param$ $return$
 */
```

注意到右下角的 No applicable contexts yet 了吗，这说明此时这个模板还没有指定应用的语言：
![指定语言](https://images.gitee.com/uploads/images/2021/1122/094149_264a4fa3_7472878.png "屏幕截图.png")
点击 Define，在弹框中勾选Java，表示将该模板应用于所有的 Java 类型文件。
![语言类型](https://images.gitee.com/uploads/images/2021/1122/094158_38e33ef2_7472878.png "屏幕截图.png")
设置 applicable contexts
还记得我们配置 Template text 时里面包含了类似于 $date$ 这样的参数，此时 IDEA 还不认识这些参数是啥玩意，下面我们对这些参数进行方法映射，让 IDEA 能够明白这些参数的含义。点击 Edit variables 按钮：
![输入图片说明](https://images.gitee.com/uploads/images/2021/1122/094211_543c664c_7472878.png "屏幕截图.png")
为每一个参数设置相对应的 Expression：
![输入图片说明](https://images.gitee.com/uploads/images/2021/1122/094219_80bc707d_7472878.png "屏幕截图.png")
设置 Expression
需要注意的是，date 和 time 的 Expression 使用的是 IDEA 内置的函数，直接使用下拉框选择就可以了，而 param 这个参数 IDEA 默认的实现很差，因此我们需要手动实现，代码如下：
```
groovyScript("def result = '';def params = \"${_1}\".replaceAll('[\\\\[|\\\\]|\\\\s]', '').split(',').toList(); for(i = 0; i < params.size(); i++) {if(params[i] != '')result+='* @param ' + params[i] + ((i < params.size() - 1) ? '\\r\\n ' : '')}; return result == '' ? null : '\\r\\n ' + result", methodParameters())
```
另外 return 这个参数我也自己实现了下，代码如下：
```
groovyScript("return \"${_1}\" == 'void' ? null : '\\r\\n * @return ' + \"${_1}\"", methodReturnType())
```
注：本文没有勾选了 Skip if defined 属性，它的意思是如果在生成注释时候如果这一项被定义了，那么鼠标光标就会直接跳过它。我并不需要这个功能，因此有被勾选该属性。
点击 OK 保存设置，大功告成！