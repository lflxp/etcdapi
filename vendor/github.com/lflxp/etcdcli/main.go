package etcdcli

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/etcd-io/etcd/client"
// )

// func main() {
// 	data := []string{"http://192.168.40.204:2379"}
// 	cfg := client.Config{
// 		Endpoints: data,
// 		Transport: client.DefaultTransport,
// 		// set timeout per request to fail fast when the target endpoint is unavailable
// 		HeaderTimeoutPerRequest: time.Second,
// 	}
// 	c, err := client.New(cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	kapi := client.NewKeysAPI(c)
// 	// set "/foo" key with "bar" value
// 	log.Print("Setting '/foo' key with 'bar' value")
// 	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		// print common key info
// 		log.Printf("Set is done. Metadata is %q\n", resp)
// 	}
// 	// get "/foo" key's value
// 	log.Print("Getting '/foo' key value")
// 	resp, err = kapi.Get(context.Background(), "/foo", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		// print common key info
// 		log.Printf("Get is done. Metadata is %q\n", resp)
// 		// print value
// 		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
// 	}
// }
