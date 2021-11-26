## 自动装配入门
@SpringBootApplication注解是一个聚合注解，其中主要包含了
- @SpringBootConfiguration
- @ComponentScan
- @EnableAutoConfiguration # 负责Spring Boot的自动装配机制
尽管@CompenentScan 仅关注@Compent，但是@SpringBootConfiguration派生自@Configuration，而@Configuration又派生自@Compent，所以@SpringBootConfiguration也可以被@CompenentScan扫描到。
它们的层次关系如下：
- @Compent
    - @Configuration
        - SpringBootConfiguration

### 注解解释

#### @ConfigurationProperties

取代传统xml配置文件中出现的bean的value值，DruidDataSource类中的属性自动赋值。

使用prefix，能够让类DataSource读取到配置文件中的值。

作用于构建Bean的方法上和作用于类上同理。如下为作用于构建Bean的方法上。

[参考博客](https://www.cnblogs.com/duanxz/p/4520571.html)

~~~properties
#数据源
spring.datasource.druid.write.url=jdbc:mysql://localhost:3306/jpa
spring.datasource.druid.write.username=root
spring.datasource.druid.write.password=1
spring.datasource.druid.write.driver-class-name=com.mysql.jdbc.Driver

spring.datasource.druid.read.url=jdbc:mysql://localhost:3306/jpa
spring.datasource.druid.read.username=root
spring.datasource.druid.read.password=1
spring.datasource.druid.read.driver-class-name=com.mysql.jdbc.Driver
~~~

~~~java
@Configuration
public class DruidDataSourceConfig {
    /**
     * DataSource 配置
     * @return
     */
    @ConfigurationProperties(prefix = "spring.datasource.druid.read")
    @Bean(name = "readDruidDataSource")
    public DataSource readDruidDataSource() {
        return new DruidDataSource();
    }

    /**
     * DataSource 配置
     * @return
     */
    @ConfigurationProperties(prefix = "spring.datasource.druid.write")
    @Bean(name = "writeDruidDataSource")
    @Primary
    public DataSource writeDruidDataSource() {
        return new DruidDataSource();
    }
}
~~~

#### @EnableConfigurationProperties

用于注入ConfigurationProperties。
