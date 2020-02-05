package Result

type JsonDatatablesResult struct {
	Meta MetaInfo    `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaInfo struct {
	Page    int    `json:"page"`    //当前第几页
	Pages   int    `json:"pages"`   //一共多少页
	Perpage int    `json:"perpage"` //每页多少条数据
	Total   int    `json:"total"`   //一共多少条数据
	Sort    string `json:"sort"`
	Field   string `json:"field"`
}
