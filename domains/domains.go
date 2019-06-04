package domains

import (
	"log"
	"net/http"
	"strings"
)

type Subdomains struct {
	subs map[string]http.Handler
}

func newSubdomains() *Subdomains {
	subs := make(map[string]http.Handler)
	return &Subdomains{subs: subs}
}

const (
	token = "token"
	ssl   = "ssl"
	www   = "www"
)

func RegisterSubdomains(instanceID, certPath string) *Subdomains {
	result := newSubdomains()
	result.Add(ssl, sslMuxSetup(certPath))

	confDomains, err := loadSettings()

	if err != nil {
		panic(err)
	}

	for _, v := range *confDomains {
		handl, err := v.SetupMux(instanceID)

		if err != nil {
			log.Printf("Register Subdomains: %s\t%s\n", v.Name, err.Error())
		}

		result.Add(v.Address, handl)
	}

	return result
}

//Add will overwrite a subdomain with the same name with the given handler.
func (s *Subdomains) Add(name string, handler http.Handler) {
	s.subs[name] = handler
}

//ServeHTTP calls the requested subdomains' handler
func (d *Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CertBot requires tests on well-known for SSL Certs
	if strings.Contains(r.URL.String(), "well-known") {
		sslHand, _ := d.subs[ssl]

		sslHand.ServeHTTP(w, r)
		return
	}

	log.Printf("Request: %s", r.URL.RequestURI())
	handleSession(r, w)

	domainParts := strings.Split(r.Host, ".")
	sdomainName := domainParts[0]
	result := d.GetMux(sdomainName)
	result.ServeHTTP(w, r)
}

func (d *Subdomains) GetMux(subdomain string) http.Handler {
	result, ok := d.subs[subdomain]

	if !ok {
		return d.subs[www]
	}

	return result
}

func sslMuxSetup(certPath string) http.Handler {
	sslMux := http.NewServeMux()

	fullCertPath := http.FileSystem(http.Dir(certPath))
	fs := http.FileServer(fullCertPath)
	challengePath := "/.well-known/acme-challenge/"

	sslMux.Handle(challengePath, fs)

	return sslMux
}
