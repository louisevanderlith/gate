package routers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/gate/domains"
)

func Setup(e resins.Epoxi, domain string) {
	domain = strings.TrimSuffix(domain, "/")

	srvc := e.Service()
	log.Printf("Building %s Profile %s: \n", domain, srvc.Profile)

	rtr := e.Router().(*mux.Router)
	subdoms := make(map[string]string)

	code, err := do.GET("", &subdoms, srvc.ID, "Router.API", "applicants", srvc.Profile)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	if code != http.StatusOK {
		log.Println(code)
		panic("unable to load applicants")
	}

	for sname, sdom := range subdoms {
		handl, err := domains.SetupMux(srvc.ID, sname)

		if err != nil {
			log.Printf("Register Subdomains: %s\t%s\n", sname, err.Error())
		}

		log.Printf("Setup: %s\n", sdom)

		if sdom == "www" {
			dchild := rtr.Host(strings.Replace(domain, ".", "", 1)).Subrouter()
			dchild.PathPrefix("/").Handler(handl)
		}

		child := rtr.Host(fmt.Sprintf("%s%s", sdom, domain)).Subrouter()
		child.PathPrefix("/").Handler(handl)
		child.Use(domains.HandleSession)
	}
}
