package main

import (
	"flag"
	"fmt"
	"github.com/coveo/uabot-server/server"
	"github.com/coveo/uabot/scenariolib"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	queueLength = flag.Int("queue-length", 100, "Length of the queue of workers")
	port        = flag.String("port", "8080", "Server port")
	sslport        = flag.String("sslport", "443", "Server port")
	routinesPerCPU = flag.Int("routinesPerCPU", 2, "Maximum number of routine per CPU")
	silent = flag.Bool("silent", false, "dump the Info prints")
)

func main() {
	flag.Parse()

	if *silent {
		scenariolib.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
	} else {
		scenariolib.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}

	scenariolib.Info.Printf("Queue Length: %v", *queueLength)
	scenariolib.Info.Printf("Server Port: %v", *port)
	scenariolib.Info.Printf("Server SSL Port: %v", *sslport)

	source := rand.NewSource(int64(time.Now().Unix()))
	random := rand.New(source)

	concurrentGoRoutine := *routinesPerCPU * runtime.NumCPU()
	scenariolib.Info.Printf("Number of workers: %v", concurrentGoRoutine)
	workPool := server.NewWorkPool(concurrentGoRoutine, int32(*queueLength))

	server.Init(workPool, random)
	router := server.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), router))
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%v", *sslport), "server.crt", "server.key", router))
}
