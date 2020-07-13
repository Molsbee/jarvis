package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const FileName = ".jarvis-config.json"

var (
	parsedConfig = false
	DataCenters  = []DataCenter{
		{
			Name:    "AU1",
			HAProxy: "http://haproxy-au1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-au1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-au1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-au1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-au1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "CA1",
			HAProxy: "http://haproxy-ca1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-ca1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-ca1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-ca1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-ca1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "CA2",
			HAProxy: "http://haproxy-ca2.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-ca2.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-ca2.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-ca2.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-ca2.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "CA3",
			HAProxy: "http://haproxy-ca3.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-ca3.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-ca3.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-ca3.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-ca3.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "DE1",
			HAProxy: "http://haproxy-de1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-de1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-de1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-de1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-de1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "DE3",
			HAProxy: "http://haproxy-de3.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-de3.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-de3.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-de3.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-de3.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "GB1",
			HAProxy: "http://haproxy-gb1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-gb1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-gb1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-gb1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-gb1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "GB3",
			HAProxy: "http://haproxy-gb3.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-gb3.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-gb3.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-gb3.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-gb3.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "IL1",
			HAProxy: "http://haproxy-il1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-il1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-il1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-il1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-il1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "NE1",
			HAProxy: "http://haproxy-ne1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-ne1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-ne1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-ne1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-ne1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "NY1",
			HAProxy: "http://haproxy-ny1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-ny1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-ny1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-ny1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-ny1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "SG1",
			HAProxy: "http://haproxy-sg1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-sg1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-sg1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-sg1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-sg1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "UC1",
			HAProxy: "http://haproxy-uc1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-uc1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-uc1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-uc1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-uc1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "VA1",
			HAProxy: "http://haproxy-va1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-va1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-va1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-va1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-va1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "VA2",
			HAProxy: "http://haproxy-va2.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-va2.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-va2.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-va2.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-va2.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "WA1",
			HAProxy: "http://haproxy-wa1.t3n.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-wa1.t3n.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-wa1.t3n.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-wa1.t3n.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-wa1.t3n.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "LB1",
			HAProxy: "http://haproxy-lb1.t3dev.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-lb1.t3dev.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-lb1.t3dev.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-lb1.t3dev.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-lb1.t3dev.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "QA5",
			HAProxy: "http://haproxy-qa5.t3dev.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-qa5.t3dev.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-qa5.t3dev.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-qa5.t3dev.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-qa5.t3dev.dom:1337/audits/_search?pretty=true",
			},
		},
		{
			Name:    "LB2",
			HAProxy: "http://haproxy-lb2.t3ppe.dom:1936/stats;csv;norefresh",
			MSSQL:   "platformsql-lb2.t3ppe.dom:1433",
			ElasticSearch: ElasticSearch{
				Foundation: "http://search-lb2.t3ppe.dom:1337/foundation/_search?pretty=true",
				Platform:   "http://search-lb2.t3ppe.dom:1337/platform/_search?pretty=true",
				Audits:     "http://search-lb2.t3ppe.dom:1337/audits/_search?pretty=true",
			},
		},
	}
)

func GetDataCenter(dataCenter string) *DataCenter {
	for _, dc := range DataCenters {
		if strings.ToLower(dc.Name) == strings.ToLower(dataCenter) {
			return &dc
		}
	}
	return nil
}

func GetConfig() Config {
	userHomeDir, _ := os.UserHomeDir()
	workingDir, _ := os.Getwd()
	directories := []string{
		fmt.Sprintf("%s/%s", userHomeDir, FileName),
		fmt.Sprintf("%s/%s", workingDir, FileName),
		fmt.Sprintf("%s/../%s", workingDir, FileName),
	}

	for _, dir := range directories {
		config, err := parseConfig(dir)
		if err == nil {
			return config
		}
	}

	log.Panic("unable to parse user config")
	return Config{}
}

func parseConfig(filePath string) (config Config, err error) {
	if _, fErr := os.Stat(filePath); !os.IsNotExist(fErr) {
		data, rErr := ioutil.ReadFile(filePath)
		if rErr != nil {
			err = fmt.Errorf("unable to read file %s", filePath)
			return
		}

		if jErr := json.Unmarshal(data, &config); jErr != nil {
			err = fmt.Errorf("unable to parse config file %s properly", filePath)
			return
		}

		return
	}

	err = fmt.Errorf("file not found at provided path %s", filePath)
	return
}

func Write(config Config) error {
	data, _ := json.Marshal(config)

	workingDirectory, _ := os.UserHomeDir()
	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s", workingDirectory, FileName), data, os.FileMode(0600)); err != nil {
		return fmt.Errorf("unable to write config file to home directory - err (%s)", err)
	}

	return nil
}
