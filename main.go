package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/gate/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	host := os.Getenv("HOST")
	profile := os.Getenv("PROFILE")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := bodies.NewService(conf.Appname, pubPath, conf.HTTPSPort, servicetype.APX)

	routr, err := droxolite.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewBasicEpoxy(srv, droxolite.GetNoTheme(host, srv.ID, profile))
	routers.Setup(poxy, srv.ID, keyPath)

	err = droxolite.BootSecure(poxy, privPath, conf.HTTPPort)

	if err != nil {
		panic(err)
	}
}
