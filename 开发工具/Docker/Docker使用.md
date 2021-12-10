### 常用参数

| Options | Mean                                               |
| ------- | -------------------------------------------------- |
| -i      | 以交互模式运行容器，通常与 -t 同时使用；           |
| -t      | 为容器重新分配一个伪输入终端，通常与 -i 同时使用； |
| -d      | 后台运行容器，并返回容器ID；                       |

## docker pull

~~~
docker pull christoftorres/honeybadger 
~~~

## docker build

### 实例

使用当前目录的 Dockerfile 创建镜像，标签为 runoob/ubuntu:v1。 仓库名/镜像名：TAG名

-t tag标签名 -m message信息

```
docker build -t runoob/ubuntu:v1 . 
```

查看是否创建成功,查看镜像

~~~
docker images
~~~

运行

~~~
docker run -i -t christoftorres/honeybadger
~~~

## docker create

docker create命令能够基于镜像创建容器。

使用docker镜像nginx:latest创建一个容器,并将容器命名为myrunoob

~~~shell
runoob@runoob:~$ docker create  --name myrunoob  nginx:latest      
09b93464c2f75b7b69f83d56a9cfc23ceb50a48a9db7652ee4c27e3e2cb1961f
~~~

该命令执行的效果类似于docker run -d，即创建一个将在系统后台运行的容器。
但是与docker run -d不同的是，docker create创建的容器并未实际启动，还需要执行docker start命令或docker run命令以启动容器。
事实上，docker create命令常用于在启动容器之前进行必要的设置。

### 实例

```ruby
# 只创建了容器
[root@localhost ~]# docker create nginx:1
520d126da89b27ebab5a44acf83e867fa61a5ccaa8c1ab1774a2009e1b1f5949
[root@localhost ~]# docker ps -a
CONTAINER ID        IMAGE               COMMAND                   CREATED             STATUS              PORTS                  NAMES
520d126da89b        nginx:1             "/bin/sh -c '[\"/usr/…"   4 seconds ago       Created                                    musing_wescoff
a75dfeff0898        centos_nginx:1      "/nginx.sh"               2 days ago          Up 2 days                                  confident_wilson
f61fe8c1b847        nginx:1             "/bin/bash"               3 days ago          Up 3 days           0.0.0.0:8080->80/tcp   epic_ride
cf2c588a7b30        centos              "/bin/bash"               6 days ago    
```

## dockerfile

~~~shell
FROM ubuntu:16.04

MAINTAINER Christof Torres (christof.torres@uni.lu)

SHELL ["/bin/bash", "-c"]
RUN apt-get update
RUN apt-get install -y sudo wget tar unzip pandoc python-setuptools python-pip python-dev python-virtualenv git build-essential software-properties-common
RUN add-apt-repository -y ppa:ethereum/ethereum
RUN apt-get update

