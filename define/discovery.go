package define

type Discovery struct {
	ID string `json:"id"`
	Request string `json:"request"`
}

type ServerAddr struct {
	Addr string `json:"addr"`
}