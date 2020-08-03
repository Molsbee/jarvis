package model

import (
	"fmt"
)

type ServerResponse struct {
	Name              string
	Status            string
	PowerState        string
	AccountID         string
	LocationID        string
	VSphere           string
	IPAddresses       []string
	UserName          string
	EncryptedPassword string
	EncryptionSeed    string
	EncryptionVersion string
	OS                string
	Type              string
	GroupUUID         string
}

func NewServerResponse(vm VM, configuration ServerConfiguration) ServerResponse {
	serverResponse := ServerResponse{
		Name:              vm.Name,
		Status:            vm.Status,
		PowerState:        configuration.PowerState,
		AccountID:         vm.AccountID,
		LocationID:        vm.LocationID,
		IPAddresses:       configuration.Network.IPAddresses,
		UserName:          vm.Credentials.UserName,
		EncryptedPassword: vm.Credentials.Password.EncryptedPassword,
		EncryptionSeed:    vm.Credentials.Password.EncryptionSeed,
		EncryptionVersion: vm.Credentials.Password.EncryptionVersion,
		OS:                vm.OS,
		Type:              vm.Type,
		GroupUUID:         vm.GroupUUID,
	}

	if len(configuration.Host.ManagementLinks) != 0 {
		serverResponse.VSphere = configuration.Host.ManagementLinks[0].URI
	}

	return serverResponse
}

func (s ServerResponse) String() string {
	return fmt.Sprintf(
		`	Name: %s
	Status: %s
	PowerState: %s
	AccountID: %s
	LocationID: %s
	VSphere: %s
	IPAddresses: %v
	Credentials:
		UserName: %s
		EncryptedPassword: %s
		EncryptionSeed: %s
		EncryptionVersion: %s
	OS: %s
	Type: %s
	GroupUUID: %s`, s.Name, s.Status, s.PowerState, s.AccountID, s.LocationID, s.VSphere, s.IPAddresses,
		s.UserName, s.EncryptedPassword, s.EncryptionSeed,
		s.EncryptionVersion, s.OS, s.Type, s.GroupUUID)
}

type IPAddress struct {
	Location     string `json:"location"`
	ID           int    `json:"id"`
	NetworkID    int    `json:"networkId"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	HardwareUUID string `json:"hardwareUUID"`
	IsPublic     bool   `json:"isPublic"`
	IsClaimed    bool   `json:"isClaimed"`
	IsClaimable  bool   `json:"isClaimable"`
}

func (i IPAddress) String() string {
	return fmt.Sprintf(
		`	Address:       %s
	Location:      %s
	ID:            %d
	NetworkID:     %d
	HardwareUUID:  %s
	IsPublic:      %v
	IsClaimed:     %v
	IsClaimable:   %v`, i.Address, i.Location, i.ID, i.NetworkID, i.HardwareUUID, i.IsPublic, i.IsClaimed, i.IsClaimable)
}
