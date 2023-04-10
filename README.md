## Simple load-balancer in Go.
#### This load-balancer based on round-robin algorithm. 
In the round-robin algorithm, requests to the load balancer will route traffic to the group of servers sequentially.<br /> 
This means if there are 6 servers, the load balancer will send requests to each of the servers one after the other (sequentially),<br />
irrespective of the volume of requests the server is currently processing.<br />
We use a round-robin count to calculate the next server to process a request after the next server is found, <br />
we increment the round-robin count so that it will pick the next server in the sequence when processing the next request. <br />