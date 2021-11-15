windows 10 修改mac地址
- 第一步：
搜索栏输入regedit，按回车键进入注册表编辑器
- 第二步：
定位到HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Class\{4D36E972-E325-11CE-BFC1-08002bE10318}
- 第三步：
在右侧的DriverDesc值中确定型号，在真实的网卡中新建字符串
命名：NetworkAddress
赋值：x2-xx-xx-xx-xx-xx 或者x6-xx-xx-xx-xx-xx (其中x为16进制字符串)
- 第四步：
禁用 ，启用网卡
查看新mac地址，ipconfig/all