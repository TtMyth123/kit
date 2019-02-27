package funs

var a1float64map = map[string]func(interface{}) (float64, error){
	"FLOOR":      FLOOR,
	"SUM":        SUM,
	"MIN":        MIN,
	"MAX":        MAX,
	"HARMEAN":    HARMEAN,
	"MEDIAN":     MEDIAN,
	"VALUE":      VALUE,
	"SUMPRODUCT": SUMPRODUCT1,
}

var a2float64map = map[string]func(interface{}, interface{}) (float64, error){
	"SUM":   SUM2,
	"POWER": POWER,
	"ROUND": func(p1 interface{}, precision interface{}) (float64, error) {
		return ROUND(p1.(float64), precision.(float64))
	},
	"COUNTIF":    COUNTIF,
	"MIN":        MIN2,
	"MAX":        MAX2,
	"HARMEAN":    HARMEAN2,
	"FLOOR":      FLOOR2,
	"MEDIAN":     MEDIAN2,
	"SUMPRODUCT": SUMPRODUCT2,
}

var a3float64map = map[string]func(interface{}, interface{}, interface{}) (float64, error){
	"POISSON_DIST":       POISSON_DIST3,
	"_xlfn.POISSON.DIST": POISSON_DIST3,
}
