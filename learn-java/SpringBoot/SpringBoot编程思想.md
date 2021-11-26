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
