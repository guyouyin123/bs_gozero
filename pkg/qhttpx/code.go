package qhttpx

const (
	//golbal 00~100
	StateOk                  = "00" // 成功
	StateInvalidParameter    = "01" // 参数错误
	StateFailed              = "02" // 失败
	StateCrossDomainPrelight = "03" // 跨域预检
	StateNoRecord            = "04" // 无记录
	StateDuplicate           = "05" // 重复创建
)
