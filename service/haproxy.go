package service

import (
	"encoding/csv"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/model"
	"net/http"
)

func GetStatsPage(dataCenter string) (list []model.HAProxyStatsResponse, err error) {
	dc := config.UserConfig.GetDataCenter(dataCenter)
	if dc == nil {
		err = fmt.Errorf("unsupported data center %s provided", dataCenter)
		return
	}

	resp, err := http.DefaultClient.Get(dc.HAProxy)
	if err != nil {
		err = fmt.Errorf("failed to get stats page information for data center %s", dc.Name)
		return
	}
	defer resp.Body.Close()

	lines, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		err = fmt.Errorf("failed to read csv information for data center %s", dc.Name)
		return
	}

	for i := 1; i < len(lines); i++ {
		l := lines[i]
		list = append(list, model.HAProxyStatsResponse{Location: dataCenter, ProxyName: l[0], ServiceName: l[1], Status: l[17]})
	}

	return
}
