package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome_go_LM/models"
)

type AreaController struct {
	beego.Controller
}

type AreaResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}

func (this *AreaController) RetData(resp interface{}) {

	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AreaController) GetArea() {

	resp := AreaResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}

	defer this.RetData(&resp)

	o := orm.NewOrm()
	var areas []models.Area
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		resp.Errno = models.RECODE_DBERR
		resp.Errmsg = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if num == 0 {
		resp.Errno = "4002"
		resp.Errmsg = "查询结果为0"
		return
	}

	resp.Data = areas
	return

}
