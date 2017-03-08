package config

import (
	"assert"
	"os"
	"os/user"
	"path/filepath"

	"fmt"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"gopkg.in/fzerorubigd/onion.v2"
	"gopkg.in/fzerorubigd/onion.v2/extraenv"
)

var (
	all []Initializer
)

// Initializer is the config initializer for module
type Initializer interface {
	// Initialize is called when the module is going to add its layer
	Initialize(*onion.Onion) []onion.Layer
	// Loaded inform the modules that all layer are ready
	Loaded()
}

//Initialize try to initialize config
func Initialize() {
	usr, err := user.Current()
	if err != nil {
		logrus.Warn(err)
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Warn(err)
	}

	assert.Nil(o.AddLayer(defaultLayer()))
	if err = o.AddLayer(onion.NewFileLayer("/etc/" + organization + "/" + appName + ".yaml")); err == nil {

		logrus.Infof("loading config from %s", "/etc/"+organization+"/"+appName+".yaml")
	}
	if err = o.AddLayer(onion.NewFileLayer(usr.HomeDir + "/." + organization + "/" + appName + ".yaml")); err == nil {
		logrus.Infof("loading config from %s", usr.HomeDir+"/."+organization+"/"+appName+".yaml")
	}
	if err = o.AddLayer(onion.NewFileLayer(dir + "/configs/" + appName + ".yaml")); err == nil {
		logrus.Infof("loading config from %s", dir+"/configs/"+appName+".yaml")
	}
	for i := range all {
		nL := all[i].Initialize(o)
		for l := range nL {
			_ = o.AddLayer(nL[l])
		}
	}

	o.AddLazyLayer(extraenv.NewExtraEnvLayer("GAD"))
	o.GetStruct("", &Config)
	// TODO {fzerorubigd}: Onion does not support slice in struct mapping

	for i := range all {
		all[i].Loaded()
	}

	j, _ := json.MarshalIndent(Config, "\t", "\t")
	fmt.Println(string(j))

}

// SetConfigParameter try to set the config parameter for the logrus base on config
func SetConfigParameter() {
	if Config.DevelMode {
		// In development mode I need colors :) candy mode is GREAT!
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, DisableColors: false})
		logrus.SetLevel(logrus.DebugLevel)

	} else {
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: false, DisableColors: true})
		logrus.SetLevel(logrus.WarnLevel)
	}

}

// Register a config module
func Register(i ...Initializer) {
	all = append(all, i...)
}
