package routers

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/gate/domains"
)

func Setup(e resins.Epoxi) {
	confDomains, err := domains.LoadSettings()

	if err != nil {
		panic(err)
	}

	router := e.Router()
	for _, v := range *confDomains {
		log.Printf("Building %s: \n", v.Domain)

		sub := router.(*mux.Router).Host(fmt.Sprintf("{subdomain:[a-z]+}%s", v.Domain)).Subrouter()

		for _, sdom := range v.Subdomains {
			handl, err := sdom.SetupMux(e.Service().ID)

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
