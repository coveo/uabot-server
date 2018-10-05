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
	queueLength    = flag.Int("queue-length", 100, "Length of the queue of workers")
	port           = flag.String("port", "8080", "Server port")
	sslport        = flag.String("sslport", "8480", "SSL Server port")
	routinesPerCPU = flag.Int("routinesPerCPU", 2, "Maximum number of routine per CPU")
	silent         = flag.Bool("silent", false, "dump the Info prints")
)

const (
	MINIMUMQUEUELENGTH int = 1
	MAXIMUMQUEUELENGTH int = 500
	DEFAULTQUEUELENGTH int = 100

	MINIMUMROUTINEPERCPU int = 1
	MAXIMUMROUTINEPERCPU int = 5
	DEFAULTROUTINEPERCPU int = 2
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "https://"+fmt.Sprintf("%v", r.Host)+fmt.Sprintf(":%v", *sslport)+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	flag.Parse()

	if *silent {
		scenariolib.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
	} else {
		scenariolib.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}

	source := rand.NewSource(int64(time.Now().Unix()))
	random := rand.New(source)

	if *queueLength < MINIMUMQUEUELENGTH || *queueLength > MAXIMUMQUEUELENGTH {
		scenariolib.Info.Printf("Queue Length is out of bounds, should be in [%v,%v], will use default value of %v ", MINIMUMQUEUELENGTH, MAXIMUMQUEUELENGTH, DEFAULTQUEUELENGTH)
		*queueLength = DEFAULTQUEUELENGTH
	}

	if *routinesPerCPU < MINIMUMROUTINEPERCPU || *routinesPerCPU > MAXIMUMROUTINEPERCPU {
		scenariolib.Info.Printf("Routine per CPU is out of bounds, should be in [%v,%v], will use default value of %v ", MINIMUMROUTINEPERCPU, MAXIMUMROUTINEPERCPU, DEFAULTROUTINEPERCPU)
		*routinesPerCPU = DEFAULTROUTINEPERCPU
	}

	scenariolib.Info.Printf("Queue Length: %v", *queueLength)
	scenariolib.Info.Printf("SSL Server Port: %v", fmt.Sprintf(":%v", *sslport))
	scenariolib.Info.Printf("Routine per CPU: %v", *routinesPerCPU)

	concurrentGoRoutine := *routinesPerCPU * runtime.NumCPU()
	scenariolib.Info.Printf("Number of workers: %v", concurrentGoRoutine)
	workPool := server.NewWorkPool(concurrentGoRoutine, int32(*queueLength))

	server.Init(workPool, random)
	routerHttps := server.NewRouter()

	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%v", *sslport), "server.crt", "server.key", routerHttps))
}
