# go-echo

Simple golang HTTP echo server. Also has a Prometheus metrics endpoint at http://localhost:2112/metrics.

Can be used as a [container image](https://hub.docker.com/r/cirri/echo):

```
docker run -dit -p 80:80 -p 2112:2112 cirri/echo:latest
```


For use with `cirri` --infra (using caddy, caddy-docker-proxy and some DNS magic):

```
docker run -dit --network cirri_proxy \
		--name echo_dev \
			cirri/echo:latest

maps the echo to `https://echo_dev.${STACKDOMAIN}` and the metrics to `https://echo_dev.${STACKDOMAIN}/metrics` using the `virtual.port` and `virtual.metrics` labels in the image.
```

## Quickstart

```
    go run echo.go -p 8080
```
    
Click [http://localhost:8080/hello](http://localhost:8080/hello?q=world) 
and the response should be similar to this:

```
    {
        "Body": {
            "String": ""
        },
        "ContentLength": 0,
        "Header": {
            "Accept": [
                "*/*"
            ],
            "Accept-Encoding": [
                "gzip, deflate"
            ],
            "Connection": [
                "keep-alive"
            ],
            "User-Agent": [
                "HTTPie/0.9.4"
            ]
        },
        "Host": "localhost:8080",
        "Method": "GET",
        "Proto": "HTTP/1.1",
        "ProtoMajor": 1,
        "ProtoMinor": 1,
        "RemoteAddr": "[::1]:53697",
        "RequestURI": "/hello?q=world",
        "Trailer": null,
        "TransferEncoding": null,
        "URL": {
            "Fragment": "",
            "Host": "",
            "Opaque": "",
            "Path": "/hello",
            "RawPath": "",
            "RawQuery": "q=world",
            "Scheme": "",
            "User": null
        }
    }
```

