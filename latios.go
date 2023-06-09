package Latios

/*
* @author: Chen Chiheng
* @date: 2023/3/13 0013 14:33
* @description:
**/

//	数据库驱动：
//	建立数据库链接：
//
//	“准备好被执行的语句”：
//		用户出入的SQL语句会经过处理生成为“准备好被执行的语句”。可以直接在此结构上执行SQL语句。
//		关闭此结构就相当于关闭当前的数据库链接。
//	事务:
//		提交与回滚。
//	Ping:
//		链接探活。
//	行：
//		特指一个SELECT语句的执行结果。返回当前表的全部列名。查询结果通常是一个集合，迭代集合是最频繁的操作。
//	语句执行的结果：
//		对于INSERT,DELETE,UPDATE语句返回影响的行数，INSERT还可以额外返回插入后最新的主键ID（若有）。

// 数据库连接池：

// 内存对齐：
//		为什么需要内存对齐？
//		对齐系数
//		对齐规则
//		结构体中字段合理的排布可以减少结构体的内存消耗。
