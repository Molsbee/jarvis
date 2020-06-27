package config

type Config struct {
	Domain struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"domain"`
	DataCenters []DataCenter `json:"dataCenters"`
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
