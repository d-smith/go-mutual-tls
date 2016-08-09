package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	clientCACert, err := ioutil.ReadFile("../cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      clientCertPool,
	}

	tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}

	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Println("Unable to speak to our server", err)
		return
	}

	defer resp.Body.Close()

	bytes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read the response", err)
		return
	}

	log.Println(string(bytes))
}
