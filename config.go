package mrpc

type (
	RpcServerConf struct {
		Addr    string `json:"Addr"`
		Timeout int64  `json:"Timeout"`
	}
	RpcClientConf struct {
		Target   string `json:"Target"`
		NonBlock bool   `json:"NonBlock"`
		Timeout  int64  `json:"Timeout"`
	}
)
