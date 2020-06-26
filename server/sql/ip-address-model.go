package sql

import (
	"fmt"
	mssql "github.com/denisenkom/go-mssqldb"
	"strings"
)

type IPAddress struct {
	ID              int                    `gorm:"column:ID"`
	NetworkID       int                    `gorm:"column:NetworkID"`
	IPAddressTypeID int                    `gorm:"column:IPAddressTypeID"`
	Address         string                 `gorm:"column:Address"`
	Description     string                 `gorm:"column:Description"`
	CreatedBy       string                 `gorm:"column:CreatedBy"`
	DateCreated     string                 `gorm:"column:DateCreated"`
	ModifiedBy      string                 `gorm:"column:ModifiedBy"`
	DateModified    string                 `gorm:"column:DateModified"`
	IsClaimed       bool                   `gorm:"column:IsClaimed"`
	HardwareUUID    mssql.UniqueIdentifier `gorm:"column:HardwareUUID"`
	IsPublic        bool                   `gorm:"column:IsPublic"`
	IsPrimary       bool                   `gorm:"column:IsPrimary"`
	IsClaimable     bool                   `gorm:"column:IsClaimable"`
}

func (i IPAddress) NakedHardwareUUID() string {
	uuid := i.HardwareUUID
	return strings.ToLower(fmt.Sprintf("%X%X%X%X%X", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]))
}

func (IPAddress) TableName() string {
	return "IPAddress"
}
