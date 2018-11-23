package pkg

import (
	"github.com/astaxie/beego"
	"github.com/lflxp/etcdapi/models"
	"github.com/lflxp/etcdcli"
	"go.etcd.io/etcd/clientv3"
)

var ConnV3 *etcdcli.V3

func init() {
	var err error
	beego.Critical(beego.AppConfig.Strings("etcd::url"))
	ConnV3, err = etcdcli.NewConnV3(beego.AppConfig.Strings("etcd::url"), "", "")
	if err != nil {
		beego.Critical(err.Error())
	}
}

func GetV3(data models.Key) (*clientv3.GetResponse, error) {
	var resp *clientv3.GetResponse
	var err error
	if data.Recursive == true {
		resp, err = ConnV3.GetWithPreifxV3(data.Key)
	} else {
		resp, err = ConnV3.GetV3(data.Key)
	}
	return resp, err
}

func SetV3(value models.Value) (*clientv3.PutResponse, error) {
	var resp *clientv3.PutResponse
	var err error
	if value.Ttl != 0 {
		resp, err = ConnV3.PutV3(value.Key, value.Value)
	} else {
		resp, err = ConnV3.PutTtlV3(value.Key, value.Value, value.Ttl)
	}
	return resp, err
}

func DeleteV3(data models.Key) (*clientv3.DeleteResponse, error) {
	var resp *clientv3.DeleteResponse
	var err error
	if data.Recursive == true {
		resp, err = ConnV3.DeleteWithPrefixV3(data.Key)
	} else {
		resp, err = ConnV3.DeleteV3(data.Key)
	}
	return resp, err
}
