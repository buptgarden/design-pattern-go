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