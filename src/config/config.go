package config

import (
	"assert"
	"flag"

	"gopkg.in/fzerorubigd/onion.v2"
)

var Verbose = flag.Bool("v", false, "Verbose status of connection")
var DialTimeout = flag.Int("t", 10, "Timeout of sytem on dial host in seconds")
var Domain = flag.String("d", "google.com", "Critiria doamin for checking connectivity")
var Port = flag.String("p", "80", "Port of host which you wanna dial with it")
var CycleTimeout = flag.Int("c", 1, "Cycle time to check connectivity in seconds")

const (
	organization = "Mimtim"
	appName      = "Netchecker"
)


var Config AppConfig
var o = onion.New()

type AppConfig struct {
	DevelMode       bool   `onion:"devel_mode"`
	Verbose      bool   `onion:"verbose"`
	DialTimeout  int    `onion:"dial_timeout"`
	Domain       string `onion:"domain"`
	Port         string `onion:"port"`
	CycleTimeout int    `onion:"cycle_timeout"`
}

func defaultLayer() onion.Layer {
	d := onion.NewDefaultLayer()
	assert.Nil(d.SetDefault("devel_mode", true))
	assert.Nil(d.SetDefault("verbose", false))
	assert.Nil(d.SetDefault("dial_timeout", 10))
	assert.Nil(d.SetDefault("domain", "google.com"))
	assert.Nil(d.SetDefault("cycle_timeout", 1))
	return d

}
