package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/google/gops/agent"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"index", "/", index,
	},
	Route{
		"everything", "/{everything:.*}", everything,
	},
}

type RequestBody struct {
	String string
	// Other types?
}

// Request is the same as http.Request minus the bits that break json.Marshall
type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           http.Header
	Body             RequestBody
	ContentLength    int64
	TransferEncoding []string
	Host             string
	//Form url.Values
	//PostForm url.Values
	//MultipartForm *multipart.Form
	Trailer    http.Header
	RemoteAddr string
	RequestURI string
	//TLS *tls.ConnectionState
	Cookies []*http.Cookie
}

const megabytes = 1048576

func echo(w http.ResponseWriter, r *http.Request) {
	e := Request{}
	e.Method = r.Method
	e.URL = r.URL
	e.Proto = r.Proto
	e.ProtoMajor = r.ProtoMajor
	e.ProtoMinor = r.ProtoMinor
	e.Header = r.Header
	e.Body = RequestBody{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1*megabytes))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	e.Body.String = string(body)
	e.ContentLength = r.ContentLength
	e.TransferEncoding = r.TransferEncoding
	e.Host = r.Host
	e.Trailer = r.Trailer
	e.RemoteAddr = r.RemoteAddr
	e.RequestURI = r.RequestURI
	e.Cookies = r.Cookies()

	// TODO detect if we're proxied and change the remoteip to the next endpoint along
	remoteAddr := r.RemoteAddr
	if v := r.Header.Get("X-Forwarded-For"); v != "" {
		remoteAddr = v
	}
	requestCount.With(prometheus.Labels{"remoteip": remoteAddr, "host": r.Host}).Inc()

	b, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(b))
}

func index(w http.ResponseWriter, r *http.Request) {
	echo(w, r)
}

func everything(w http.ResponseWriter, r *http.Request) {
	echo(w, r)
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)
		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func main() {
	port := flag.Int("p", 80, "Port")
	flag.Parse()

	go runMetrics()

	router := newRouter()
	log.Println(fmt.Sprintf("Listening on port %v...", *port))
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), router))
}

////// prom metrics
var (
	requestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "echo_request_count",
			Help: "The number of requests to echo",
		},
		[]string{"remoteip", "host"},
	)
)

func runMetrics() {
	fmt.Printf("Starting gops metrics\n")
	if err := agent.Listen(agent.Options{}); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	Port := ":2112"
	fmt.Printf("Starting Prometheus metrics endpoint on Port %s\n", Port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Printf("Fail to listen to Port %s error: %v, exiting\n", Port, err)
		os.Exit(1)
	}
}
