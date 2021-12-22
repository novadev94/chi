package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// Returns GetRequestHttpError so callers have finer tune handle of error
// Separated from AppError because they serve different purpose(?)
func MakeGetRequest(url string, hearders map[string]string, query map[string]string, timeout time.Duration, response interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for k, v := range hearders {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := res.Body
	defer body.Close()

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return GetRequestHttpError{
			StatusCode: res.StatusCode,
			Body:       bytes,
		}
	}
	return json.Unmarshal(bytes, response)
}

func MakePostRequest(url string, headers map[string]string, query map[string]string, timeout time.Duration, response interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := res.Body
	defer body.Close()

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return GetRequestHttpError{
			StatusCode: res.StatusCode,
			Body:       bytes,
		}
	}
	return json.Unmarshal(bytes, response)
}

func fromWei(amount *big.Int) float64 {
	const decimals = 18

	if amount == nil {
		return 0
	}

	floatAmount := new(big.Float).SetInt(amount)
	power := new(big.Float).SetInt(new(big.Int).Exp(
		big.NewInt(10), big.NewInt(decimals), nil,
	))
	res := new(big.Float).Quo(floatAmount, power)
	result, _ := res.Float64()
	return result
}

func StringToBig(str string) (*big.Int, bool) {
	if str == "" {
		str = "0"
	}
	n := new(big.Int)
	return n.SetString(str, 10)
}

func GetMinNumber(a, b int64) string {
	if a < b {
		return strconv.FormatInt(a, 10)
	}
	return strconv.FormatInt(b, 10)
}

func PrettyPrint(obj interface{}) string {
	json, _ := json.MarshalIndent(obj, "", "\t")
	return string(json)
}

type GetRequestHttpError struct {
	StatusCode int
	Body       []byte
}

func (a GetRequestHttpError) Error() string {
	return fmt.Sprintf("status: %d, body: %s", a.StatusCode, string(a.Body))
}
