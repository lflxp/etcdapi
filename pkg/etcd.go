package pkg

import (
	"github.com/astaxie/beego"
	"github.com/etcd-io/etcd/client"
	"github.com/lflxp/etcdapi/models"
	"github.com/lflxp/etcdcli"
)

var Conn *etcdcli.V2

func init() {
	var err error
	beego.Critical(beego.AppConfig.Strings("etcd::url"))
	Conn, err = etcdcli.NewConnV2(beego.AppConfig.Strings("etcd::url"), "", "")
	if err != nil {
		beego.Critical(err.Error())
	}
}

func Get(data models.Key) (*client.Response, error) {
	var resp *client.Response
	var err error
	if data.Recursive == true {
		resp, err = Conn.GetWithPreifxV2(data.Key)
	} else {
		resp, err = Conn.GetV2(data.Key)
	}
	return resp, err
}

func Set(value models.Value) (*client.Response, error) {
	var resp *client.Response
	var err error
	if value.Ttl != 0 {
		resp, err = Conn.PutV2(value.Key, value.Value)
	} else {
		resp, err = Conn.PutTtlV2(value.Key, value.Value, value.Ttl)
	}
	return resp, err
}

func SetDir(value models.Value) (*client.Response, error) {
	var resp *client.Response
	var err error
	resp, err = Conn.SetDir(value.Key)
	return resp, err
}

func Delete(data models.Key) (*client.Response, error) {
	var resp *client.Response
	var err error
	if data.Recursive == true {
		resp, err = Conn.DeleteWithPrefixV2(data.Key)
	} else {
		resp, err = Conn.DeleteV2(data.Key)
	}
	return resp, err
}
