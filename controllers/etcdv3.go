package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/lflxp/etcdapi/models"
	"github.com/lflxp/etcdapi/pkg"
)

type EtcdV3Controller struct {
	BaseController
}

// @Title Get
// @Description etcdV3 查询
// @Param	key		body 	models.Key	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv3/get [post]
func (this *EtcdV3Controller) GetValueV3() {
	var key models.Key
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	resp, err := pkg.GetV3(key)
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

	this.ServeJSON()
}

// @Title Set
// @Description etcdV3 修改设置
// @Param	key		body 	models.Value	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv3/set [put]
func (this *EtcdV3Controller) SetValueV3() {
	var key models.Value
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	if key.Key != "" && key.Value != "" {
		resp, err := pkg.SetV3(key)
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
// @Description etcdV3 修改设置
// @Param	key		body 	models.Key	true		"the objectid you want to get"
// @Success 200 {string} success
// @Failure 403 :key is empty
// @router /etcdv3/delete [delete]
func (this *EtcdV3Controller) DeleteValueV3() {
	var key models.Key
	json.Unmarshal(this.Ctx.Input.RequestBody, &key)
	if key.Key != "" {
		resp, err := pkg.DeleteV3(key)
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
