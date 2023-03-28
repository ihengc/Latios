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

### 读取指定字节数的数据

若指定字节数n大于缓冲区的大小，我们没必要将读取的数据存入缓冲区中而后从缓冲区中取出数据，我们可以直接进行I/O操作，读取n字节的数据。

否则读取缓冲区大小的数据（处理系统调用错误），读位置移动。