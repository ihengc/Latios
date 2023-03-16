package loader

/*
* @author: Chen Chiheng
* @date: 2023/3/13 0013 17:08
* @description:
**/

// 加载数据
// 同步加载，异步加载。
// 异步加载，使用数据库连接池，并发量为数据库连接池的数量。若多各加载器共用一个数据库，则会出现数据库连接不够的情况。
// 加载耗时：
// 网络IO，磁盘IO，数据转换。。

// Loader 数据加载接口。从指定的数据源加载（获取数据）。数据源可以是磁盘文件，数据库，网络等。
type Loader interface {
	// LimitedLoad 加载数据，加载出数据最大为传入字节数组的长度。若数据超过
	LimitedLoad([]byte) (n int, err error)
}

// Converter 数据转换接口，负责数据源数据格式与Golang结构之间的转换。
type Converter interface {
}
