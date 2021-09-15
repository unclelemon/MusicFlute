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