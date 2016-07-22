# go-echo

Simple golang HTTP echo server

## Quickstart

    go run echo.go -p 8080
    
Click [http://localhost:8080/hello](http://localhost:8080/hello?q=world) 
and the response should be
    
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


