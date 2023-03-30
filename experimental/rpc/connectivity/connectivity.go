package connectivity

/*
 * @author: Chen Chiheng
 * @date: 2023/3/29 21:34
 * @description:
 */

// State 定义连接状态。
type State int

func (s State) Strings() string {
	switch s {
	case Idle:
		return "IDLE"
	case Connecting:
		return "CONNECTING"
	case Ready:
		return "READY"
	case TransientFailure:
		return "TRANSIENT_FAILURE"
	case Shutdown:
		return "SHUTDOWN"
	default:
		return "INVALID_STATE"
	}
}

// 连接将在下面5个状态中进行转换。可以看成一个状态机。
const (
	// Idle 表示连接空闲（一个处于 Ready 状态的连接在一段时间内未发生RPC则会转到此状态）。
	// 对于在此状态下连接的RPC操作，此状态将会转换到 Connecting 。
	// 设计此状态的目的是为了降低服务器连接负载。
	Idle State = iota
	// Connecting 表示正在建立连接。
	Connecting
	// Ready 表示此连接可以用来发送或接收消息。
	Ready
	// TransientFailure 表示此连接发生错误，期望恢复。
	// 比如在进行TCP握手时发送失败。
	TransientFailure
	// Shutdown 表示此连接已开始关闭。
	Shutdown
)
