package main

import (
	"fmt"
	"github.com/atadzan/load-balancer/balancer"
	"github.com/atadzan/load-balancer/server"
	"net/http"
	"os"
)

func main() {
	servers := []server.Server{
		server.NewSimpleServer("https://bloomberg.com"),
		server.NewSimpleServer("https://google.com"),
		server.NewSimpleServer("https://uber.com"),
	}

	lb := balancer.NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request")
		lb.ServeProxy(rw, r)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Server is redirecting on the localhost at: %s\n", lb.Port)
	err := http.ListenAndServe(":"+lb.Port, nil)
	if err != nil {
		fmt.Sprintf("error: %v\n", err)
		os.Exit(1)
	}

}