# COPY z3-4.7.1-x64-ubuntu-16.04.zip honeybadger
# Install z3
# RUN unzip z3-4.7.1-x64-ubuntu-16.04.zip && rm z3-4.7.1-x64-ubuntu-16.04.zip && mv z3-4.7.1-x64-ubuntu-16.04/bin/* /usr/local/bin && rm -r z3-4.7.1-x64-ubuntu-16.04
RUN wget https://github.com/Z3Prover/z3/releases/download/z3-4.7.1/z3-4.7.1-x64-ubuntu-16.04.zip --no-check-certificate && unzip z3-4.7.1-x64-ubuntu-16.04.zip && rm z3-4.7.1-x64-ubuntu-16.04.zip && mv z3-4.7.1-x64-ubuntu-16.04/bin/* /usr/local/bin && rm -r z3-4.7.1-x64-ubuntu-16.04
# Install solidity
RUN wget https://github.com/ethereum/solidity/releases/download/v0.4.25/solidity_0.4.25.tar.gz && tar -xvzf solidity_0.4.25.tar.gz && rm solidity_0.4.25.tar.gz && cd solidity_0.4.25 && ./scripts/install_deps.sh && ./scripts/build.sh && cd .. && rm -r solidity_0.4.25
# Install go
RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz && tar -xvf go1.9.2.linux-amd64.tar.gz && rm go1.9.2.linux-amd64.tar.gz && cp go/bin/* /usr/local/bin && mv go /usr/local && mkdir -p ~/go; echo "export GOPATH=$HOME/go" >> ~/.bashrc && echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc
# Install geth
RUN wget https://github.com/ethereum/go-ethereum/archive/v1.8.16.zip && unzip v1.8.16.zip && rm v1.8.16.zip && cd go-ethereum-1.8.16 && make all && mv build/bin/* /usr/local/bin && cd .. && rm -r go-ethereum-1.8.16
# Create virtualenv
RUN virtualenv venv && source venv/bin/activate
# Upgrade pip
RUN pip install --upgrade pip setuptools wheel
# Install requests
RUN pip install requests
# Install web3
RUN pip install web3==0.1.9

WORKDIR /root
COPY datasets/honeypots honeypots
COPY honeybadger honeybadger

~~~

## docker 删除镜像
1. 查询所有镜像
  ~~~
  docker ps -a
  
  CONTAINER ID   IMAGE         COMMAND                  CREATED       STATUS                   PORTS                                                  NAMES
  925a665ae7b6   mysql:5.7     "docker-entrypoint.s…"   7 weeks ago   Up About an hour         0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql
  92f7c45445f2   hello-world   "/hello"                 7 weeks ago   Exited (0) 7 weeks ago                                                          elegant_booth
  ~~~
  可以看出来上面有 mysql：5.7 和hello-world两个镜像

2. 停止容器运行

   该容器的ID 可以看出来是：925a665ae7b6

    ~~~
    docker stop 925a665ae7b6
    ~~~

3. 删除容器

   ~~~
   docker rm 925a665ae7b6
   ~~~

4. 删除镜像

   ~~~
   docker rmi mysql:5.7
   
   
   Untagged: mysql:5.7
   Untagged: mysql@sha256:b8814059bbd9c80b78fe4b2b0b70cd70fe3772b3c5d8ee1edfa46791db3224f9
   Deleted: sha256:8a8a506ccfdc7699b62e0818774fa06c7e1f30d17b4695d2c8be42848870e2ef
   Deleted: sha256:ba884392d0236efc0b56a5177c1d95fa11627aeeb1a094ddb6b79af8a974f239
   Deleted: sha256:e24812a440029006d5515d620552882486acac4685ddc1cee8c41114871d5b33
   Deleted: sha256:e5bf95b5be3c9b3a31d22bee3844f30b7eabf1d56186ab13df5fd7635f3e5947
   Deleted: sha256:57cd1e9fd1efe26b2bd726093c561c7ff90edbb3392e6ab94dc54a669b45a489
   Deleted: sha256:80f5487a88b8061855e99782979ed6069a8dd1c7dfbb1eb63fe42a4a9d119436
   Deleted: sha256:f791a6c727931d41c51f8bf24ee32a4dbf0169f372b174f1ff89b4836b97c48e
   Deleted: sha256:4c88df098412e11a98936509f3cede57f87154b350b0f75d96713f6e1dd56101
   Deleted: sha256:fdba3a2cd286d9a5f65fc00f5254048855ae7dc00f3b3fa3356981eb9a7fe6d0
   Deleted: sha256:8b3a69042e0da82429d28be0c474e73290ba4908730de22b2200a7aac9b245bd
   Deleted: sha256:90afe56a0643f5bf1b1e8ee147b40a8e12b3fdd7e26bc2d2c50180d68dd524d0
   Deleted: sha256:e81bff2725dbc0bf2003db10272fef362e882eb96353055778a66cda430cf81b
   ~~~
## 安装mysql-docker

### 配置文件（可跳过）

~~~
# 进入已启动的容器
docker exec -it mysql bin/bash
# 退出进入的容器
exit;

创建my.cnf配置文件
touch /mydata/mysql/my.cnf

因为有目录映射，所以我们可以直接在镜像外执行
vi /mydata/mysql/conf/my.cnf 

my.conf添加如下内容：
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
default_authentication_plugin=mysql_native_password
init_connect='SET collation_connection = utf8_unicode_ci'
init_connect='SET NAMES utf8'
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
skip-name-resolve

保存(注意评论区该配置不对，不是collection而是collation)
# 
log_timestamps=SYSTEM：日志记录使用UTC时区，需要修改成系统时区
character-set-server=utf8：设置字符集为utf8
default_authentication_plugin：mysql8.0后的默认认证插件是caching_sha2_password，会导致我们用Navicat连接不上的问题，避免麻烦直接改成以前的认证方式
lower_case_table_names=1：
因为windows默认是1，linux默认是0，Mac OS X默认是 2，开发中很容易因为本地数据库和线上数据库配置不统一，而导致出问题（windows和mac不能设置为0，最多改成1和2）
下面比较的意思是指：在代码里sql用到表名的地方，用大写表名能否查出数据库里的小写表名。
0：创建的时候区分大小写，查询比较的时候也区分大小写
1：代表创建表是一律小写的，但比较的时候不区分大小写（输入大写表名也能查到）
2：代表创建表区分大小写，但比较的时候是小写的

docker restart mysql

~~~



### 安装

注意：mysql 8.0 必须要有下面这一行

-v /mydata/mysql/mysql-files:/var/lib/mysql-files \

   ~~~
   # 下载
   sudo docker pull mysql:5.7
   
   # 启动 mysql 5.0  
   # 注意-d 后面跟的是容器名
   # --name指定容器名字 -v目录挂载 -p指定端口映射  -e设置mysql参数 -d后台运行
   sudo docker run -p 3306:3306 --name mysql5 \
   -v /mydata/mysql/log:/var/log/mysql \
   -v /mydata/mysql/data:/var/lib/mysql \
   -v /mydata/mysql/conf:/etc/mysql \
   -v /mydata/mysql/mysql-files:/var/lib/mysql-files \
   -e MYSQL_ROOT_PASSWORD=12345678 \
   -d mysql:5.7
   
   MYSQL_ROOT_PASSWORD ：root的密码
   MYSQL_USER ：建一个普通用户
   MYSQL_PASSWORD ：普通用户的密码
   
   # 进入已启动的容器
   docker exec -it mysql bin/bash
   # 退出进入的容器
   exit;
   
   因为有目录映射，所以我们可以直接在镜像外执行
   vi /mydata/mysql/conf/my.conf 
   
   [client]
   default-character-set=utf8
   [mysql]
   default-character-set=utf8
   [mysqld]
   init_connect='SET collation_connection = utf8_unicode_ci'
   init_connect='SET NAMES utf8'
   character-set-server=utf8
   collation-server=utf8_unicode_ci
   skip-character-set-client-handshake
   skip-name-resolve
   
   # 重启
   docker restart mysql
   
   ~~~
mysql 8.0
~~~
docker run -d -p 3806:3306 --name mysql8x --privileged=true -v /docker/mysql/conf:/etc/mysql/conf.d -v /docker/mysql/logs:/logs -v /docker/mysql/data80:/var/lib/mysql -v /etc/localtime:/etc/localtime -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

参数说明：
-p 3806:3306：把容器内的3306端口映射到本机的3806端口，我们远程连接的时候也是连3806
–privileged=true：大约在0.6版，privileged被引入docker。
使用该参数，container内的root拥有真正的root权限。否则，container内的root只是外部的一个普通用户权限。privileged启动的容器，可以看到很多host上的设备，并且可以执行mount。甚至允许你在docker容器中启动docker容器。
-v /docker/mysql/conf/my.cnf:/etc/my.cnf：映射配置文件
-v /docker/mysql/data:/var/lib/mysql：映射数据目录
-v /etc/localtime:/etc/localtime：映射linux时间文件（为了让容器时间和主机时间同步）
-e MYSQL_PASSWORD=“123456

设置环境变量参数说明，这里我为了方便只设置了root密码：
-e MYSQL_USER=“woniu” ：添加woniu用户
-e MYSQL_PASSWORD=“123456”：设置添加的用户密码
-e MYSQL_ROOT_PASSWORD=“123456”：设置root用户密码

~~~



###   Mysql操作

   ~~~
   # 进入已启动的容器
   docker exec -it mysql bin/bash
   
   # 查看版本
   mysql -V
   
   #登录mysql
   mysql -u root -p
   ALTER USER 'root'@'localhost' IDENTIFIED BY 'hzlove321';
   
   #添加远程登录用户
   CREATE USER 'hzlin'@'%' IDENTIFIED WITH mysql_native_password BY '12345678';
   GRANT ALL PRIVILEGES ON *.* TO 'hzlin'@'%';
   ~~~

### 查看启动日志

~~~
docker logs mysql

mysqld: Error on realpath() on '/var/lib/mysql-files' (Error 2 - No such file or directory)
2021-12-10T05:45:34.555629Z 0 [ERROR] [MY-010095] [Server] Failed to access directory for --secure-file-priv. Please make sure that directory exists and is accessible by MySQL Server. Supplied value : /var/lib/mysql-files
2021-12-10T05:45:34.555651Z 0 [ERROR] [MY-010119] [Server] Aborting

# 解决方法 添加如下一行
-v /mydata/mysql/mysql-files:/var/lib/mysql-files \
~~~



## docker 配置加速器

中国官方加速镜像

--registry-mirror=https://registry.docker-cn.com

网易163镜像加速

--registry-mirror=http://hub-mirror.c.163.com

中科大镜像加速

--registry-mirror=https://docker.mirrors.ustc.edu.cn

阿里云镜像加速

--registry-mirror=https://{your_id}.mirror.aliyuncs.com

daocloud镜像加速

--registry-mirror=http://{your_id}.m.daocloud.io

1. 创建文件夹

   ~~~
   # mkdir -p ：递归创建目录，即使上级目录不存在，会按目录层级自动创建目录
   [root@centos-7 go-gin-chat]# mkdir -p /etc/docker
   ~~~
   
2. 创建文件
   ~~~
   [root@centos-7 go-gin-chat]# vim /etc/docker/daemon.json
   内容如下
   {
       "registry-mirrors": ["http://f1361db2.m.daocloud.io"]     
   }
   或者
   sudo mkdir -p /etc/docker
   sudo tee /etc/docker/daemon.json <<-'EOF'
   {
     "registry-mirrors": ["https://chqac97z.mirror.aliyuncs.com"]
   }
   EOF
   
   # 出现镜像依旧下载很慢，可以更换镜像重新下载
   # 或者 使用下面的命令
   sudo mkdir -p /etc/docker
   sudo tee /etc/docker/daemon.json <<-'EOF'
   {
     "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]
   }
   EOF
   
   sudo systemctl daemon-reload
   sudo systemctl restart docker
   
   ~~~
   
3. 重启
   ~~~
   # 重启 daemon 守护进程
   sudo systemctl daemon-reload
   sudo systemctl restart docker
   
   ~~~


### 镜像拉取错误

~~~
dig @114.114.114.114 registry-1.docker.io 

;; ANSWER SECTION:
registry-1.docker.io.	46	IN	A	3.209.182.229
registry-1.docker.io.	46	IN	A	18.214.230.110
registry-1.docker.io.	46	IN	A	54.85.56.253
registry-1.docker.io.	46	IN	A	3.226.210.61
registry-1.docker.io.	46	IN	A	54.161.109.204
registry-1.docker.io.	46	IN	A	52.72.232.213
registry-1.docker.io.	46	IN	A	3.213.204.48
registry-1.docker.io.	46	IN	A	34.192.145.113
~~~

修改hosts

~~~
vim /etc/hosts
或者
cat >> /etc/hosts<<EOF 末尾追加一行
或者 其中 -e 表示激活转义
echo -e "hellow\nworld" >> lb.txt

3.209.182.229 registry-1.docker.io
~~~

重启docker

~~~
systemctl restart docker
~~~

