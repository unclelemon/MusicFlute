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
~~~java
// 创建一个固定长度的线程池，每当提交一个任务时就创建一个线程，直到达到线程池的最大数量，这时候线程池的规模不再变化
    // 如果某个线程由于发生了未预期的Exception而结束，那么线程池会补充一个新的线程
    ExecutorService es1 = Executors.newFixedThreadPool(NTHREADS);
    
    // 创建一个可缓存的线程池，如果线程池的当前规模超过了处理需求时候，那么将回收空闲的线程，而当需求增加时，则可以添加新的线程
    // 线程池的规模不存在任何限制
    ExecutorService es2 = Executors.newCachedThreadPool();
    
    // 创建单一线程
    ExecutorService es3 = Executors.newSingleThreadExecutor();

    // 创建固定长度的线程池，而且以延迟或者定时的方式执行任务。
    ExecutorService es4 = Executors.newScheduledThreadPool(NTHREADS);
~~~