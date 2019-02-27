package funs

var a2inter = map[string]func(interface{}, interface{}) interface{}{
	"IFERROR": IFERROR,
}

var a4inter = map[string]func(interface{}, interface{}, interface{}, interface{}) interface{}{
	"VLOOKUP": func(p1 interface{}, p2 interface{}, p3 interface{}, p4 interface{}) interface{} {
		var index int
		var approx bool
		if result, ok := p3.(int); ok {
			index = result
		} else if result, ok := p3.(float64); ok {
			index = int(result)
		} else {
			return "N/A"
		}

		if result, ok := p4.(int); ok {
			approx = result == 1
		} else if result, ok := p4.(float64); ok {
			approx = result == 1.0
		} else {
			approx = false
		}

		return VLOOKUP(p1, p2, index, approx)
	},
}
