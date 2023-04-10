package balancer

import (
	"fmt"
	"github.com/atadzan/load-balancer/server"
	"net/http"
)

type LoadBalancer struct {
	Port            string
	roundRobinCount int
	servers         []server.Server
}

func NewLoadBalancer(port string, servers []server.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) GetAvailableServer() server.Server {
	// dynamic round-robin algorithm
	lbServer := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !lbServer.IsAlive() {
		lb.roundRobinCount++
		lbServer = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
	return lbServer
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetAvailableServer()
	fmt.Printf("Forwarding request to addresses: %v \n", targetServer.GetAddress())
	targetServer.Serve(rw, r)
}
