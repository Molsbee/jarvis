package elasticsearch

import (
	"fmt"
	"strings"
)

type main struct {
	URL string
}

type Environment interface {
	Name() string
	URL() string
}

type environment struct {
	name string
	url  string
}

func (e environment) Name() string {
	return e.name
}

func (e environment) URL() string {
	return e.url
}

var (
	DEV  = environment{name: "dev", url: "http://search-lb1.t3dev.dom:1337/main/_search?pretty=true"}
	PPE  = environment{name: "ppe", url: "http://search-lb2.t3ppe.dom:1337/main/_search?pretty=true"}
	PROD = environment{name: "prod", url: "http://search-uc1.t3n.dom:1337/main/_search?pretty=true"}
)

func GetEnvironment(env string) (e Environment, err error) {
	switch env {
	case DEV.name:
		e = DEV
		return
	case PPE.name:
		e = PPE
		return
	case PROD.name:
		e = PROD
		return
	}

	err = fmt.Errorf("%s is an unsupported environment - supported values (%s, %s, %s)", env, DEV.name, PPE.name, PROD.name)
	return
}

func Main(env Environment) *main {
	return &main{
		URL: env.URL(),
	}
}

func (m *main) GetServerDetails(name string) (s ServerResponse, err error) {
	vm, vmErr := m.getVirtualMachineInfo(name)
	if vmErr != nil {
		err = fmt.Errorf("unable to pull vm document for server name: (%s) - %s", name, vmErr)
		return
	}

	serverConfiguration, sErr := m.getServerConfiguration(name)
	if sErr != nil {
		err = fmt.Errorf("unable to pull server configuration document for server name: (%s) - %s", name, vmErr)
		return
	}

	s = ServerResponse{
		Name:        vm.Name,
		PowerState:  serverConfiguration.PowerState,
		AccountID:   vm.AccountID,
		LocationID:  vm.LocationID,
		VSphere:     serverConfiguration.Host.ManagementLinks[0].URI,
		IPAddresses: serverConfiguration.Network.IPAddresses,
		Credentials: vm.Credentials,
		OS:          vm.OS,
		Type:        vm.Type,
		GroupUUID:   vm.GroupUUID,
	}
	return
}

func (m *main) GetServerDetailsByHardwareUUID(uuid string) (s ServerResponse, err error) {
	vm, vmErr := m.getVirtualMachineInfoByHardwareUUID(uuid)
	if vmErr != nil {
		err = fmt.Errorf("unable to pull vm document for server id: (%s) - %s", uuid, vmErr)
		return
	}
	serverConfiguration, sErr := m.getServerConfigurationByHardwareUUID(uuid)
	if sErr != nil {
		err = fmt.Errorf("unable to pull server configuration document for server id: (%s) - %s", uuid, vmErr)
		return
	}

	s = ServerResponse{
		Name:        vm.Name,
		PowerState:  serverConfiguration.PowerState,
		AccountID:   vm.AccountID,
		LocationID:  vm.LocationID,
		VSphere:     serverConfiguration.Host.ManagementLinks[0].URI,
		IPAddresses: serverConfiguration.Network.IPAddresses,
		Credentials: vm.Credentials,
		OS:          vm.OS,
		Type:        vm.Type,
		GroupUUID:   vm.GroupUUID,
	}
	return
}

func (m *main) getVirtualMachineInfo(name string) (v VM, err error) {
	response, err := post(m.URL, fmt.Sprintf(`{
			"filter": {
    			"term": {"couchbaseDocument.doc.vm.name": "%s"}
			}
		}`, strings.ToLower(name)))
	if err != nil {
		return
	}

	v = response.Hits.Hits[0].Source.Doc.VM
	return
}

func (m *main) getVirtualMachineInfoByHardwareUUID(uuid string) (v VM, err error) {
	response, err := post(m.URL, fmt.Sprintf(`{
			"filter": {
    			"term": {"couchbaseDocument.doc.vm.id": "%s"}
			}
		}`, strings.ToLower(uuid)))
	if err != nil {
		return
	}

	v = response.Hits.Hits[0].Source.Doc.VM
	return
}

func (m *main) getServerConfiguration(name string) (s ServerConfiguration, err error) {
	response, err := post(m.URL, fmt.Sprintf(`{
			"filter": {
    			"term": {"couchbaseDocument.doc.serverConfiguration.name": "%s"}
			}
		}`, strings.ToLower(name)))
	if err != nil {
		return
	}

	s = response.Hits.Hits[0].Source.Doc.ServerConfiguration
	return
}

func (m *main) getServerConfigurationByHardwareUUID(uuid string) (s ServerConfiguration, err error) {
	response, err := post(m.URL, fmt.Sprintf(`{
			"filter": {
    			"term": {"couchbaseDocument.doc.serverConfiguration.id": "%s"}
			}
		}`, strings.ToLower(uuid)))
	if err != nil {
		return
	}

	s = response.Hits.Hits[0].Source.Doc.ServerConfiguration
	return
}
