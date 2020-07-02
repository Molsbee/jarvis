package clc

import (
	"github.com/Molsbee/jarvis/elasticsearch"
)

func GetServerDetails(env elasticsearch.Environment, name string) (elasticsearch.ServerResponse, error) {
	return elasticsearch.Main(env).GetServerDetails(name)
}

func GetServerDetailsByHardwareUUID(env elasticsearch.Environment, uuid string) (elasticsearch.ServerResponse, error) {
	return elasticsearch.Main(env).GetServerDetailsByHardwareUUID(uuid)
}
