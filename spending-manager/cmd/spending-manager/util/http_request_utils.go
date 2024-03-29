package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func WriteBodyToObject(body io.Reader, objectToWrite any) {
	responseBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(responseBody, &objectToWrite)
}
