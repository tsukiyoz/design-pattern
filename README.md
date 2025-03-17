# Go 版设计模式（58 种）

> 云原生 AI 实战营项目之一，更多精彩项目见：[云原生 AI 实战营](https://konglingfei.com/)。

本仓库是 [《Go 设计模式 61 讲》课程](https://konglingfei.com/cloudai/catalog/design-pattern.html) 的源码仓库。里面介绍了 58 种常见的设计模式，及 Go 版代码实现。

如果你想全面学习这些设计模式，欢迎加入 [云原生 AI 实战营](https://konglingfei.com)。里面不仅包含了 《Go 设计模式 61 讲》课程，还包含了大量其他高质量的 Go、云原生、AI Infra 课程。

<br/>

在 《Go 设计模式 61 讲》课程中一共介绍了 58 种设计模式，这些设计模式列表如下。

## 创建型模式（Creational Patterns）

| 模式名 | 英文名 | 状态 | 
| --- | --- | --- |
|[简单工厂模式](./creational/simplefactory)|Simple Factory|✔|
|[工厂方法模式](./creational/factorymethod)|Factory Method|✔|
|[抽象工厂模式](./creational/abstractfactory)|Abstract Factory|✔|
|[建造者模式](./creational/builder)|Builder|✔|
|[原型模式](./creational/prototype)|Prototype|✔|
|[单例模式](./creational/singleton)|Singleton|✔|
|[New 模式](./creational/new)|New|✔|
|[函数选项模式](./creational/functionaloption)|Functional Options|✔|
|[对象池模式](./creational/objectpool)|Object Pool|✔|

## 行为型模式（Behavioral Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[中介者模式](./behavioral/mediator)|Mediator|✔|
|[观察者模式](./behavioral/observer)|Observer|✔|
|[命令模式](./behavioral/command)|Command|✔|
|[迭代器模式](./behavioral/iterator)|Iterator|✔|
|[模版方法模式](./behavioral/templatemethod)|Template Method|✔|
|[策略模式](./behavioral/strategy)|Strategy|✔|
|[状态模式](./behavioral/state)|State|✔|
|[备忘录模式](./behavioral/memento)|Memento|✔|
|[解释器模式](./behavioral/interpreter)|Interpreter|✔|
|[责任链模式](./behavioral/chainofresponsibility)|Chain of Responsibility|✔|
|[访问者模式](./behavioral/observer)|Visitor|✔|
|[注册表模式](./behavioral/registry)|Registry|✔|
|[上下文模式](./behavioral/context)|Context|✔|

## 结构型模式（Structural Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[外观模式](./structural/facade)|Facade|✔|
|[适配器模式](./structural/adapter)|Adapter|✔|
|[代理模式](./structural/proxy)|Proxy|✔|
|[组合模式](./structural/composite)|Composite|✔|
|[享元模式](./structural/flyweight)|Flyweight|✔|
|[装饰器模式](./structural/decorator)|Decorator|✔|
|[桥接模式](./structural/adapter)|Bridge|✔|

## 同步模式（Synchronization Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[条件变量模式](./synchronization/condition)|Condition Variable|✔|
|[互斥锁模式](./synchronization/lockmutex)|Lock/Mutex|✔|
|[监视器模式](./synchronization/monitor)|Monitor|✔|
|[读写锁模式](./synchronization/readwritelock)|Read-Write Lock|✔|
|[信号量模式](./synchronization/semaphore)|Semaphore|✔|

## 并发模式（Concurrency Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[屏障模式](./concurrency/barrier)|N-Barrier|✔|
|[有界并行性模式](./concurrency/boundedparallelism)|Bounded Parallelism|✔|
|[广播模式](./concurrency/broadcast)|Broadcast|✔|
|[协程模式](./concurrency/coroutines)|Coroutines|✔|
|[生成器模式](./concurrency/generator)|Generator|✔|
|[反应器模式](./concurrency/reactor)|Reactor|✔|
|[并行模式](./concurrency/parallelism)|Parallelism|✔|
|[生产者消费者模式](./concurrency/producerconsumer)|Producer Consumer|✔|
|[批处理模式](./concurrency/batcher)|Batch Processing|✔|

## 消息传递模式（Messaging Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[扇入模式](./messaging/fanin)|Fan-In|✔|
|[扇出模式](./messaging/fanout)|Fan-Out|✔|
|[未来与承诺模式](./messaging/futurespromises)|Futures & Promises|✔|
|[发布订阅模式](./messaging/pubsub)|Publish/Subscribe|✔|
|[推模式与拉模式](./messaging/pushpull)|Push & Pull|✔|

## 稳定型模式（Stability Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[隔离模式](./stability/bulkhead)|Bulkheads|✔|
|[断路器模式](./stability/circuit)|Circuit-Breaker|✔|
|[截止期限模式](./stability/deadline)|Deadline|✔|
|[快速失败模式](./stability/failfast)|Fail-Fast|✔|
|[握手模式](./stability/handshaking)|Handshaking|✔|
|[稳态模式](./stability/steadystate)|Steady-State|✔|
|[限流模式](./stability/ratelimiting)|Rate Limiting|✔|
|[重试模式](./stability/retrier)|Retrier|✔|

## 分析模式（Profiling Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[计时函数模式](./profiling/timing)|Timing Functions|✔|

## 反模式（Anti-Patterns）

| 模式名 | 英文名 | 状态 |
| --- | --- | --- |
|[串联故障模式](./anti/cascadingfailures)|Cascading Failures|✔|

## 参考资料

| 推荐顺序| 项目名 | 参考星级 | 参考状态 |
| --- | --- | --- | --- |
| 1 | [crazybber/go-pattern-examples](https://github.com/crazybber/go-pattern-examples) | ★★★★★ |✔|
| 2 | [crazybber/awesome-patterns](https://github.com/crazybber/awesome-patterns) | ★★★★★ |✔|
| 3 | [tmrts/go-patterns](https://github.com/tmrts/go-patterns) | ★★★★ |✔|
| 4 | [senghoo/golang-design-pattern](https://github.com/senghoo/golang-design-pattern) | ★★★ |✔|
| 5 | [lee501/go-patterns](https://github.com/lee501/go-patterns) | ★★★ |✔|
| 6 | [mohuishou/go-design-pattern](https://github.com/mohuishou/go-design-pattern) | ★★★ |✔|

## Contacts

<img src="./images/three-code.png" alt="" width="900" />
