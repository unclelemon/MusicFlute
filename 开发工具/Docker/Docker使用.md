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

