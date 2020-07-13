package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"net/http"
	"strings"
)

func post(url string, data string) (v ElasticResponse, err error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	domainCredentials := config.GetConfig().Domain
	request.SetBasicAuth(domainCredentials.Username, domainCredentials.Password)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		err = fmt.Errorf("failed to perform post request to elasticsearch - err (%s)", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		err = fmt.Errorf("failed to convert response to interface - err (%s)", err)
	}

	return
}
