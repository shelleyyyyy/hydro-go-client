package clean

import(
	"hydro/client/conn"
	"fmt"
	"encoding/json"
)

type Historical struct {
	S1Avg int `json:"s1avg"`
	S1Max int `json:"s1max"`
	S1Min int `json:"s1min"`
	S2Avg int `json:"s2avg"`
	S2Max int `json:"s2max"`
	S2Min int `json:"s2min"`
	S3Avg int `json:"s3avg"`
	S3Max int `json:"s3max"`
	S3Min int `json:"s3min"`
}

type Hist struct {
	Avg float64 `json:"avg"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func CreateModule(collName string) {
	err, items := conn.GetAllData(collName)

	if err != nil {
		fmt.Println(err)
		return
	}

	s1_slice := []float64{}
	s2_slice := []float64{}
	s3_slice := []float64{}

	for _, v := range items {
		s1_slice = append(s1_slice, v.One)
		s2_slice = append(s2_slice, v.Two)
		s3_slice = append(s3_slice, v.Three)
	}

	hist1 := Hist{
		getAvg(s1_slice), 
		getMin(s1_slice),
		getMax(s1_slice),
	}

	hist2 := Hist{
		getAvg(s2_slice), 
		getMin(s2_slice),
		getMax(s2_slice),
	}

	hist3 := Hist{
		getAvg(s3_slice), 
		getMin(s3_slice),
		getMax(s3_slice),
	}

	b1, err := json.Marshal(hist1)
	if(err != nil){
		fmt.Println(err)
		return
	}
	
	b2, err := json.Marshal(hist2)
	if(err != nil){
		fmt.Println(err)
		return
	}

	b3, err := json.Marshal(hist3)
	if(err != nil){
		fmt.Println(err)
		return
	}

	if collName == "temp" {
		conn.PostData("temp_h_1", b1)
		conn.PostData("temp_h_2", b2)
		conn.PostData("temp_h_3", b3)
		conn.DeleteCollection("temp")
	} else {
		conn.PostData("humid_h_1", b1)
		conn.PostData("humid_h_2", b2)
		conn.PostData("humid_h_3", b3)
		conn.DeleteCollection("humid")
	}
}

func getMax(s []float64) float64 {
	max := 0.0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func getMin(s []float64) float64 {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func getAvg(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum = sum + v
	}
	avg := sum / float64(len(s))
	return avg
}