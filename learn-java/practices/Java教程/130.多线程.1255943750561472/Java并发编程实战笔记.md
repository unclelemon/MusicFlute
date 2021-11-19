# JUC (java.util.concurrent) 学习
# 第二部分 结构化并发应用程序
## 第6章 任务执行
### 6.2 Executor 框架
Executor 是个简单的接口
~~~java
public interface Executor {
    void execute(Runnable command);
}
~~~
### 6.2.3 线程池
