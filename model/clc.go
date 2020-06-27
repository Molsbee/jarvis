package model

import "fmt"

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
	return fmt.Sprintf(`Address:       %s
Location:      %s
ID:            %d
NetworkID:     %d
HardwareUUID:  %s
IsPublic:      %v
IsClaimed:     %v
IsClaimable:   %v`, i.Address, i.Location, i.ID, i.NetworkID, i.HardwareUUID, i.IsPublic, i.IsClaimed, i.IsClaimable)
}
