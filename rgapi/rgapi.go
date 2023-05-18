package rgapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	BaseURL string = "api.riotgames.com/lol"

	apiKey string = ""

	//ErrAPIKeyNotSet is the error returned when no global API key has been set
	ErrAPIKeyNotSet = errors.New("rgapi: API key has not been set. If you need a key visit https://developer.riotgames.com/")
	shortRateChan   rateChan
	longRateChan    rateChan
)

type rateChan struct {
	RateQueue   chan bool
	TriggerChan chan bool
}

// RiotError contains the http status code of the erro
type RiotError struct {
	StatusCode int
}

func SetAPIKey(key string) {
	apiKey = key
}

func isKeySet() bool {
	return apiKey != ""
}

func apiGet(region, requestURL string, target interface{}) error {
	if !isKeySet() {
		return ErrAPIKeyNotSet
	}

	checkRateLimiter(shortRateChan)
	checkRateLimiter(longRateChan)

	client := &http.Client{}
	completeURL := fmt.Sprintf("https://%s.%s%s", region, BaseURL, requestURL)
	fmt.Println(completeURL)
	req, err := http.NewRequest("GET", completeURL, nil)
	req.Header.Add("X-Riot-Token", apiKey)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	checkTimeTrigger(shortRateChan)
	checkTimeTrigger(longRateChan)
	if resp.StatusCode != http.StatusOK {
		return RiotError{StatusCode: resp.StatusCode}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return err
}

// SetShortRateLimit allows a custom rate limit to be set. For at the time of this writing the default
// for a development API key is 20 requests every 1 second
func SetShortRateLimit(numRequests int, perTime time.Duration) {
	shortRateChan = rateChan{
		RateQueue:   make(chan bool, numRequests),
		TriggerChan: make(chan bool),
	}
	go rateLimitHandler(shortRateChan, perTime)
}

func rateLimitHandler(RateChan rateChan, perTime time.Duration) {
	returnChan := make(chan bool)
	go timeTriggerWatcher(RateChan.TriggerChan, returnChan)
	for {
		<-returnChan
		<-time.After(perTime)
		go timeTriggerWatcher(RateChan.TriggerChan, returnChan)
		length := len(RateChan.RateQueue)
		for i := 0; i < length; i++ {
			<-RateChan.RateQueue
		}
	}
}

func timeTriggerWatcher(timeTrigger chan bool, returnChan chan bool) {
	timeTrigger <- true
	returnChan <- true
}

// SetLongRateLimit allows a custom rate limit to be set. For at the time of this writing the default
// for a development API key is 100 requests every 2 minutes
func SetLongRateLimit(numRequests int, perTime time.Duration) {
	longRateChan = rateChan{
		RateQueue:   make(chan bool, numRequests),
		TriggerChan: make(chan bool),
	}
	go rateLimitHandler(longRateChan, perTime)
}

func checkRateLimiter(RateChan rateChan) {
	if RateChan.RateQueue != nil && RateChan.TriggerChan != nil {
		RateChan.RateQueue <- true
	}
}

func checkTimeTrigger(RateChan rateChan) {
	if RateChan.RateQueue != nil && RateChan.TriggerChan != nil {
		select {
		case <-RateChan.TriggerChan:
		default:
		}
	}
}

// Error prints the error message for a RiotError
func (err RiotError) Error() string {
	return fmt.Sprintf("Error: HTTP Status %d", err.StatusCode)
}
