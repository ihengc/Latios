package rpc

import "Latios/experimental/rpc/backoff"

/*
* @author: Chen Chiheng
* @date: 2023/3/30 0030 16:06
* @description:
**/

type dialOptions struct {
	bs backoff.Strategy
}
