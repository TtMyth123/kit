package funs

// add ty
var a1StringMap = map[string]func(interface{}) (string, error){
	"CONCATENATE": CONCATENATE,
	"LEFT":        LEFT1,
	"RIGHT":       RIGHT1,
}

// add ty
var a2StringMap = map[string]func(interface{}, interface{}) (string, error){
	"CONCATENATE": CONCATENATE2,
	"TEXT":        TEXT,
	"LEFT":        LEFT2,
	"RIGHT":       RIGHT2,
}

// add ty
var a3StringMap = map[string]func(interface{}, interface{}, interface{}) (string, error){
	"CONCATENATE": CONCATENATE3,
	"MID":         MID,
}

// add ty
var a4StringMap = map[string]func(interface{}, interface{}, interface{}, interface{}) (string, error){
	"CONCATENATE": CONCATENATE4,
	"REPLACE":     REPLACE4,
}

// add ty
var anStringMap = map[string]func(...interface{}) (string, error){
	"CONCATENATE": CONCATENATEn,
}
