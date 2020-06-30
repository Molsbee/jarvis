package config

import "strings"

type Config struct {
	Domain struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"domain"`
	DataCenters []DataCenter `json:"dataCenters"`
}

func (c Config) GetDataCenter(dataCenter string) *DataCenter {
	for _, dc := range c.DataCenters {
		if strings.ToLower(dc.Name) == strings.ToLower(dataCenter) {
			return &dc
		}
	}
	return nil
}

type DataCenter struct {
	Name          string `json:"name"`
	HAProxy       string `json:"haProxy"`
	MSSQL         string `json:"mssql"`
	ElasticSearch struct {
		Main        string `json:"main"`
		Foundation  string `json:"foundation"`
		Platform    string `json:"platform"`
		Audits      string `json:"audits"`
		Credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"credentials"`
	} `json:"elasticSearch"`
}
