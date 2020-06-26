package service

import (
	"github.com/Molsbee/jarvis/common_model"
	"github.com/Molsbee/jarvis/server/config"
	"github.com/Molsbee/jarvis/server/sql"
	"log"
)

func FindIPAddress(addressString string) (ipAddress common_model.IPAddress) {
	for _, dc := range config.UserConfig.DataCenters {
		client, err := sql.NewSQLClient(dc.MSSQL)
		if err != nil {
			log.Printf("failed to create a database connection dc: %s")
		}

		address, _ := client.FindIPAddress(addressString)
		client.Close()

		if address.Address == addressString {
			ipAddress = common_model.IPAddress{
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
			break
		}
	}

	return
}

//func FindServerByIPAddress(ipAddress string) {
//
//}
