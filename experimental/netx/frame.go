package netx

/*
 * @author: Chen Chiheng
 * @date: 2023/4/3 22:22
 * @description:
 */

// 本文件对应用层的数据包进行抽象，封装。

type Frame interface {
}

type FrameReaderWriter struct {
	// 读缓冲
	// 写缓冲
	// 对于头部固定的数据帧，可以直接初始化一个头部缓冲。
}
