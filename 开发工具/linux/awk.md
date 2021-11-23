寻找A文件有，B文件没有的数据
~~~shell
# 查找到一条不同数据就停止运行
# 读取filename1.txt中的每一行数据，之后将其中在filename2.txt中出现过的数据删除，打印留下的数据。
awk '{if(ARGIND == 1) {val[$0]}else{if($0 in val) delete val[$0]}} END{for(i in val) {print i ; break}}' filename1.txt filename2.txt
~~~