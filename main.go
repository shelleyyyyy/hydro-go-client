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
	// "hydro/client/mqtt"
	// "hydro/client/conn"
	"hydro/client/clean"
	"hydro/client/mqtt"
)

func main() {

	// topics := []string{"tmp", "humid"}
	// var subs []po.Sub

	// for _, v := range topics {
	// 	tmp := po.Sub{"192.168.0.101", 1883, v, "init", 0}
	// 	subs = append(subs, tmp)

	// }

	// sings := []*po.Sing{
	// 	{"humid", "val"},
	// 	{"tmp", "val"},
	// 	{"humid", "val"},
	// 	{"tmp", "val"},
	// 	{"humid", "val"},
	// 	{"tmp", "val"},
	// }

	// sub_obj := po.Sub{
	// 	"192.168.0.101",
	// 	1883,
	// 	sings,
	// 	0,
	// }

	// sub_obj.Activate()

	// mqtt.SendLive()

	// go subs[0].Activate()

	// go func() {
	// 	subs[0].Activate()
	// }()

	// for true {

	// }

	// fmt.Println("\n\nFINISHED\n" + subs[1].Payload)

	// go mqtt.SendTable()
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
