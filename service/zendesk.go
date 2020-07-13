package service

import (
	"encoding/json"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/model"
	"net/http"
)

const zendeskAPI = "https://t3n.zendesk.com/api/v2/"

func GetZendeskTickets() ([]model.ZendeskTicket, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf("%s/views/incoming/tickets.json", zendeskAPI), nil)
	zendeskCredentials := config.GetConfig().Zendesk
	request.SetBasicAuth(zendeskCredentials.Username, zendeskCredentials.Password)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from zendesk api - err (%s)", err)
	}

	t := struct {
		Tickets []model.ZendeskTicket `json:"tickets"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, fmt.Errorf("failed to decode response to object - err (%s)", err)
	}

	return t.Tickets, nil
}
