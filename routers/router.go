package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/gate/domains"
)

/*
func Boot(httpsPort, httpPort int) error {
	r := mux.Router()
	r.HandleFunc("/ping", BasicHandler)

	log.Fata(http.ListenAn)
}*/

func Setup(e resins.Epoxi, instanceID, certPath string) {
	confDomains, err := domains.LoadSettings()

	if err != nil {
		panic(err)
	}

	router := e.Router()
	for _, v := range *confDomains {
		log.Printf("Building %s: \n", v.Domain)
		//SSL
		fullCertPath := http.FileSystem(http.Dir(certPath))
		fs := http.FileServer(fullCertPath)
		challengePath := "/.well-known/acme-challenge/"

		sub := router.(*mux.Router).Host(fmt.Sprintf("{subdomain:[a-z]+}%s", v.Domain)).Subrouter()
		sub.Handle(challengePath, fs)

		for _, sdom := range v.Subdomains {
			handl, err := sdom.SetupMux(instanceID)

			if err != nil {
				log.Printf("Register Subdomains: %s\t%s\n", sdom.Name, err.Error())
			}

			log.Printf("Setup: %s\n", sdom.Address)
			child := sub.Host(fmt.Sprintf("%s%s", sdom.Address, v.Domain)).Subrouter()
			child.PathPrefix("/").Handler(handl)
			child.Use(domains.HandleSession)
		}
	}
}
