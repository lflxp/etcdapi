// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/lflxp/etcdapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v2",
			beego.NSInclude(
				&controllers.EtcdV2Controller{},
			),
		),
		beego.NSNamespace("/v3",
			beego.NSInclude(
				&controllers.EtcdV3Controller{},
			),
		),
	)
	beego.AddNamespace(ns)
}
