package urlKit

import (
	"fmt"
	"testing"
)

func TestGetUrl(t *testing.T) {
	okStrUrl := `https://ag2.in868.net/app/control/corprator/report/report_agent.php?hall_id=&scid=&cid=11149187&said=11149242&aid=11149672&game_type=LT&wagers_class=wagers&DATE_START=2019-04-15&DATE_END=2019-04-21&report_kind=A&pay_type=&bet_type=&wtype=&is_pay=&lt_num=`
	baseUrl := "https://ag2.in868.net/app/control/corprator/report/report_agent.php"
	param := make(map[string]string)
	param["hall_id"] = ""
	param["scid"] = ""
	param["cid"] = "11149187"
	param["said"] = "11149242"
	param["aid"] = "11149672"
	param["game_type"] = "LT"
	param["wagers_class"] = "wagers"
	param["DATE_START"] = "2019-04-15"
	param["DATE_END"] = "2019-04-21"
	param["report_kind"] = "A"
	param["pay_type"] = ""
	param["bet_type"] = ""
	param["wtype"] = ""
	param["is_pay"] = ""
	param["lt_num"] = ""

	strUrl := GetUrl(baseUrl, param)
	if len(okStrUrl) != len(strUrl) {
		t.Error("GetUrl(", baseUrl, param, ")不正确。")
	} else {
		fmt.Println("GetUrl Ok:", strUrl)
	}

}
