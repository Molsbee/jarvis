package elasticsearch

import (
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/model"
	"strings"
)

type main struct {
	URL string
}

func Main(env config.Environment) *main {
	return &main{
		URL: env.ElasticSearchMainURL(),
	}
}

func (m *main) GetServerDetails(name string) (s model.ServerResponse, err error) {
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

	s = model.NewServerResponse(vm, serverConfiguration)
	return
}

func (m *main) GetServerDetailsByHardwareUUID(uuid string) (s model.ServerResponse, err error) {
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

	s = model.NewServerResponse(vm, serverConfiguration)
	return
}

func (m *main) getVirtualMachineInfo(name string) (v model.VM, err error) {
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

func (m *main) getVirtualMachineInfoByHardwareUUID(uuid string) (v model.VM, err error) {
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

func (m *main) getServerConfiguration(name string) (s model.ServerConfiguration, err error) {
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

func (m *main) getServerConfigurationByHardwareUUID(uuid string) (s model.ServerConfiguration, err error) {
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
