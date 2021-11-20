## windows
- 1 添加清华源：命令行中直接输入以下命令
~~~
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge 
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/msys2/
 
# 设置搜索时显示通道地址
conda config --set show_channel_urls yes

# 注意如果需要pytorch, 还需要添加pytorch的镜像
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/
~~~

- 2 添加中科大源
~~~
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/main/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/conda-forge/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/msys2/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/bioconda/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/menpo/
 
conda config --set show_channel_urls yes
~~~
或者修改文件
C:\users\username\.condarc

## Linux
linux：/home/username/.condarc
将以下配置文件写在vim ~/.condarc文件中 
添加如下命令
~~~
channels:
  - https://mirrors.ustc.edu.cn/anaconda/pkgs/main/
  - https://mirrors.ustc.edu.cn/anaconda/cloud/conda-forge/
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
  - defaults
show_channel_urls: true
~~~

删源
换回conda的默认源。查看了conda config的文档后，发现直接删除channels即可。
~~~
conda config --remove-key channels
~~~