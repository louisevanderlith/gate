package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/gate/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	host := os.Getenv("HOST")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	httpsport, _ := strconv.Atoi(os.Getenv("HTTPSPORT"))
	appName := os.Getenv("APPNAME")
	profile := os.Getenv("PROFILE")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	// Register with router
	srv := bodies.NewService(appName, profile, pubPath, host, httpsport, servicetype.APX)

	routr, err := do.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, profile))
	routers.Setup(poxy, host)

	err = droxolite.BootSecure(poxy, privPath, httpport)

	if err != nil {
		panic(err)
	}
}
