package netx

/*
* @author: Chen Chiheng
* @date: 2023/4/4 0004 11:18
* @description:
**/

func fnWrapper(fn func()) (err error) {
	fn()
	defer func() {
		if re := recover(); re != nil {
			err = re.(error)
		}
	}()
	return
}
