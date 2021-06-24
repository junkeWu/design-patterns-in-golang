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
6. 适配器模式
7. 桥接模式
8. 代理模式
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


# 二、单例模式
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

# 三、建造型模式
  建造型模式用于创建过程稳定，但配置多变的对象。在《设计模式》一书中的定义是：将一个复杂的构建与其表示相分离，使得同样的构建过程可以创造不同的表示。 经典的【建造者-指挥者】模式现在已经不太常用了，现在建造者模式主要用来通过链式调用生成不同的配置。
  比如我们要制作一杯珍珠奶茶，他的制作过程是稳定的，除了必须要知道奶茶的种类和规格外，是否加珍珠和是否加冰是可选的。
  建造者模式中有一下几个角色需要我们熟悉：
Product: 这是我们要创建的复杂对象(一般都是很复杂的对象才需要使用建造者模式)。
Builder: 抽象的一个东西， 主要是用来规范我们的建造者的。
ConcreteBuilder: 具体的Builder实现， 这是今天的重点，主要用来根据不用的业务来完成要创建对象的组件的创建。
Director: 这个的作用是规范复杂对象的创建流程。

# 四、原型模式
  原型模式：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。举个例子，比如有一天，周杰伦到奶茶店点了一份不加冰的原味奶茶，你说我是周杰伦的忠实粉，我也要一份跟周杰伦一样的。
  原型模式用于创建重复的对象，同时又能保证性能。一般是通过实现 Clone()接口来实现。
  <h3>五大要素<h3>
模式名称：原型模式
  目的（What）：快速创建一个需要创建的对象的克隆，由此返回一个新的内存实例
  解决的问题（Why）：当直接创建对象的代价比较大时，使用这种模式
  解决方案（How）：实现Clone()方法，类型实现这个方法，快速地生成和原型对象一样的实例。比如细胞分类等等。
  解决效果：
  优点：
    性能提高，当创建对象需要一系列繁琐操作的时候，使用原型模式可以提高一定的性能。
  缺点：
    对象必须实现Clone()方法.

# 结构型模式

# 一、适配器模式
  将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作。
  适配的意思是适应、匹配。通俗地讲，适配器模式适用于 有相关性但不兼容的结构，源接口通过一个中间件转换后才可以适用于目标接口，这个转换过程就是适配，这个中间件就称之为适配器。
# 二、桥接模式
桥接模式最主要的特点就是松耦合。这里举一个比较好懂的例子。
我们每个人的家里都有一些电器（比如灯泡和风扇），每个电器都是由开关控制，开关也有不同的种类（圆形的，方形的，或者是很小的那种开关），它们之间通过线路连接，可以把线路理解为它们之间的桥梁。在任何情况下，我们可以在不更改线路（桥梁）的情况下，对电器或者开关做更换，比如我们可以在不更换开关的情况下，换一个新的灯泡，也可以在不更换灯泡的情况下，更换开关（比如从方形开关换成圆形开关），甚至可以把灯泡和风扇互换，只要我们乐意。
因此对于桥接模式而言，当不同的事物被联系到一起时，可以更换它们其中的任意一个而不受影响。在上面的例子中，电器是一个抽象类（灯泡和风扇是它的具体实现类），开关是另一个抽象类（圆形开关和方形开关是它的具体实现类）。这两个抽象类通过桥接的形式连接在一起（线路），在这种情况下，我们可以替换抽象类中的任意一个而让整体系统不受到影响。
  <h3>五大要素<h3>
来看看桥接模式的五大要素：

模式名称：桥接模式
  目的（What）：将抽象部分与实现部分分离，使它们都可以独立的变化
  解决的问题（Why）：在有多种可能会变化的情况下，用继承会造成类爆炸问题，扩展起来不灵活。当实现系统可能有多个角度分类（多个抽象类），每一种角度会有不同的变化。
  解决方案（How）：把每个角度的分类分离出来，让它们独立变化，减少它们之间的耦合。
  解决效果：
  优点：
    抽象和实现的分离
    优秀的扩展能力
    实现细节对客户透明
  缺点：
    桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系（桥接）建立在抽象层，要求开发者针对抽象进行设计与编程。
# 三、组合模式
组合模式和桥接模式的组合完全不一样。组合模式用于 整体与部分的结构，当整体与部分有相似的结构，在操作时可以被一致对待时，就可以使用组合模式。
文件夹和子文件夹的关系：文件夹中可以存放文件，也可以新建文件夹，子文件夹也一样。
总公司子公司的关系：总公司可以设立部门，也可以设立分公司，子公司也一样。
组合模式最主要的功能就是让用户可以一致对待整体和部分结构,将两者都作为一个相同的组件，所以我们先新建一个抽象的组件类.
组合模式,指的是就是用接口做抽象,在这个接口内部组合自己,形成一个树型结构.
像组织架构,课表,等等可以形成树状的实体,都可以用组合模式来做一个实现,这样我们可以不断的演化节点类型而对总体的节点不用修改,不断的扩展这棵树的体积而不用动旧代码。
· 优点
高层模块调用简单
节点自由增加
· 缺点
使用组合模式时，其叶子和树枝的声明都是实现类。而不是接口，违反了依赖倒置原则。

# 结构型模式
# 一、观察者模式
观察者模式是一种行为设计模式， 允许你定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象。
拥有一些值得关注的状态的对象通常被称为目标， 由于它要将自身的状态改变通知给其他对象， 我们也将其称为发布者 （publisher）。 所有希望关注发布者状态变化的其他对象被称为订阅者 （subscribers）。

观察者模式建议你为发布者类添加订阅机制， 让每个对象都能订阅或取消订阅发布者事件流。 不要害怕！ 这并不像听上去那么复杂。 实际上， 该机制包括 1） 一个用于存储订阅者对象引用的列表成员变量； 2） 几个用于添加或删除该列表中订阅者的公有方法。

· 观察者结构
1. 发布者 （Publisher） 会向其他对象发送值得关注的事件。 事件会在发布者自身状态改变或执行特定行为后发生。 发布者中包含一个允许新订阅者加入和当前订阅者离开列表的订阅构架。
2. 当新事件发生时， 发送者会遍历订阅列表并调用每个订阅者对象的通知方法。 该方法是在订阅者接口中声明的。
3. 订阅者 （Subscriber） 接口声明了通知接口。 在绝大多数情况下， 该接口仅包含一个 update更新方法。 该方法可以拥有多个参数， 使发布者能在更新时传递事件的详细信息。
4. 具体订阅者 （Concrete Subscribers） 可以执行一些操作来回应发布者的通知。 所有具体订阅者类都实现了同样的接口， 因此发布者不需要与具体类相耦合。
5. 订阅者通常需要一些上下文信息来正确地处理更新。 因此， 发布者通常会将一些上下文数据作为通知方法的参数进行传递。 发布者也可将自身作为参数进行传递， 使订阅者直接获取所需的数据。