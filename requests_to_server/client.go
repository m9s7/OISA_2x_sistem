package requests_to_server

import (
	"encoding/json"
	"fmt"
	"github.com/rogpeppe/retry"
	"io"
	"net/http"
	"time"
)

var Mozzart = httpClient()
var Maxbet = httpClient()
var Merkurxtip = httpClient()
var Soccerbet = httpClient()

func httpClient() *http.Client {

	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

var RetryStrategy = retry.Strategy{
	Delay:       100 * time.Millisecond,
	MaxDelay:    4 * time.Second,
	MaxDuration: 20 * time.Second,
	Factor:      2,
}

func GetJson(client *http.Client, req *http.Request, target interface{}) error {

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	return json.NewDecoder(res.Body).Decode(target)
}
