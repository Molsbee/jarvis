package config

type Config struct {
	Zendesk struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"zendesk"`
	Domain struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"domain"`
}

type DataCenter struct {
	Name          string `json:"name"`
	HAProxy       string `json:"haProxy"`
	MSSQL         string `json:"mssql"`
	ElasticSearch ElasticSearch
}

type ElasticSearch struct {
	Foundation string
	Platform   string
	Audits     string
}
