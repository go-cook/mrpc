package mrpc

type (
	RpcServerConf struct {
		ServiceConf
		Addr    string `json:"Addr"`
		Timeout int64  `json:"Timeout"`
	}
	RpcClientConf struct {
		Target   string `json:"Target"`
		NonBlock bool   `json:"NonBlock"`
		Timeout  int64  `json:"Timeout"`
	}
	ServiceConf struct {
		Name       string `json:"Name"`
		Mode       string `json:"Mode"`
		MetricsUrl string `json:"MetricsUrl"`
		Prometheus string `json:"Prometheus"`
	}
)
