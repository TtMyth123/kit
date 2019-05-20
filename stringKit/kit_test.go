package stringKit

import "testing"

func TestGetBetweenStr(t *testing.T) {
	s := ""
	//s = GetBetweenStr("aaabbcc", "aa", "c")
	//if s!="abb" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}
	//
	//s = GetBetweenStr("1234567890", "a", "1")
	//if s!="" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}
	//
	//s = GetBetweenStr("1234567890", "1", "1")
	//if s!="" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}
	//
	//s = GetBetweenStr("1234567890", "1", "2")
	//if s!="" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}
	//s = GetBetweenStr("1234567890", "1", "3")
	//if s!="2" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}


	ss := `&amp;scid=&amp;cid=11149187&amp;game_type=LT&amp;wagers_class=wagers&amp;DATE_START=2019-04-15&amp;DATE_END=2019-04-21&amp;report_kind=A&amp;pay_type=&amp;bet_type=&amp;wtype=&amp;is_pay=&amp;lt_num=">313946.00</a>`
	s = GetBetweenStr(ss, ";cid=", "&")
	if s!="11149187" {
		t.Error("TestGetBetweenStr()不正确。")
	}

	//ss = `
    //            <a href="report_co.php?hall_id=&amp;scid=&amp;cid=11149187&amp;game_type=LT&amp;wagers_class=wagers&amp;DATE_START=2019-04-15&amp;DATE_END=2019-04-21&amp;report_kind=A&amp;pay_type=&amp;bet_type=&amp;wtype=&amp;is_pay=&amp;lt_num=">313946.00</a>
    //          `
	//s = GetBetweenStr(ss, "cid=", "&")
	//if s!="11149187" {
	//	t.Error("TestGetBetweenStr()不正确。")
	//}
}
