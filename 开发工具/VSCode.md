## 下载
~~~shell
# 镜像源的替换
vscode.cdn.azure.cn
将
https://az764295.vo.msecnd.net/stable/3a6960b964327f0e3882ce18fcebd07ed191b316/VSCodeUserSetup-x64-1.62.2.exe
替换成
https://vscode.cdn.azure.cn/stable/3a6960b964327f0e3882ce18fcebd07ed191b316/VSCodeUserSetup-x64-1.62.2.exe
~~~

## vscode设置成中文

vscode默认的语言是英文，对于英文不好的小伙伴可能不太友好。简单几步教大家如何将vscode设置成中文。

1. 按快捷键“Ctrl+Shift+P”。
2. 在“vscode”顶部会出现一个搜索框。
3. 输入“configure language”，然后回车。
4. “vscode”里面就会打开一个语言配置文件。
5. 将“en-us”修改成“zh-cn”。
6. 按“Ctrl+S”保存设置。
7. 关闭“vscode”，再次打开就可以看到中文界面了。

## 常用快捷键

以下以Windows为主，windows的 Ctrl，mac下换成Command就行了

对于 **行** 的操作：

- 重开一行：光标在行尾的话，回车即可；不在行尾，ctrl` + enter` 向下重开一行；ctrl+`shift + enter` 则是在上一行重开一行
- 删除一行：光标没有选择内容时，ctrl` + x` 剪切一行；ctrl +`shift + k` 直接删除一行
- 移动一行：`alt + ↑` 向上移动一行；`alt + ↓` 向下移动一行
- 复制一行：`shift + alt + ↓` 向下复制一行；`shift + alt + ↑` 向上复制一行
- ctrl + z 回退

对于 **词** 的操作：

- 选中一个词：ctrl` + d`

搜索或者替换：

- ctrl` + f` ：搜索
- ctrl` + alt + f`： 替换
- ctrl` + shift + f`：在项目内搜索

通过**Ctrl + `** 可以打开或关闭终端

Ctrl+P 快速打开最近打开的文件

Ctrl+Shift+N 打开新的编辑器窗口

Ctrl+Shift+W 关闭编辑器

Home 光标跳转到行头

End 光标跳转到行尾

Ctrl + Home 跳转到页头

Ctrl + End 跳转到页尾

Ctrl + Shift + [ 折叠区域代码

Ctrl + Shift + ] 展开区域代码

Ctrl + / 添加关闭行注释

Shift + Alt +A 块区域注释

## 设置大小写切换

1、按ctrl+k,s，即打开快捷键修改的界面。也可以在文件-首选项-键盘快捷方式里面打开。

2、在这个页面搜索"转换为大写/转换为小写"，或者“upper/lower”。

3、双击该行或者点左边的加号，可以录入快捷键，录入完毕按回车保存

大写alt+d， 小写 alt + x

## 删除空行

~~~java
([\/]\*.+?\*[\/]|//.+)    //清除注释
^\s*(?=\r?$)\n            // 删除空行
~~~

