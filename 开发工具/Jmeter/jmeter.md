[参考博客1](https://dzone.com/articles/how-to-load-test-wsdl-authentication-with-jmeter)



Add an HTTP Header Manager to the HelloUser sample.

HelloWorld -> Add -> Config Element -> HTTP Header Manager

Add two rows:

- *Content-Type: text/xml; charset=utf-8*
- *SOAPAction: "http://tempuri.org/HelloUser"*

Xpath语法
正如json提取器需要准确编写json表达式一样，xpath提取器也需要准确编写xpth表达式。这里简单说明一下xpath表达式的写法，xpath表达式分为相对路径和绝对路径两种写法，通常采用相对路径的写法。
（1）通过元素属性编写表达式：//标签[@属性=‘属性值’]，示例：//input[@name='username']
同时使用两个属性定位：//input[@name='username'and @class='input']
同名元素拥有某个属性而没有另一个属性：//input[@name='username'and not(@class)]
（2）通过索引编写表达式：通过元素层级定位，示例：//header/div/nav/ui/li[3](注意:xpath中序号从1开始)
（3）通过text编写表达式：text为完全匹配，需要输入全部文字；示例：//a[text()='登录']
（4）通过contains编写表达式：contains为部分匹配，只需输入部分文字；示例：//a[contains(text(),'登录')]或//a[contains(text(),'登')]
（5）通过string编写表达式：元素内容被子元素截断时，通过text无法定位，可以使用string；示例：//a[string()='登录']

### 案例

右键单击线程组元素并选择“添加 --> 配置元素 --> 用户定义变量”以添加“用户定义变量”元素。在用户定义的变量 UI 屏幕上，单击“添加”按钮。设置变量的“名称”并将值保留为空白。

标签：***<ABCD>TARGET</ABCD>***

Xpath：**//\*[local-name()='ABCD']/text ()**

右键单击“线程组”元素并选择“***添加 --> 后处理器 --> BeanShell Postpocessor*** ”

您可以使用为 bean shell 编写的以下脚本

**打印
**您在上面创建的用户定义变量的值：**print("Beanshell processing SOAP response");** **打印（“ACQIDD”+${ACQIDD}）；**