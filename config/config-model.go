package config

type Config struct {
	Zendesk ZendeskCredentials `json:"zendesk"`
	Domain  DomainCredentials  `json:"domain"`
}

type ZendeskCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DomainCredentials struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
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
