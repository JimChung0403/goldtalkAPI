package dao

import (
	"goldtalkAPI/pkg/client"

	"goldtalkAPI/conf"
	"testing"
)

func TestMain(m *testing.M) {
	conf.LoadConfigFile("../../conf/service.conf")
	client.InitClients(conf.Conf)
	m.Run()
}
