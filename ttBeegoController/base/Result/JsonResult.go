package Result

import (
	"github.com/TtMyth123/kit/ttBeegoController/base/enums"
)

type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  interface{}          `json:"obj"`
}
