# design-pattern-go
this is a design pattern implement by go language

# factory method
工厂方法的目的，是给一类对象，抽象共同的方法，并且使用工厂方法来构建，实现使用者和具体的类的隔离，从而保持灵活性。
缺陷：工厂方法的类型构建比较死板

# abstract factory
抽象工厂的目的，比较工厂方法，多了一个系列产品的概率，所以由每个具体的工厂生产一个系列的产品。这个关注的是一套产品的组合抽象。使用者和具体的物品也是隔离的。
缺陷：会多出很多类，风格的抽象也比较考究
