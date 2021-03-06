package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

func GetSubscribers() (Item, error) {

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("id", os.Getenv("YOUTUBE_CHANNEL_ID"))
	q.Add("part", "statistics")

	req.URL.RawQuery = q.Encode()

	// create new http client to send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	// unmarshal bytes from io.readall call into the response object
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Item{}, err
	}

	// return the first item in the array
	return response.Items[0], nil
}
