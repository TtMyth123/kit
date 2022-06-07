package stringKit

import (
	"fmt"
	"github.com/ttmyth123/kit/strconvEx"
	"strconv"
	"testing"
	"time"
)

func TestAA(t *testing.T) {

	f := 1234.3
	fmt.Println("a1:", fmt.Sprintf("%.2f", f))
	fmt.Sprintf("")
	a := strconv.FormatFloat(strconvEx.Decimal(f), 'G', 9, 32)

	f = 1234.356
	fmt.Println("a2:", fmt.Sprintf("%.2f", f))
	a1 := strconv.FormatFloat(strconvEx.Decimal(f), 'G', 9, 32)

	f = 1234000.3
	fmt.Println("a2:", fmt.Sprintf("%.2f", f))
	a2 := strconv.FormatFloat(strconvEx.Decimal(f), 'G', 9, 32)
	f = 1234.33423
	fmt.Println("a3:", fmt.Sprintf("%.2f", f))
	a3 := strconv.FormatFloat(strconvEx.Decimal(f), 'G', 9, 32)

	fmt.Println(a, a1, a2, a3) // 100.12

}

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
	if s != "11149187" {
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

func TestGetNumGuid(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go aaaGetNumGuid()
	}
	time.Sleep(time.Second * 100)
}

func aaaGetNumGuid() {
	a := GetNumGuid(2)
	fmt.Println(a)
}

func TestGetBetweenStr11(t *testing.T) {
	s := `
 <input type="hidden" id="goldidtmp" value="">

		</table>
	</div>
	<input type="hidden" value="8XX6@85.9,6XX6@85.9" id="num_str" name="num_str">
 <input type="hidden" id="goldidtmp" value="">
 <input type="hidden" id="allsettmp" value="">
 <input type="hidden" id="ptmp" value="">
`
	a := GetBetweenStrEx(s, `" id="num_str" name="num_str"`, `value="`)
	fmt.Println(a)
}
