package model

type VM struct {
	ID      string `json:"id"`
	Archive struct {
		StorageMB int `json:"storageMB"`
	} `json:"archive"`
	ChangeInfo      ChangeInfo
	Credentials     Credential
	CustomFields    []string `json:"customFields"`
	Description     string   `json:"description"`
	MaintenanceMode bool     `json:"maintenanceMode"`
	ManagedOS       bool     `json:"managedOS"`
	Name            string   `json:"name"`
	DisplayName     string   `json:"displayName"`
	OS              string   `json:"os"`
	Status          string   `json:"status"`
	Storage         string   `json:"storage"`
	Template        bool     `json:"template"`
	Type            string   `json:"type"`
	AccountID       string   `json:"accountID"`
	LocationID      string   `json:"locationID"`
	GroupUUID       string   `json:"groupUuid"`
}

type Credential struct {
	UserName string `json:"userName"`
	Password struct {
		EncryptedPassword string `json:"encryptedValue"`
		EncryptionSeed    string `json:"encryptionSeed"`
		EncryptionVersion string `json:"encryptionVersion"`
	} `json:"password"`
}

type ServerConfiguration struct {
	ID         string `json:"id"`
	AccountID  string `json:"accountID"`
	LocationID string `json:"locationID"`
	Name       string `json:"name"`
	Host       struct {
		Type               string   `json:"type"`
		DataStores         []string `json:"dataStores"`
		HostName           string   `json:"hostName"`
		ManagementHostName string   `json:"managementHostName"`
		ManagementLinks    []struct {
			Name string `json:"name"`
			URI  string `json:"uri"`
		}
		RegisteredName string `json:"registeredName"`
	}
	Network struct {
		IPAddresses []string `json:"ipAddresses"`
		DNSName     string   `json:"dnsName"`
	}
	PowerState string `json:"powerState"`
	ChangeInfo ChangeInfo
	Hardware   struct {
		CoresPerSocket int `json:"coresPerSocket"`
		Disks          []struct {
			ID         string `json:"id"`
			CapacityMB int    `json:"capacityMB"`
			Partitions []struct {
				CapacityMB int    `json:"capacityMB"`
				Path       string `json:"path"`
			} `json:"partitions"`
		}
		MemoryMB  int `json:"memoryMB"`
		Sockets   int `json:"sockets"`
		StorageMB int `json:"storageMB"`
	}
	Snapshots []interface{}
}

type ChangeInfo struct {
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	ModifiedBy   string `json:"modifiedBy"`
	ModifiedDate string `json:"modifiedDate"`
}
