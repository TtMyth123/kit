package funs

var a1boolmap = map[string]func(interface{}) bool{
	"OR":      OR,
	"AND":     AND,
	"ISERROR": ISERROR,
}

var a2boolmap = map[string]func(interface{}, interface{}) bool{
	"OR":  OR2,
	"AND": AND2,
}

var a3boolmap = map[string]func(interface{}, interface{}, interface{}) bool{
	"OR":  OR3,
	"AND": AND3,
}
