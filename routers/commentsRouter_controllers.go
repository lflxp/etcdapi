package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"],
		beego.ControllerComments{
			Method: "DeleteValueV2",
			Router: `/etcdv2/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"],
		beego.ControllerComments{
			Method: "GetValueV2",
			Router: `/etcdv2/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV2Controller"],
		beego.ControllerComments{
			Method: "SetValueV2",
			Router: `/etcdv2/set`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"],
		beego.ControllerComments{
			Method: "DeleteValueV3",
			Router: `/etcdv3/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"],
		beego.ControllerComments{
			Method: "GetValueV3",
			Router: `/etcdv3/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"] = append(beego.GlobalControllerRouter["github.com/lflxp/etcdapi/controllers:EtcdV3Controller"],
		beego.ControllerComments{
			Method: "SetValueV3",
			Router: `/etcdv3/set`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
