##  发展历史

- All in one :所有的业务模块都写道一个Service。java文件当中

- MVC/三层架构: 将各个功能根据层次进行了划分，但是所有代码仍然在同一台计算机中编写，并发能力有限。

- RPC(RPC（Remote Procedure Call）远程过程调用):可以让一个项目部署在不同的计算机当中，但是此种模式的ip+端口比较分散，有一定的维护难度。

- SOA面向服务的架构：通过注册中心将分散的ip+端口集中管理。zookeeper是比较著名的注册中心，是一个树形结构。

  Dubbo是一个比较著名的SOA架构

## Dubbo 基础概念

- 注册中心： Registry
- 提供方：Provider
- 消费方： Consumer
- 监听器：Monitor

### 运行步骤

1. Provider将服务放到服务容器Container（例如Tomcat），Container运行Provider提供的服务。

2. 发布服务到注册中心。
3. 消费方Consumer从注册中心订阅服务。
4. 注册中心推送服务（notify），并且会实施监听服务方是否发生改变。
5. 消费方调用服务方提供的服务
6. 监听器可以监听发送方和接收方。

## 开发dubbo程序

### 准备环境

1. linux 安装jdk （rpm版本 rpm -ivh jdk.rpm 默认安装到 /usr/java 目录中）

~~~shell
rpm -ivh jdk.rpm
~~~

- -i 显示套件相关信息

- -v 显示指令执行过程

- -h 或者--hash 套件安装时列出标记

  配置环境变量

  ~~~shell
  vim /etc/profile # 在文件最后追加
  export JAVA_HOME=/usr/java/jdk
  export CLASSPATH=$JAVA_HOME/lib:$CLASSPATH
  export PATH=$JAVA_HOME/bin:$PATH
  source /etc/profile # 刷新环境变量
  ~~~

2. linux中安装注册中心zookeeper

下载apace zookeeper-3.4.12.tar.gz

~~~shell
tar -zxvf zookeeper-3.4.12.tar.gz # 解压
cd zookeeper-3.4.12/conf
mv zoo_sample.cfg zoo.cfg # 重命名配置文件
--------------------------
~~~

更改zoo.cfg配置文件

- 端口号clientPort=2081

- dataDir=全路径/data # 更改为自己想要存放服务的路径文件夹

启动zookeeper

~~~shell
/bin/zkServer.sh start  # 开启
/bin/zkServer.sh stop   # 关闭
/bin/zkServer.sh status # 查看状态
~~~

### 开发服务方

- 开发功能Imp，注册到dubbo和spring容器中

- 配置文件,集成Spring

~~~xml
  <!-- 配置dubbo的应用名称-->
  <dubbo:application name ="app-server" />
  <!-- 配置注册中心地址-->
  <dubbo:registry protocol="zookeeper" address="zookeeper://ip:2081" />
    
  <!-- 配置dubbo扫描包 -->
  <dubbo:annotation package=="org.stau.****"/>
                    
  <context:component-scan base-package="*****"></dubbo:component-scan>
~~~

###  开发消费方

~~~xml
  <!-- 配置dubbo的应用名称-->
  <dubbo:application name ="app-consumer" />
  <!-- 配置注册中心地址-->
  <dubbo:registry protocol="zookeeper" address="zookeeper://ip:2081" />
    
  <!-- 配置dubbo扫描包 -->
  <dubbo:annotation package=="org.stau.controllr"/>
                    
  <!-- 配置Spring扫描包 -->
  <context:component-scan base-package="org.stau.controllr"></dubbo:component-scan>
~~~

从提供方拷贝接口并使用@Reference  将远程服务中的服务类进行注入

### 监视器安装

### 实战

- 传统MVC模块

- 建立父工程，让MVC三个模块都依赖父工程，以便将共用的jar包独立出来
- 建立统一的pojo类，以便MVC三个模块能通用
- 提供相同的代码（接口）

在maven中引入一个maven中不存在的jar包

例如将本地存在的ojdbc7.jar 变成maven格式的jar包

~~~shell
mvn install:install-file -DgroupId=com.oracle -DartifactId=ojdbc7 -Dversion=10.2.0.5.0 -Dpackaging=jar -Dfile=d:\ojdbc7.jar
~~~

在mvn仓库所在的文件夹中执行以上命令。

然后安装依赖

~~~xml
<dependency>
	<groupId>com.oracle</groupId>
  <artifactedId>ojdbc7</artifactedId>
  <version>10.2.0.5.0</version>
</dependency>
~~~

如自己新建了一个students-pojo jar项目。需要被父工程所依赖。从students-pojo 中的pom.xml复制出来，被父工程所依赖

~~~xml
<dependency>
	<groupId>org.dubbo</groupId>
  <artifactedId>students-pojo</artifactedId>
  <version>0.0.1-SNAPSHOT</version>
</dependency>
~~~

目前共有六个模块

1. 视图层 View
2. 模型层 model
3. 控制层Controller
4. 父工程 （存放通用jar包）
5. 公共pojo （被MVC使用）
6. 公共接口 （通用的方法）

父工程依赖公共pojo，dependency公共接口。如此依赖，依赖于父工程的MVC模块就自然而然的引入了pojo。其中的公共接口在model中实现，view使用dubbo依赖于公共接口，在通过dubbo注入服务，就能够调用生产者的服务。

其中View  在一台消费者机器上，model和controller在另外一台生产者机器。

View 使用@Reference 注入在生产者机器的类。而controller 使用dubbo提供的@service提供服务。
