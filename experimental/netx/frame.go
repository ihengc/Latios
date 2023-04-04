package netx

/*
* @author: Chen Chiheng
* @date: 2023/4/3 0003 15:48
* @description:
**/

// FrameType 数据帧类型。
type FrameType uint8

const (
	FrameSYN    FrameType = 0x0 // FrameSYN 用来发起连接。
	FrameACK    FrameType = 0x1 // FrameACK 回应发起连接。
	FramePing   FrameType = 0x2 // FramePing 心跳数据帧。
	FrameData   FrameType = 0x3 // FrameData 数据帧。
	FramePush   FrameType = 0x4 // FramePush 服务主动推送的数据帧。
	FrameGoAway FrameType = 0x5 // FrameGoAway 服务主动断开连接前发送此数据帧。
)

// FrameHeader 数据帧头部。
type FrameHeader struct {
	// Type 数据帧类型。
	Type FrameType
	// Length 数据帧总长度。通过固定头部数据帧的长度，使用此字段可以计算出数据帧中正文的大小。
	Length uint32
	// IsCompressed 是否压缩，可选（默认不开启）。通常会使用Protobuf进行数据的序列化，
	// Protobuf的数据比较紧凑，我们应该尽量保证单次传输的数据短小。尽量避免压缩开销。
	IsCompressed bool
	// ID 数据帧ID。在使用UDP协议（非面向连接的协议）需要此字段。
	ID uint32
}

// DataFrame 数据帧。
type DataFrame struct {
	FrameHeader
	payload []byte
}

// PingFrame 心跳帧。
type PingFrame struct {
	FrameHeader
	payload []byte
}
