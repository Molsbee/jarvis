package elasticsearch

import "fmt"

type ServerResponse struct {
	Name        string
	PowerState  string
	AccountID   string
	LocationID  string
	VSphere     string
	IPAddresses []string
	Credentials Credential
	OS          string
	Type        string
	GroupUUID   string
}

func (s ServerResponse) String() string {
	return fmt.Sprintf(
		`	Name: %s
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
	GroupUUID: %s`, s.Name, s.PowerState, s.AccountID, s.LocationID, s.VSphere, s.IPAddresses,
		s.Credentials.UserName, s.Credentials.Password.EncryptedPassword, s.Credentials.Password.EncryptionSeed,
		s.Credentials.Password.EncryptionVersion, s.OS, s.Type, s.GroupUUID)
}
