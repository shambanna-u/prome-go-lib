# prome-go-lib
Instrumentation of go-lang code with Prometheus library.

* This code will call the https://httpstat.us/503 and https://httpstat.us/200 endpoints and get response code and response time.
* It exposes the metrics in prometheus format on http://{{host_ip}}:2222/metrics .
* Github actions are also configured to run build and test cases. 
* Automatic PR merge will happen once the test and build is passed.
* Webhook is configured to the docker hub.
* Any code push or PR merge to the main branch will trigger the build.
* Image will be available in docker hub shambuu/prome:latest, we can pull the latest image from the following command.
```
  $ docker pull shambuu/prome:latest
```
* kube directory contains kubernetes resource files. 

Run the following command to create the resources.
```
   $ kubectl apply -f kube/
```
![Image of kubestatus](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/kube-status.PNG)
* Services are exported with following urls

  - Prometheus: http://{{node_ip}}:30100/
  - Grafana: http://{{node_ip}}:30200/
  - Application: http://{{node_ip}}:30300/metrics

* If you check the application endpoint output will be similar to this

![Image of metrics](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/metrics.PNG)
* Prometheus endpoints are already configured to get from application by using service endpoint.

![Image of targe](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/target.PNG)

* Run the PQL to get the metrics

![Image of pql](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/prome-ui.PNG)

* Grafana dashboard is already configured with Data source


![Image of grafana](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/datasource.PNG)

* dashboard.json is added in this repo, import the dashboard.

![Image of output](https://github.com/shambanna-u/prome-go-lib/blob/main/.img/dashboard.PNG)

