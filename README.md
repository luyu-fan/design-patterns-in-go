# design-patterns-in-go

设计模式的Go语言练习实践

## About

设计模式是软件开发过程中构建可扩展、可重用、可维护程序的经验沉淀。设计模式不是在直接编写功能代码本身，而是通过
对构建过程、对象组合已经行为表示进行抽象和归类，总结出各类场景下的设计规范，利用这些规范编写出更好的代码。

## 创建型模式

规范了如何在系统中创建对象，并利用这些类及对象对系统进行参数化。

- :star:[工厂方法](./1-simplefactory/readme.md)
- :star:[抽象工厂](./2-abstractfactory/readme.md)
- :star:[单例](./5-singleton/readme.md)
- [建造者](./3-builder/readme.md)
- [原型](./4-prototype/readmd.md)

## 结构型模式

- :star:[适配器](./6-adapter/readme.md)
- :star:[桥接](./7-bridge/readme.md)
- :star:[组合](./8-composite/readme.md)
- [装饰器](./9-decorator/readme.md)
- :star:[外观](./10-facade/readme.md)
- [享元](./11-flyweight/readme.md)
- [代理](./12-proxy/readme.md)

## 行为型模式

行为型模式主要考虑到对变化以及通信的封装。当一个程序经常性发生变化时，这个模式就定义一个封装这个变化方面的对象，
当程序的其它部分依赖这个方面时便可以通过这个对象进行协作。

一些模式总是引入一些被用作参数的对象。对象之间也存在通信，它们都会根据具体问题而被分为封装和分布处理。

- [责任链](./13-chain/readme.md)
- :star:[命令](./14-command/readme.md)
- [解释器](./15-interpreter/readme.md)
- :star:[迭代器](./16-iterator/readme.md)
- :star:[中介者](./17-mediator/readme.md)
- [备忘录](./18-memento/readme.md)
- :star:[观察者](./19-observer/readme.md)
- :star:[状态](./20-state/readme.md)
- :star:[策略](./21-strategy/readme.md)
- [模板](./22-template/readme.md)
- [访问者](./23-visitor/readme.md)
