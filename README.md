# prome-go-lib
Instrumentation of go-lang code with Prometheus library.

* This code will call the https://httpstat.us/503 and https://httpstat.us/200 endpoints and get response code and response time.
* It exposes the metrics in prometheus format on http://{{host_ip}}:2222/metrics .
* Github actions also configured to run buld and test cases. 
* Autometic PR merge will happens once test and build is passed.
* Webhook is configured to the docker hub.
* Any code push or PR merge to main branch will triggre the build.
* Image will be availble in docker hub shambuu/prome:latest, we can pull the latest image from following command.
```
  $ docker pull shambuu/prome:latest
```
* kube directory contains kuberenets resourece files. 

Run the following command to create the resources.
```
   $ kubectl apply -f kube/
```
### Services are exported with following urls

- Prometheus: http://{{node_ip}}:30100/

- Grafana: http://{{node_ip}}:30200/

- Applicaion: http://{{node_ip}}:30300/
