package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/gate/routers"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	//host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)
	privPath := path.Join(keyPath, privName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPSPort, servicetype.APX)

	err = srv.Register()

	if err != nil {
		panic(err)
	}

	poxy := droxolite.NewEpoxy(srv)
	routers.Setup(poxy, srv.ID, keyPath)

	err = poxy.BootSecure(privPath, conf.HTTPPort)

	if err != nil {
		panic(err)
	}
}
