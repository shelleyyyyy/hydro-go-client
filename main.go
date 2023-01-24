package main

import ( // "bytes"
	// "fmt"
	// // // "io/ioutil"
	// // "math/rand"
	// // // "net/http"
	// "encoding/json"
	// "hydro/client/conn"
	// "hydro/client/mqtt"
	// "hydro/client/conn"
	"hydro/client/clean"
	"hydro/client/mqtt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	go mqtt.SendLive()

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("05:00").Do(func() {
		CleanUp()
	})

	s.StartBlocking()

	// CleanUp()

}

func CleanUp() {

	clean.CreateModule("humid")
	clean.CreateModule("temp")
	// time.Sleep(86400 * time.Second)
}
