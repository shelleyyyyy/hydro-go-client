package main

import (
	// "bytes"
	// "fmt"
	// // // "io/ioutil"
	// // "math/rand"
	"time"
	// // // "net/http"
	// "encoding/json"
	// "hydro/client/conn"
	"hydro/client/mqtt"
	// "hydro/client/conn"
	"hydro/client/clean"

)

func main() {



	go mqtt.SendTable()
	go mqtt.SendLive()
	CleanUp()

	// _, items := conn.GetAllData("temp")

	// fmt.Println(len(items))

	// conn.DeleteCollection("temp")
	// conn.DeleteData("temp", "x7fsre41fk8z094")

	// clean.CreateModule("humid")

	// temp := Data{1, 2, 3}

	// b_temp, err := json.Marshal(temp)
	// if(err != nil){
	// 	fmt.Println(err)
	// 	return
	// }

	// conn.UpdateData("live_temp", "ij81m47voyokipj", b_temp)
	
}

func CleanUp() {

	for true {
		clean.CreateModule("humid")
		clean.CreateModule("temp")
		time.Sleep(86400 * time.Second)
	}
}