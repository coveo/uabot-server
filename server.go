package main

import (
	"flag"
	"fmt"
	"github.com/adambbolduc/uabot/server"
	"github.com/erocheleau/uabot/scenariolib"
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
)

func main() {
	flag.Parse()

	scenariolib.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	scenariolib.Info.Printf("Queue Length: %v", *queueLength)
	scenariolib.Info.Printf("Server Port: %v", *port)

	source := rand.NewSource(int64(time.Now().Unix()))
	random := rand.New(source)

	concurrentGoRoutine := runtime.NumCPU()
	workPool := server.NewWorkPool(concurrentGoRoutine, int32(*queueLength))

	server.Init(workPool, random)
	router := server.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), router))
}