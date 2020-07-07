package elasticsearch

import "github.com/Molsbee/jarvis/model"

type ElasticResponse struct {
	Hits struct {
		Hits []struct {
			ID     string  `json:"_id"`
			Index  string  `json:"_index"`
			Score  float32 `json:"_score"`
			Source struct {
				Doc struct {
					VM                  model.VM                  `json:"vm"`
					ServerConfiguration model.ServerConfiguration `json:"serverConfiguration"`
				} `json:"doc"`
				Meta struct {
					Expiration int    `json:"expiration"`
					Flags      int    `json:"flags"`
					ID         string `json:"id"`
					REV        string `json:"rev"`
				}
			} `json:"_source"`
			Type string `json:"_type"`
		} `json:"hits"`
		MaxScore float32 `json:"max_score"`
		Total    int     `json:"total"`
	} `json:"hits"`
	Shards struct {
		Failed     int `json:"failed"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
}
