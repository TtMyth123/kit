package funs

var a1int64map = map[string]func(interface{}) (int, error){
	"LEN": LEN,
}

var a2int64map = map[string]func(interface{}, interface{}) (int, error){
	"MATCH": func(p1, p2 interface{}) (int, error) {
		return MATCH(p1, p2, 1)
	},
	"FIND":   FIND2,
	"SEARCH": SEARCH2,
}

var a3int64map = map[string]func(interface{}, interface{}, interface{}) (int, error){
	"MATCH": func(p1, p2 interface{}, p3 interface{}) (int, error) {
		return MATCH(p1, p2, int(p3.(float64)))
	},
	"FIND":   FIND3,
	"SEARCH": SEARCH3,
}
