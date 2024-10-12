package main

import (
	"context"
	"os"

	"github.com/magicsong/kidecar/pkg/assembler"
	"github.com/magicsong/kidecar/pkg/info"
	"github.com/magicsong/kidecar/pkg/manager"
	"github.com/magicsong/kidecar/pkg/plugins"
	flag "github.com/spf13/pflag"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "/opt/sidecar/config.yaml", "config file path")
}

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("manager-examples")
	flag.Parse()
	sidecar := assembler.NewSidecar()
	if err := sidecar.LoadConfig(configPath); err != nil {
		log.Error(err, "failed to load config")
		os.Exit(1)
	}
	mgr, err := manager.NewManager()
	if err != nil {
		log.Error(err, "failed to create manager")
		panic(err)
	}
	info.SetGlobalKubeInterface(mgr)
	sidecar.SetupWithManager(mgr)
	// add plugins
	for _, v := range plugins.PluginRegistry {
		if err := sidecar.AddPlugin(v); err != nil {
			log.Error(err, "failed to add plugin", "pluginName", v.Name())
			panic(err)
		}
	}
	ctx := context.TODO()
	if err := sidecar.Start(ctx); err != nil {
		panic(err)
	}
}
