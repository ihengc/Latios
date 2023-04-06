package netx

import "io"

/*
* @author: Chen Chiheng
* @date: 2023/4/3 0003 15:48
* @description:
**/

const (
	frameHeaderSize = 10
	maxFrameSize    = 1<<24 - 1
)

// FrameType 数据帧类型。
type FrameType uint8

const (
	FrameSYN     FrameType = 0x0 // FrameSYN 用来发起连接。
	FrameACK     FrameType = 0x1 // FrameACK 回应发起连接。
	FramePing    FrameType = 0x2 // FramePing 心跳数据帧。
	FrameData    FrameType = 0x3 // FrameData 数据帧。
	FramePush    FrameType = 0x4 // FramePush 服务主动推送的数据帧。
	FrameGoAway  FrameType = 0x5 // FrameGoAway 服务主动断开连接前发送此数据帧。
	FrameKickOut FrameType = 0x6 // FrameKickOut 踢出用户数据帧。
)

// Flags 帧标志位类型。
type Flags uint8

const (
	FlagPingAck    Flags = 0x1 // FlagPingAck ping回应帧。
	FlagFragmented Flags = 0x2 // FlagFragmented 是否是分片帧。
)

type Frame interface {
	Header() FrameHeader
	Payload() []byte
}

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
	// Timestamp 时间戳，记录包发送时的时间。
	Timestamp int64
	// Flag 帧标志位。
	Flag Flags
	// FragmentOffSet 分片偏移。
	FragmentOffSet uint32
}

// DataFrame 数据帧。
type DataFrame struct {
	FrameHeader
	payload []byte
}

// PingFrame 心跳帧。
type PingFrame struct {
	FrameHeader
}

// FrameReaderWriter 帧读写。
type FrameReaderWriter struct {
	r io.Reader
	// headerBuf 若header长度固定可以设置header缓冲。
	headerBuf [frameHeaderSize]byte
	// getReadBuf 获取读缓冲。
	// 若需求的缓冲区大小大于readBuf则会创建一个新的缓冲区，否则返回readBuf。
	getReadBuf func(size uint32) []byte
	// readBuf 读缓冲区。一定程度上减少内存分配。
	readBuf []byte

	w io.Writer
	// writeBuf 写出缓冲区。
	writeBuf []byte
}

// ReadFrame 读取帧。
func (frw *FrameReaderWriter) ReadFrame() ([]Frame, error) {
	// 读取帧的头部。
	if _, err := io.ReadFull(frw.r, frw.headerBuf[:]); err != nil {
		return nil, err
	}
	return nil, nil
}
