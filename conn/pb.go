package conn

import(
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strconv"
)

type ResponseData struct {
	Page int 
	PerPage int 
	TotalItems int 
	TotalPages int 
	Items []Item
}

type Item struct {
	CollectionID string
	CollectionName string
	Created string
	ID string
	One float64
	Two float64
	Three float64
	Updated string
}

func PostData(collID string, jsonData []byte) error {
	httpposturl := fmt.Sprintf("http://45.79.208.204:8080/api/collections/%s/records", collID)

	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return error	
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}

func GetData(collID string, pageNum string) (error, ResponseData) {
	requestURL := fmt.Sprintf("http://45.79.208.204:8080/api/collections/%s/records?page=%s", collID, pageNum)
	res, err := http.Get(requestURL)

	var jsonData ResponseData

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return err, jsonData
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err, jsonData
	}
	
	err = json.Unmarshal([]byte(resBody), &jsonData)

	if err != nil {
		fmt.Println("HERE")
		panic(err)
	}

	return nil, jsonData
}

func GetAllData(collID string) (error, []Item ) {
	var items []Item

	returnLength := 30
	i := 0
	for returnLength == 30 {
		err, data := GetData(collID, strconv.Itoa(i))
		if err != nil {
			return err, items
		}

		items = append(items, data.Items...)

		fmt.Println(len(data.Items))
		fmt.Println(i)

		i = i + 1
		returnLength = len(data.Items)
	}

	return nil, items
}

func UpdateData(collName string, recordID string, jsonData []byte) error {
	httpposturl := fmt.Sprintf("http://45.79.208.204:8080/api/collections/%s/records/%s", collName, recordID)

	request, error := http.NewRequest("PATCH", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return error	
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func DeleteData(collName string, recordID string) error {
	deleteurl := fmt.Sprintf("http://45.79.208.204:8080/api/collections/%s/records/%s", collName, recordID)

    todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
    jsonReq, err := json.Marshal(todo)
    req, err := http.NewRequest(http.MethodDelete, deleteurl, bytes.NewBuffer(jsonReq))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    defer resp.Body.Close()


	return nil

}

func DeleteCollection(collName string) error {
	
	items_length := 1

	for items_length != 0 {
		err, data := GetData(collName, "1")
		if err != nil {
			return err
		}

		items_length = len(data.Items)

		for _, v := range data.Items {
			err = DeleteData(collName, v.ID)

			if err != nil {
				return err
			}
		}
	}
	
	return nil
}
