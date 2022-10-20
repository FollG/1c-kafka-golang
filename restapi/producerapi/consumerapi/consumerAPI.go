package restapi

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const (
	HOST               = "http://localhost:8082"
	CONSUMER           = "testConsumer"
	GROUP              = "testGroup"
	NEW_CONSUMER       = "%s/consumers/%s"
	SUBSCRIBE_CONSUMER = "%s/consumers/%s/instances/%s/subscription"
	FETCH_CONSUMER     = "%s/consumers/%s/instances/%s/records"
	DELETE_CONSUMER    = "%s/consumers/%s/instances/%s"
	CONTENT_TYPE       = "application/vnd.kafka.json.v2+json"
)

func DoHelper(client *http.Client, url string, body []byte) error {
	bufferBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, url, bufferBody)
	if err != nil {
		return err
	}
	fmt.Printf("-->Call %s\n", req.URL)
	fmt.Printf("-->Body %s\n", string(body))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyResp := bufio.NewScanner(resp.Body)
	fmt.Printf("<--Response %s\n", resp.Status)
	for bodyResp.Scan() {
		fmt.Printf("<--Body %s\n", bodyResp.Text())
	}
	return nil
}

func main() {
	client := http.Client{}
	// New consumer
	url := fmt.Sprintf(NEW_CONSUMER, HOST, GROUP)
	body := fmt.Sprintf(`{"name":"%s", "format": "json"}`, CONSUMER)
	err := DoHelper(&client, url, []byte(body))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
}
