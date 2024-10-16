# design-pattern-go
this is a design pattern implement by go language

# factory method
工厂方法的目的，是给一类对象，抽象共同的方法，并且使用工厂方法来构建，实现使用者和具体的类的隔离，从而保持灵活性。
缺陷：工厂方法的类型构建比较死板

# abstract factory
抽象工厂的目的，比较工厂方法，多了一个系列产品的概率，所以由每个具体的工厂生产一个系列的产品。这个关注的是一套产品的组合抽象。使用者和具体的物品也是隔离的。
缺陷：会多出很多类，风格的抽象也比较考究


# builder
建造者模式，解决的是对象的构造方法太复杂的问题。如果构造参数过多，或者构造函数太多，都有不够优雅的问题。对比前两个模式，这个更新是关注的对象垂直的复杂性。工厂模式关注的水平层面的复杂性。


# prototype
原型模式，或者说克隆模式。为了让调用者不了解目标类的情况下，复制对象。


# adaptor
适配器是一个简单的结构模式，用于解决存量代码的不兼容问题。使用中间层做双向转化。

# bridge
桥接模式，是在面对复杂的对象结构是，使用分层，或者分类思想，给对象分出抽象层，和具体实现层，然后用组合的模式，构建对象间的桥梁。这里注意分清主次，由谁来引用谁。所以桥接模式也是一种组合。

# composite
组合模式，你的对象一定要是树状结构才可以使用。让你能提纲挈领的使用这些复杂的对象。

# decorator
装饰器模式是给一个对象，提供更多的使用表现的模式。装饰器的方法执行顺序由引用定义的顺序实现。优势，减少了定时各种特殊类的繁琐。

# facade
外观模式，是一个简单的模式，目的是给一复杂的模块，提供一个简单的统一的调用接口，向调用者屏蔽复杂性。

# flyweight
享元模式，对象要区分可变和不可变，然后复用不可变的属性为享元。从而实现性能的优化。

# proxy
代理模式，对于复杂的对象，或者不容易改变的外部对象，使用代理模式，扩展他的能力。并且管理生命周期。

# chain of responsibility
责任链模式，是典型的流水线处理一个业务逻辑，并且实现流程的可插拔。实际代码中应用

# command
命令模式，是将复杂的指令封装成一个统一的对象，这样便于前后端的分离。调用者不用关心指令的复杂性。并且能够实现指令执行历史和回滚等功能。

# iterator
迭代器模式，用户遍历集合对象，而不用关心对象的具体实现。

# mediator
中介模式，是一种行为模式，如果对象间有复杂的调用关系，可以尝试使用中介模式，将对象隔离开。解除他们的耦合。但是要防止上帝对象的出现

# memento
备忘录模式，一种不用关系详细的复制的拷贝模式，原始对象自己要实现备份。然后备份的形态要统一，并且注意防止修改。

# observer
观察者模式，即发布订阅，订阅者要有统一的修改接口，发布者拥有对订阅列表的管理接口，以及通知机制。

# state
状态模式， 在计算机中有限状态自动机是经常用到的。但是复杂的状态切换和判断，会导致代码的腐坏。所以可以将每一状态做成一个类，然后状态类在指定的操作里互相转化。

# strtegy
策略模式，针对具体算法的抽象。

# tempalte 
模版方法，和工厂方法类似，工厂方法抽象的大的构建过程。模版方法是执行流程的抽象， 并且依赖子类的实现，实现特定的执行。

# visitor
访问者模式，将对象和其具体的算法分离的好设计模式，唯一缺点是在对象中需要添加 accept 方法。但是好处是，后续修改都不用再变化，只需要切换不同的访问者即可。