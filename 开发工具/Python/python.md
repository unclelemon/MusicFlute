### Python操作文件

###### 利用python来创建文件，并且写入文件内容

~~~python
def create_file():   
  # 读取目录下所有文件名 
  for filename in os.listdir(data):        
    print(filename)        
    # 打开文件。若文件不存在则创建文件        
    file = open(honeypot + filename + ".c", 'w')        
    # 写入内容        
    file.write("1")        
    # 关闭文件        
    file.close()
if __name__ == '__main__':    
  create_file()

~~~

### LOG

由于在logging.basicConfig()中的level 的值设置为logging.DEBUG, 所有debug, info, warning, error, critical 的log都会打印到控制台。

默认为warning，输出到屏幕上，info不会输出到屏幕上。

~~~python
import logging
logging.basicConfig(format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s',
                    level=logging.DEBUG)
logging.debug('debug 信息')
logging.info('info 信息')
logging.warning('warning 信息')
logging.error('error 信息')
logging.critical('critial 信息')
~~~

