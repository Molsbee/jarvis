package clc

import (
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/model"
	"github.com/Molsbee/jarvis/service/clc/elasticsearch"
)

func GetServerDetails(env config.Environment, name string) (model.ServerResponse, error) {
	return elasticsearch.Main(env).GetServerDetails(name)
}

func GetServerDetailsByHardwareUUID(env config.Environment, uuid string) (model.ServerResponse, error) {
	return elasticsearch.Main(env).GetServerDetailsByHardwareUUID(uuid)
}
