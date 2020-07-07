package clc

import (
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/model"
	"github.com/Molsbee/jarvis/service/clc/sql"
	"log"
	"time"
)

func FindIPAddress(ip string) (ipAddress *model.IPAddress) {
	ch := make(chan model.IPAddress)
	for _, dc := range config.DataCenters {
		go findIPAddress(ip, dc, ch)
	}

loop:
	for i := 0; i < 30; i++ {
		select {
		case a := <-ch:
			ipAddress = &a
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	return
}

func findIPAddress(ip string, dc config.DataCenter, ch chan model.IPAddress) {
	client, err := sql.NewSQLClient(dc.MSSQL)
	if err != nil {
		log.Printf("failed to create a database connection dc: %s", dc.Name)
	}

	address, _ := client.FindIPAddress(ip)
	defer client.Close()

	if address.Address == ip {
		ipAddress := model.IPAddress{
			Location:     dc.Name,
			ID:           address.ID,
			NetworkID:    address.NetworkID,
			Address:      address.Address,
			Description:  address.Description,
			HardwareUUID: address.NakedHardwareUUID(),
			IsPublic:     address.IsPublic,
			IsClaimed:    address.IsClaimed,
			IsClaimable:  address.IsClaimable,
		}
		ch <- ipAddress
	}
}
