## Windows 安装docker
windows 10 企业版，旗舰版 等 除去家庭版 可以开启虚拟化技术。安装docker。

在任务管理器中-》性能-》查看虚拟化技术 enable 或者disable；
如果是disable ，进入bios。一般开机按F2,或者按照各个厂商的主板进入bios。开启虚拟化技术。

## windows 到 docker 互传文件
使用ps 查看所有在运行的docker容器
docker ps 
最后一列是容器名字
![Image description](https://images.gitee.com/uploads/images/2021/1120/193543_a63b16bf_7472878.png "截屏2021-11-20下午7.34.23.png")

- 从windows传文件到docker
docker cp filename 容器名字:容器路径 
~~~
docker cp d:\file nusld:/root/sdjk/
~~~
- 从docker传文件到windows
docker cp 容器名字:容器路径 filename 