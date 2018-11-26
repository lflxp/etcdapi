package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/lflxp/etcdapi/models"
	"github.com/lflxp/etcdapi/pkg"
)

type EtcdV2Controller struct {
	BaseController
}

// @Title Get
// @Description etcdv2 查询
// @Param	key		body 	models.Key	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv2/get [post]
func (this *EtcdV2Controller) GetValueV2() {
	var key models.Key
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	if key.Key != "" {
		resp, err := pkg.Get(key)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"status": "failed",
				"data":   err.Error(),
			}
		} else {
			this.Data["json"] = map[string]interface{}{
				"status": "success",
				"data":   resp,
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"status": "failed",
			"data":   "key is none",
		}
	}
	this.ServeJSON()
}

// @Title Set
// @Description etcdv2 修改设置
// @Param	key		body 	models.Value	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv2/set [put]
func (this *EtcdV2Controller) SetValueV2() {
	var key models.Value
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	if key.Key != "" && key.Value != "" {
		resp, err := pkg.Set(key)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"status": "failed",
				"data":   err.Error(),
			}
		} else {
			this.Data["json"] = map[string]interface{}{
				"status": "success",
				"data":   resp,
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"status": "failed",
			"data":   fmt.Sprintf("key %s or value %s is none", key.Key, key.Value),
		}
	}
	this.ServeJSON()
}

// @Title Set
// @Description etcdv2 修改设置
// @Param	key		body 	models.Value	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv2/setdir [put]
func (this *EtcdV2Controller) SetValueDirV2() {
	var key models.Value
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	beego.Critical("11111111111", key)
	if key.Key != "" {
		resp, err := pkg.SetDir(key)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"status": "failed",
				"data":   err.Error(),
			}
		} else {
			this.Data["json"] = map[string]interface{}{
				"status": "success",
				"data":   resp,
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"status": "failed",
			"data":   fmt.Sprintf("key %s or value %s is none", key.Key, key.Value),
		}
	}
	this.ServeJSON()
}

// @Title Delete
// @Description etcdv2 修改设置
// @Param	key		body 	models.Key	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv2/delete [delete]
func (this *EtcdV2Controller) DeleteValueV2() {
	var key models.Key
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	if key.Key != "" {
		resp, err := pkg.Delete(key)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"status": "failed",
				"data":   err.Error(),
			}
		} else {
			this.Data["json"] = map[string]interface{}{
				"status": "success",
				"data":   resp,
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"status": "failed",
			"data":   fmt.Sprintf("key %s none", key.Key),
		}
	}
	this.ServeJSON()
}
