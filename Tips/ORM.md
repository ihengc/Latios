## 数据库种类

通常`ORM`框架为了支持多种数据库，会抽象一个“方言（dialect）”接口。

此接口的抽象很自然，因为不同数据库的数据类型，表定义的方式，约束等都不相同，通过`dialect`接口可以很好的屏蔽掉这些问题。

`dialect`接口构成：

* 获取当前是哪种数据库。
*