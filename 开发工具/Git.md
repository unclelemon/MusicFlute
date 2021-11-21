#### A simple command-line tutorial:

Git global settings:

~~~shell
git config --global user.name "Unclecode"
git config --global user.email "linhaizeng163@163.com"
~~~

Create git repository:

```shell
mkdir dev-tool
cd dev-tool
git init
touch README.md
git add README.md
git commit -m "first commit"
git remote add origin https://gitee.com/unclecode/dev-tool.git
git push -u origin master
```

Existing repository?

```shell
cd existing_git_repo
git remote add origin https://gitee.com/unclecode/dev-tool.git
git push -u origin master
```

#### 配置github的ip

第一步：

[查找ip地址](https://websites.ipaddress.com/github.com)

第二步:

配置ip

- Windows

  C:\Windows\System32\drivers\etc\hosts

  ~~~
  140.82.112.4 github.com
  ~~~

- Linux

  ~~~shell
  sudo vi /etc/hosts
  ~~~

### Git镜像加速

也就是github clone加速的时候，前面的域名用https://hub.fastgit.org/ 来替换掉https://github.com就行啦！