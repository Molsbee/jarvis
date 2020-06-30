package model

import "fmt"

type HAProxyStatsResponse struct {
	Location    string
	ProxyName   string
	ServiceName string
	Status      string
}

func (h HAProxyStatsResponse) String() string {
	return fmt.Sprintf("%s - %-20s %-20s %s", h.Location, h.ProxyName, h.ServiceName, h.Status)
}
