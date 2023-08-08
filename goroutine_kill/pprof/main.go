package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"runtime/debug"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	debug.PrintStack()
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(10)
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case <-ch:
				time.Sleep(time.Second * 2)
			}
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/{category}", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Fatal(http.ListenAndServe(":6060", r))
}
