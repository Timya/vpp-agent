package frr

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/unrolled/render"
	"go.ligato.io/cn-infra/v2/datasync"
	"go.ligato.io/cn-infra/v2/db/keyval"
	"go.ligato.io/cn-infra/v2/db/keyval/etcd"
	"go.ligato.io/cn-infra/v2/examples/tutorials/04_kv-store/model"
	"go.ligato.io/cn-infra/v2/infra"
)

//go:generate protoc --proto_path=model --go_out=model ./model/frr.proto

type FRRPlugin struct {
	Deps
	watchCloser chan string
}

type Deps struct {
	infra.PluginDeps
	// REST    rest.HTTPHandlers
	KVStore     keyval.KvProtoPlugin
	watchCloser chan string
}

const keyPrefix = "/frrPlugin/"

func (p *FRRPlugin) fooHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			p.Log.Errorf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		formatter.Text(w, http.StatusOK, fmt.Sprintf("Hello %s", body))
	}
}

func (p *FRRPlugin) Init() (err error) {
	p.Log.Info("Loading Config")

	if _, err = p.loadConfig(); err != nil {
		return err
	}

	// p.Log.Infof("config: %+v", p.config)
	/* 	if p.REST != nil {
	   		p.REST.RegisterHTTPHandler("/greeding", p.fooHandler, "GET")
	   	}
	*/
	p.initKVS()

	return nil
}

func (p *FRRPlugin) initKVS() error {
	if p.KVStore.Disabled() {
		return fmt.Errorf("KV store is disabled")
	}

	watcher := p.KVStore.NewWatcher(keyPrefix)

	// Start watching for changes
	err := watcher.Watch(p.onChange, p.watchCloser, "greetings/")
	if err != nil {
		return err
	}

	return nil
}

func (p *FRRPlugin) onChange(resp datasync.ProtoWatchResp) {
	value := new(model.Greetings)
	// Deserialize data
	if err := resp.GetValue(value); err != nil {
		p.Log.Errorf("GetValue for change failed: %v", err)
		return
	}
	p.Log.Infof("%v change, Key: %q Value: %+v", resp.GetChangeType(), resp.GetKey(), value)
}

// AfterInit is executed after agent initialization.
func (p *FRRPlugin) AfterInit() error {
	go p.updater()
	return nil
}

func (p *FRRPlugin) updater() {
	broker := p.KVStore.NewBroker(keyPrefix)

	// Retrieve value from KV store
	value := new(model.Greetings)
	found, _, err := broker.GetValue("greetings/hello", value)
	if err != nil {
		p.Log.Errorf("GetValue failed: %v", err)
	} else if !found {
		p.Log.Info("No greetings found..")
	} else {
		p.Log.Infof(" iFound some greetings: %+v", value)
	}

	// Wait few seconds
	time.Sleep(time.Second * 2)

	p.Log.Infof("updating..")

	// Prepare data
	value = &model.Greetings{
		Greeting: "Hello1111111111111",
	}

	// Update value in KV store
	if err := broker.Put("greetings/hello", value); err != nil {
		p.Log.Errorf("Put failed: %v", err)
	}
}

func (p *FRRPlugin) Close() error {
	log.Println("frr closed")
	return nil
}

func NewFRRPlugin() *FRRPlugin {

	p := &FRRPlugin{
		watchCloser: make(chan string),
	}
	// p.REST = &rest.DefaultPlugin
	p.PluginName = "frr"
	p.SetName("frr")
	p.Setup()
	//initialize kvs
	p.KVStore = &etcd.DefaultPlugin

	return p
}
