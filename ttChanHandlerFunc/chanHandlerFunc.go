package ttChanHandlerFunc

type ChanData struct {
	Key  string
	Data interface{}
}

type ChanHandlerFunc struct {
	gChan         chan ChanData
	mpHandlerFunc map[string]map[string]HandlerFunc
}

type HandlerFunc func(data interface{})

func NewChanHandlerFunc() *ChanHandlerFunc {
	aGChan := new(ChanHandlerFunc)
	aGChan.gChan = make(chan ChanData)
	aGChan.mpHandlerFunc = make(map[string]map[string]HandlerFunc)

	go aGChan.run()
	return aGChan
}

func (this *ChanHandlerFunc) AddData(key string, data interface{}) {
	aChanData := ChanData{
		Key:  key,
		Data: data,
	}

	this.gChan <- aChanData
}

func (this *ChanHandlerFunc) AddHandlerFunc(key string, funKey string, fun HandlerFunc) bool {
	a := this.mpHandlerFunc[key]
	if a == nil {
		this.mpHandlerFunc[key] = make(map[string]HandlerFunc)
	}

	this.mpHandlerFunc[key][funKey] = fun

	return true
}

func (this *ChanHandlerFunc) DelHandlerFunc(key string) {
	delete(this.mpHandlerFunc, key)
}

func (this *ChanHandlerFunc) run() {
	for msg := range this.gChan {
		a := this.mpHandlerFunc[msg.Key]
		for _, fun := range a {
			fun(msg.Data)
		}
	}

	//for msg := range this.gChan {
	//	this.mpHandlerFunc.Range(func(key, funV interface{}) bool {
	//		fun := funV.(HandlerFunc)
	//		go fun(msg)
	//		return true
	//	})
	//}
}
