package connector

import (
	"fmt"
	"net/http"
	"spending-manager/cmd/spending-manager/config"
)

func GetCurrentUser(requestHeader http.Header) (*http.Response, error) {
	methodName := "/getUserByToken"
	requestURL := fmt.Sprintf("%s%s", config.Config.Connector.Auth, methodName)
	authRequest, _ := http.NewRequest("GET", requestURL, nil)
	authRequest.Header = requestHeader
	return http.DefaultClient.Do(authRequest)
}
