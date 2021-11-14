import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;

/**
 * @author linhaizeng
 * @Description TODO
 * @date 2021/11/10-7:31 下午
 */
public class ThreadPoolTest {
    public static void main(String[] args) {
        ExecutorService es = Executors.newScheduledThreadPool(3);
        ((ScheduledExecutorService) es).scheduleAtFixedRate(new Runnable() {
            @Override
            public void run() {
                loge("time:");
            }
        }, 0, 40, TimeUnit.MILLISECONDS);//0表示首次执行任务的延迟时间，40表示每次执行任务的间隔时间，TimeUnit.MILLISECONDS执行的时间间隔数值单位



    }
}
