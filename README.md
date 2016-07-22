# go-echo

Simple golang HTTP echo server

## Quickstart

    go run echo.go
    
Click [http://localhost:8080/hello](http://localhost:8080/hello?q=world) 
and the response should be

    {
        "method": "GET",
        "query": {
            "q": "world"
        }
        "path": "/hello",
        "body": ""
    }

For POST request the body will be set in the response

