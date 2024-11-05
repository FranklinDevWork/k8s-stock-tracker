# k8s-stock-tracker

This project consists of a Golang backend using the Gin API framework, where it provides one endpoint that will aggregate stock data configured via the environment variables `SYMBOL` and `NDAYS`. 

The project can be run via Docker using the following commands:

```shell
$ docker build -t k8s-stock-tracker .
$ docker run -p 8080:8080 --env API_KEY="<GET_ONE>" k8s-stock-tracker
```

You can also specify the `SYMBOL` and `NDAYS` which defaults to MSFT and 7 days. 

This will allow you to make requests to the webservice on port 8080, e.g. 

```shell
$ curl localhost:8080
{"symbol":"MSFT","number_of_days":7,"average_closing_price":426.51500000000004,"results":[{"1. open":"415.3600","2. high":"416.1600","3. low":"406.3000","4. close":"406.3500","5. volume":"53970981"},{"1. open":"416.1200","2. high":"418.9600","3. low":"413.7501","4. close":"418.7800","5. volume":"14206115"},{"1. open":"422.1800","2. high":"422.4800","3. low":"415.2600","4. close":"418.7400","5. volume":"18900201"},{"1. open":"429.8300","2. high":"433.1190","3. low":"428.5700","4. close":"432.1100","5. volume":"13396364"},{"1. open":"414.8000","2. high":"417.7200","3. low":"412.4456","4. close":"416.8600","5. volume":"18266980"},{"1. open":"461.2200","2. high":"466.4600","3. low":"458.8550","4. close":"466.2500","5. volume":"18196100"}]}
```

## Kubernetes 

This project can also be deployed to a kubernetes cluster using the provided helm charts, first you will need to update `api-secret.example` in the `helm/templates` directory with a base64 encoded api key, and rename the file replacing `.example` with `.yaml`, and stripping out the `-example`.

```
helm install k8s-stock-tracker ./helm
``` 
