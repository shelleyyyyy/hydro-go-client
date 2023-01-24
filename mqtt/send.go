package mqtt

import (
	"encoding/json"
	"fmt"
	"hydro/client/conn"
	"hydro/client/po"
	"time"

	"github.com/go-co-op/gocron"
)

type Data struct {
	One   float64 `json:"one"`
	Two   float64 `json:"two"`
	Three float64 `json:"three"`
}

func SendLive() {

	sings := []*po.Sing{
		{"s1/humid", 0},
		{"s1/tmpf", 0},
		{"s2/humid", 0},
		{"s2/tmpf", 0},
		{"s3/humid", 0},
		{"s3/tmpf", 0},
	}

	sub_obj := po.Sub{"192.168.0.101", 1883, sings, 0}

	go sub_obj.Activate()

	s := gocron.NewScheduler(time.UTC)

	s.Every(2).Seconds().Do(func() {
		temp := Data{sub_obj.Topics[1].Payload, sub_obj.Topics[3].Payload, sub_obj.Topics[5].Payload}
		humid := Data{sub_obj.Topics[0].Payload, sub_obj.Topics[2].Payload, sub_obj.Topics[4].Payload}

		b_temp, err := json.Marshal(temp)
		if err != nil {
			fmt.Println(err)
			return
		}

		b_humid, humid_err := json.Marshal(humid)
		if humid_err != nil {
			fmt.Println(err)
			return
		}

		conn.UpdateData("live_temp", "ij81m47voyokipj", b_temp)
		conn.UpdateData("live_humid", "ihdig8d0pfxz6qa", b_humid)
		time.Sleep(2 * time.Second)
		fmt.Println("SENDING")
	})

	s.Every("5m").Do(func() {
		temp := Data{sub_obj.Topics[1].Payload, sub_obj.Topics[3].Payload, sub_obj.Topics[5].Payload}
		humid := Data{sub_obj.Topics[0].Payload, sub_obj.Topics[2].Payload, sub_obj.Topics[4].Payload}

		b_temp, err := json.Marshal(temp)
		if err != nil {
			fmt.Println(err)
			return
		}

		b_humid, humid_err := json.Marshal(humid)
		if humid_err != nil {
			fmt.Println(err)
			return
		}

		conn.PostData("temp", b_temp)
		conn.PostData("humid", b_humid)
		fmt.Println("SENDING")
	})

	s.StartAsync()

	// for true {
	// 	temp := Data{sub_obj.Topics[1].Payload, sub_obj.Topics[3].Payload, sub_obj.Topics[5].Payload}
	// 	humid := Data{sub_obj.Topics[0].Payload, sub_obj.Topics[2].Payload, sub_obj.Topics[4].Payload}

	// 	b_temp, err := json.Marshal(temp)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	b_humid, humid_err := json.Marshal(humid)
	// 	if humid_err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	conn.UpdateData("live_temp", "ij81m47voyokipj", b_temp)
	// 	conn.UpdateData("live_humid", "ihdig8d0pfxz6qa", b_humid)
	// 	conn.PostData("temp", b_temp)
	// 	conn.PostData("humid", b_humid)
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("SENDING")
	// }
}

// func SendTable() {
// 	sings := []*po.Sing{
// 		{"s1/humid", 0},
// 		{"s1/tmpf", 0},
// 		{"s2/humid", 0},
// 		{"s2/tmpf", 0},
// 		{"s3/humid", 0},
// 		{"s3/tmpf", 0},
// 	}

// 	sub_obj := po.Sub{"192.168.0.101", 1883, sings, 0}

// 	go sub_obj.Activate()

// 	temp := Data{sub_obj.Topics[1].Payload, sub_obj.Topics[3].Payload, sub_obj.Topics[5].Payload}
// 	humid := Data{sub_obj.Topics[0].Payload, sub_obj.Topics[2].Payload, sub_obj.Topics[4].Payload}

// 	b_temp, err := json.Marshal(temp)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	b_humid, humid_err := json.Marshal(humid)
// 	if humid_err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	conn.PostData("temp", b_temp)
// 	conn.PostData("humid", b_humid)
// 	fmt.Println("SENDING")
// }

// func getRandom() int {
// 	return rand.Intn(200)
// }
