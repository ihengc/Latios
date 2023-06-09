# I/O缓冲

需要理解为什么需要I/O缓冲。I/O到底做了哪些事？

## 读缓冲

读缓冲结构：

* 数据缓冲区。
* 文件描述符（或者能进行I/O操作的对象）可由外部传入。
* 读位置。
* 写位置。

默认缓冲区大小为4096字节。数据缓冲区可以理解为一个环形字节数组。环形数组将被分为两个区间：

1. 读位置到写位置，表示未被读取的数据区间。
2. 写位置到读位置，表示还可以存放数据的区间。

判断读缓冲中是否还有未读的数据：`写位置减读位置是否大于0`。

### 读取指定字节数的数据

若指定字节数n大于缓冲区的大小，我们没必要将读取的数据存入缓冲区中而后从缓冲区中取出数据，我们可以直接进行I/O操作，读取n字节的数据。

否则读取缓冲区大小的数据（处理系统调用错误），读位置移动。

### 读满缓冲区

这是一个很常见的操作，正常情况下读缓冲在执行系统调用时，总是希望一次读满整个缓冲区。

如何读满缓冲区？

首先需要判断当前读缓冲区的状态，判断缓冲区中有未读取的数据（读位置到写位置的区间）。

如何处理出现无法读取到数据的情况？

进行有限次数的重试。

### 读取到指定分隔符

这里涉及到如何查找分隔符的问题。

### 跳过指定字节数

## 写缓冲
