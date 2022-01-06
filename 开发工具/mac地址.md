windows 10 修改无线网卡mac地址
- 第一步：
搜索栏输入regedit，按回车键进入注册表编辑器
- 第二步：
定位到HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Class\{4D36E972-E325-11CE-BFC1-08002bE10318}
- 第三步：
在右侧的DriverDesc值中确定型号（可以通过适配器选项），在真实的网卡中新建字符串
![网络适配器](../images/%E9%80%82%E9%85%8D%E5%99%A8.png)

命名：NetworkAddress
赋值：x2-xx-xx-xx-xx-xx 或者x6-xx-xx-xx-xx-xx (其中x为16进制字符串)
- 第四步：
更改适配器选项，
禁用 ，启用网卡
查看新mac地址，ipconfig/all

note:36-C9-3D-EA-B0-0A
删除Networkaddress，之后重启，恢复原有ip地址。