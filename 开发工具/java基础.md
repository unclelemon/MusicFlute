## JAVA基础
### SimpleDateFormat
![SimpleDateFormat格式](../images/image.png)
SimpleDateFormat sdf=new SimpleDateFormat("yyyy-MM-dd HH:mm:ss"); 
转换的格式：yyyy是完整的西元年，MM是月份，dd是日期， 至于HH:mm:ss.SSS 时分秒

例如MM是月份，mm是分；
HH是24小时制，而hh是12小时制 
ss是秒，SSS是毫秒
DD代表 day of year  今年的第N天，dd才是这个月的多少天。

## JUC高并发
### ThreadLocal 介绍

[参考网站](https://www.cnblogs.com/fsmly/p/11020641.html)

多线程访问同一个共享变量的时候容易出现并发问题，特别是多个线程对一个变量进行写入的时候，为了保证线程安全，一般使用者在访问共享变量的时候需要进行额外的同步措施才能保证线程安全性。ThreadLocal是除了加锁这种同步方式之外的，另一种保证线程安全的手段。当我们创建一个变量的时候，每个线程都有着一份该变量，并且都操作自己的变量，不去干涉其他线程，可规避多线程操作同一个变量造成的不安全性。

### ThreadLocal的简单实用

~~~java
package test;

public class ThreadLocalTest {

    static ThreadLocal<String> localVar = new ThreadLocal<>();

    static void print(String str) {
        //打印当前线程中本地内存中本地变量的值
        System.out.println(str + " :" + localVar.get());
        //清除本地内存中的本地变量
        localVar.remove();
    }

    public static void main(String[] args) {
        Thread t1  = new Thread(new Runnable() {
            @Override
            public void run() {
                //设置线程1中本地变量的值
                localVar.set("localVar1");
                //调用打印方法
                print("thread1");
                //打印本地变量
                System.out.println("after remove : " + localVar.get());
            }
        });

        Thread t2  = new Thread(new Runnable() {
            @Override
            public void run() {
                //设置线程1中本地变量的值
                localVar.set("localVar2");
                //调用打印方法
                print("thread2");
                //打印本地变量
                System.out.println("after remove : " + localVar.get());
            }
        });

        t1.start();
        t2.start();
    }
}
~~~
### 枚举类
往enum中添加方法。其中的RED都是enum创建出来的具体事例。
~~~
public enum Color {
    RED("红色", 1), GREEN("绿色", 2), BLANK("白色", 3), YELLO("黄色", 4);
    // 成员变量
    private String name;
    private int index;

    // 构造方法
    private Color(String name, int index) {
        this.name = name;
        this.index = index;
    }

    // 普通方法
    public static String getName(int index) {
        for (Color c : Color.values()) {
        if (c.getIndex() == index) {
            return c.name;
        }
        }
        return null;
    }

    // get set 方法
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getIndex() {
        return index;
    }

    public void setIndex(int index) {
        this.index = index;
    }
    }
~~~