package ttBeegoController

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/ttBeegoController/base/Result"
	"ttmyth123/kit/ttBeegoController/base/enums"
	"ttmyth123/kit/ttBeegoController/cache"
)

type TtBaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
	viewDir        string //view目录
}

/**
60s*30=30分钟
*/
const Cache_Timeout = 60 * 30

func (this *TtBaseController) Prepare() {
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.viewDir = this.controllerName[0 : len(this.controllerName)-10]
}
func (this *TtBaseController) GetControllerName() string {
	return this.controllerName
}

func (this *TtBaseController) GetActionName() string {
	return this.actionName
}

func (this *TtBaseController) GetViewDir() string {
	return this.viewDir
}
func (this *TtBaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	res := &Result.JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}
func (this *TtBaseController) JsonDataGridResult(Total int64, Rows interface{}) {
	res := &Result.JsonDataGridResult{Code: enums.JRCodeSucc, Total: Total, Rows: Rows, RecordsTotal: Total}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *TtBaseController) JsonDatatablesResult(Page, OnePageCount, Total int, Rows interface{}) {
	_ ,Pages:= sqlKit.GetOffset(Total, OnePageCount, Page)
	aMetaInfo := Result.MetaInfo{Page: Page, Pages: Pages, Perpage: OnePageCount, Total: Total}
	res := &Result.JsonDatatablesResult{Meta: aMetaInfo, Data: Rows}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}
func (this *TtBaseController) Logout(name string) {
	Id, _ := this.GetInt("UserId", 0)
	SID := this.GetString("SID")
	key := fmt.Sprintf("%d_%s_%s", Id, SID, name)
	cache.DelCache(key);
	this.DelSession(name);
}

// 重定向
func (this *TtBaseController) RedirectTt(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}
func (this *TtBaseController) GetTtSession(name string) interface{} {
	Id, _ := this.GetInt("UserId", 0)
	SID := this.GetString("SID")

	key := fmt.Sprintf("%d_%s_%s", Id, SID, name)
	var v interface{}
	e := cache.GetCache(key, &v)
	if e != nil {
		v = this.GetSession(name)
		if v != nil {
			return v
		}
	}
	return v
}

func (this *TtBaseController) SetTtSession(name string, v interface{}) {
	Id, _ := this.GetInt("UserId", 0)
	SID := this.GetString("SID")

	key := fmt.Sprintf("%d_%s_%s", Id, SID, name)
	cache.SetCache(key, v, Cache_Timeout)
	this.SetSession(name, v)
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (this *TtBaseController) SetTpl(template ...string) {
	layout := "shared/layout_page.html"

	var tplName string
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		actionName := strings.ToLower(this.actionName)
		tplName = this.viewDir + "/" + actionName + ".html"
	}

	this.Layout = layout
	this.TplName = tplName
	return
}
