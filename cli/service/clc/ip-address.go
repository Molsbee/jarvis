package clc

import (
	"encoding/json"
	"fmt"
	"github.com/Molsbee/jarvis/common_model"
	"net/http"
)

func FindIPAddress(addressString string) (ipAddress common_model.IPAddress, err error) {
	resp, hErr := http.DefaultClient.Get("http://localhost:4000/clc/ipAddresses/" + addressString)
	if hErr != nil {
		err = fmt.Errorf("error calling javis server %s", hErr)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("non 200 status code received from server")
	}

	if jErr := json.NewDecoder(resp.Body).Decode(&ipAddress); jErr != nil {
		err = fmt.Errorf("unable to convert response body - %s", jErr)
	}
	return
}
