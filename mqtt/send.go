package mqtt

import(
	"encoding/json"
	"hydro/client/conn"
	"fmt"
	"time"
	"math/rand"
)

type Data struct {
	One int `json:"one"`
	Two int `json:"two"`
	Three int `json:"three"`
}

func SendLive() {
	for true {
		temp := Data{getRandom(), getRandom(), getRandom()}
		humid := Data{getRandom(), getRandom(), getRandom()}
	
		b_temp, err := json.Marshal(temp)
		if(err != nil){
			fmt.Println(err)
			return
		}
	
		b_humid, humid_err := json.Marshal(humid)
		if(humid_err != nil){
			fmt.Println(err)
			return
		}
	
		conn.UpdateData("live_temp", "ij81m47voyokipj", b_temp)
		conn.UpdateData("live_humid", "ihdig8d0pfxz6qa", b_humid)
		time.Sleep(1 * time.Second)
		fmt.Println("SENDING")
	}
}

func SendTable() {
	for true {
		temp := Data{getRandom(), getRandom(), getRandom()}
		humid := Data{getRandom(), getRandom(), getRandom()}
	
		b_temp, err := json.Marshal(temp)
		if(err != nil){
			fmt.Println(err)
			return
		}
	
		b_humid, humid_err := json.Marshal(humid)
		if(humid_err != nil){
			fmt.Println(err)
			return
		}
	
		conn.PostData("temp", b_temp)
		conn.PostData("humid", b_humid)
		time.Sleep(5 * time.Second)
		fmt.Println("SENDING")
	}
}



func getRandom() int {
	return rand.Intn(200)
}
