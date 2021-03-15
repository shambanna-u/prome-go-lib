# prome-go-lib
Instrumentation of go-lang code with Prometheus library 

- This code will call the https://httpstat.us/503 and https://httpstat.us/200 endpoints and get response code and response time.

- It exposes the metrics in prometheus format on http://localhost:2222/metrics.

- kube directory contains kuberenets resourece files.


Run the following command to create the resources.

    $ kubectl apply -f kube/


### Services are exported with following urls

- Prometheus: http://{{node_ip}}:30100/

- Grafana: http://{{node_ip}}:30200/

- Applicaion: http://{{node_ip}}:30300/




