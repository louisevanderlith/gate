package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"

	"fmt"

	"github.com/louisevanderlith/gate/domains"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
)

func main() {
	mode := os.Getenv("RUNMODE")
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	privName := os.Getenv("PRIVATEKEY")
	pubPath := path.Join(keyPath, pubName)

	appName := beego.BConfig.AppName
	// Register with router
	srv := mango.NewService(mode, appName, pubPath, enums.APP)

	httpsPort := beego.AppConfig.String("httpsport")

	err := srv.Register(httpsPort)

	if err != nil {
		panic(err)
	}

	httpPort := beego.AppConfig.String("httpport")
	setupHost(httpPort, httpsPort, srv.ID, keyPath, pubName, privName)
}

func setupHost(httpPort, httpsPort, instanceID, certPath, publicKey, privateKey string) {
	subs := domains.RegisterSubdomains(instanceID, certPath)

	go serveHTTP2(subs, httpsPort, certPath, publicKey, privateKey)

	err := http.ListenAndServe(":"+httpPort, http.HandlerFunc(redirectTLS))

	if err != nil {
		panic(err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	moveURL := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
	http.Redirect(w, r, moveURL, http.StatusPermanentRedirect)
}

func serveHTTP2(domains *domains.Subdomains, httpsPort, certPath, publicKey, privateKey string) {
	publicKeyPem := readBlocks(path.Join(certPath, publicKey))
	privateKeyPem := readBlocks(path.Join(certPath, privateKey))
	cert, err := tls.X509KeyPair(publicKeyPem, privateKeyPem)

	if err != nil {
		panic(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	srv := &http.Server{
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Addr:         ":" + httpsPort,
		Handler:      domains,
	}

	log.Println("Listening...")

	err = srv.ListenAndServeTLS("", "")

	if err != nil {
		panic(err)
	}
}

func readBlocks(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return file
}
