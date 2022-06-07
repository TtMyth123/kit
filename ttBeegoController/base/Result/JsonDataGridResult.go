package Result

import "github.com/TtMyth123/kit/ttBeegoController/base/enums"

type JsonDataGridResult struct {
	Code         enums.JsonResultCode `json:"code"`
	Total        int64                `json:"total"`
	Rows         interface{}          `json:"rows"`
	RecordsTotal int64                `json:"recordsTotal"`
}
