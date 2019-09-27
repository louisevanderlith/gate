package domains

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/do"
)

type Subdomain struct {
	Address string
	Name    string
}

type DomainSetting struct {
	Domain     string
	Subdomains []Subdomain
}

type Settings []*DomainSetting

//LoadSettings returns the data contained in the 'domains.json' config file.
func LoadSettings() (*Settings, error) {
	dbConfPath := droxolite.FindFilePath("domains.json", "conf")
	content, err := ioutil.ReadFile(dbConfPath)

	if err != nil {
		return nil, err
	}

	settings := &Settings{}
	err = json.Unmarshal(content, settings)

	if err != nil {
		return nil, err
	}

	return settings, nil
}

func SetupMux(instanceID, subName string) (http.Handler, error) {
	rawURL, err := do.GetServiceURL(instanceID, subName, false)

	if err != nil {
		return nil, err
	}

	vshost, err := url.Parse(rawURL)

	if err != nil {
		return nil, err
	}
	
	log.Printf("Proxy: %v\n", vshost)
	return httputil.NewSingleHostReverseProxy(vshost), nil
}
