package main

import (
	"log"

	"go.ligato.io/cn-infra/v2/agent"
	"go.ligato.io/cn-infra/v2/logging"
	frr "go.ligato.io/vpp-agent/v3/plugins/frr"
)

func main() {
	/* p := new(HelloWorld)
	p.SetName("demo plugin")
	a := agent.NewAgent(agent.Plugins(p))
	if err := a.Run(); err != nil {
		log.Fatalln(err)
	} */

	p := frr.NewFRRPlugin()
	p.Log.SetLevel(logging.DebugLevel)
	/* 	restConf := rest.Config{
	   		Endpoint: "0.0.0.0" + ":9292",
	   	}
	   	option := rest.UseConf(restConf)
	   	p.REST = rest.NewPlugin(option, ChangeLogLevel) */
	p.Deps.Log.SetLevel(logging.DebugLevel)
	p.Log.Debugf("FRR plugin starting wit name %v", p.PluginName.String())
	p.Log.Debugf("FRR plugin starting wit Config name %v", p.Cfg.GetConfigName())
	a := agent.NewAgent(agent.AllPlugins(p))
	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}

	/* broker := p.KVStore.NewBroker("/frrPlugin")

	//trial on proto

	// Retrieve value from KV store
	value := new(frr.Greetings)
	found, _, err := broker.GetValue("greetings/hello", value)
	if err != nil {
		p.Log.Errorf("GetValue failed: %v", err)
	} else if !found {
		p.Log.Info("No greetings found..")
	} else {
		p.Log.Infof("Found some greetings: %+v", value)
	}

	// Wait few seconds
	time.Sleep(time.Second * 2)

	p.Log.Infof("updating..")

	value = &frr.Greetings{
		Greeting: "hello",
	}

	if err := broker.Put("greetings/hello", value); err != nil {
		p.Log.Errorf("Put failed: %v", err)
	}

	fmt.Println(p.String()) */

}
