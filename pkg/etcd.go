package pkg

import (
	"github.com/astaxie/beego"
	"github.com/etcd-io/etcd/client"
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

func Get(name string) (*client.Response, error) {
	resp, err := Conn.GetV2(name)
	return resp, err
}

func Set(k, v string) (*client.Response, error) {
	resp, err := Conn.PutV2(k, v)
	return resp, err
}

func SetTtl(k, v string, ttl int64) (*client.Response, error) {
	resp, err := Conn.PutTtlV2(k, v, ttl)
	return resp, err
}

func Delete(k string) (*client.Response, error) {
	resp, err := Conn.DeleteV2(k)
	return resp, err
}

func DeleteAll(k string) (*client.Response, error) {
	resp, err := Conn.DeleteWithPrefixV2(k)
	return resp, err
}
