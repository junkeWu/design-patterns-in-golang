<h4>design-patterns-in-golang</h4>

<h3>23种设计模式, go语言实现版</h3>

# 设计模式的六大原则
设计模式的时间丰富多彩，比如生产一个个产品的工厂模式，衔接俩个不相关接口的适配器模式，用不用方式做同一件事的策略模式、构建步骤稳定、根据构建过程的不同配置构建不同对象的建造者模式。

面向对象结合设计模式，才能真正体会到程序变得可维护、可复用、可扩展性、灵活性好。设计模式对于程序员而言并不陌生，每人程序员在编程时或多或少地接触设计模式。无论是在大型程序的架构中，亦或是在源码的学习中，设计模式都扮演着非常重要的角色。

<h2>设计模式基于六大原则</h2>

```
. 开闭原则：一个软件实体如类、模块和函数应该对修改封闭，对扩展开放。
. 单一职责原则：一个类只做一件事，一个类应该只有一个引起它修改的原因。
. 里氏替换原则：子类应该可以完全替换父类。也就是说在使用继承时，只扩展新功能，而不要破坏父类原有的功能。
. 依赖倒置原则： 细节应该依赖于抽象，抽象不应依赖于细节。把抽象层放在程序设计的高层，并保持稳定，程序的细节变化由低层的实现层来完成。
. 迪米特法则：又名「最少知道原则」，一个类不应知道自己操作的类的细节，换言之，只和朋友谈话，不和朋友的朋友谈话。
. 接口隔离原则：客户端不应依赖它不需要的接口。如果一个接口在实现时，部分方法由于冗余被客户端空实现，则应该将接口拆分，让实现类只需依赖自己需要的接口方法。
```
所有的设计模式都是为了程序能更好的满足这六大原则。设计模式一共有 23 种。

# 创建型模式分类 5
1. 单例模式
2. 原型模式
3. 工厂方法模式
4. 抽象工厂模式
5. 建造者模式

# 结构型模式概述 7
6. 代理模式
7. 适配器模式
8. 桥接模式
9. 装饰模式
10. 外观模式
11. 享元模式
12. 组合模式

# 行为型模式概述 11
13. 模板方法模式
14. 策略模式
15. 命令模式
16. 责任链模式
17. 状态模式
18. 观察者模式
19. 中介者模式
20. 迭代器模式
21. 访问者模式
22. 备忘录模式
23. 解释器模式

# 一、工厂模式
  在平时编程中，构建对象最常用的方式是 new 一个对象。乍一看这种做法没什么不好，而实际上这也属于一种硬编码。在平时编程中，构建对象最常用的方式是 new 一个对象。乍一看这种做法没什么不好，而实际上这也属于一种硬编码。

 <h3>1.1简单工厂模式</h3>
 <h3>1.2工厂方法模式</h3>
 <h3>1.3抽象工厂模式</h3>
代码路径：https://github.com/junkeWu/design-patterns-in-golang/tree/main/creational

# 一、工厂模式
  在平时编程中，构建对象最常用的方式是 new 一个对象。乍一看这种做法没什么不好，而实际上这也属于一种硬编码。在平时编程中，构建对象最常用的方式是 new 一个对象。乍一看这种做法没什么不好，而实际上这也属于一种硬编码。

<h3>1.1简单工厂模式</h3>
  总而言之，简单工厂模式就是让一个工厂类承担构建所有对象的职责。调用者需要什么产品，让工厂生产出来即可。它的弊端也显而易见：
* 一是如果需要生产的产品过多，此模式会导致工厂类过于庞大，承担过多的职责，变成超级类。当苹果生产过程需要修改时，要来修改此工厂。梨子生产过程需要修改时，也要来修改此工厂。也就是说这个类不止一个引起修改的原因。违背了单一职责原则。
* 二是当要生产新的产品时，必须在工厂类中添加新的分支。而开闭原则告诉我们：类应该对修改封闭。我们希望在添加新功能时，只需增加新的类，而不是修改既有的类，所以这就违背了开闭原则。

<h3>1.2工厂方法模式</h3>
  为了解决简单工厂模式的这两个弊端，工厂方法模式应运而生，它规定每个产品都有一个专属工厂。比如苹果有专属的苹果工厂，梨子有专属的梨子工厂. 
* 当生产的产品种类越来越多时，工厂类不会变成超级类。工厂类会越来越多，保持灵活。不会越来越大、变得臃肿。如果苹果的生产过程需要修改时，只需修改苹果工厂。梨子的生产过程需要修改时，只需修改梨子工厂。符合单一职责原则。 
* 当需要生产新的产品时，无需更改既有的工厂，只需要添加新的工厂即可。保持了面向对象的可扩展性，符合开闭原则。

<h3>1.3抽象工厂模式</h3>
  工厂方法模式可以进一步优化，提取出工厂接口，然后苹果工厂和梨子工厂都实现此接口. 
* 实际上抽象工厂模式主要用于替换一系列方法。
例如将程序中的 SQL Server 数据库整个替换为 Access 数据库，使用抽象方法模式的话，只需在 IFactory 接口中定义好增删改查四个方法，
让 SQLFactory 和 AccessFactory 实现此接口，调用时直接使用 IFactory 中的抽象方法即可，调用者无需知道使用的什么数据库，我们就可以非常方便的整个替换程序的数据库，并且让客户端毫不知情。


# 一、单例模式
  单例模式非常常见，某个对象全局只需要一个实例时，就可以使用单例模式。它的优点也显而易见：
* 它能够避免对象重复创建，节约空间并提升效率 
* 避免由于操作不同实例导致的逻辑错误
  单例模式有两种实现方式：饿汉式和懒汉式。

<h3>2.1懒汉式</h3>
缺点：非线程安全。当正在创建时，有线程来访问在判断有无实例存在后，创建实例前，此时就会再创建，单例类就会有多个实例了。
<h3>2.2懒汉式加锁</h3>
缺点：缺点：虽然解决并发的问题，但每次加锁是要付出代价的
<h3>2.3饿汉式</h3>
缺点：如果singleton创建初始化比较复杂耗时时，加载时间会延长。
<h3>2.4双重锁</h3>
优点：减少加锁的几率，提高性能
<h3>2.5sync.Once实现</h3>
优点：golang提供的方法的，只会执行一次。