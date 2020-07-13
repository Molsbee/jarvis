package sql

import (
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/jinzhu/gorm"
	"net/url"
)

type Client struct {
	DB *gorm.DB
}

func NewSQLClient(uri string) (c *Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to create database connection to sql: %s", uri)
		}
	}()

	query := url.Values{}
	query.Add("database", "T3MAIN")

	domain := config.GetConfig().Domain
	u := url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(fmt.Sprintf("%s\\%s", domain.Name, domain.Username), domain.Password),
		Host:     uri,
		RawQuery: query.Encode(),
	}

	db, gErr := gorm.Open("mssql", u.String())
	if gErr != nil {
		err = fmt.Errorf("error opening %s database connection - %s", uri, gErr)
	}
	db.LogMode(false)

	c = &Client{
		DB: db,
	}

	return
}

func (c *Client) FindIPAddress(ipAddress string) (address IPAddress, err error) {
	err = c.DB.Where("Address = ?", ipAddress).Find(&address).Error
	return
}

func (c *Client) Close() {
	_ = c.DB.Close()
}
