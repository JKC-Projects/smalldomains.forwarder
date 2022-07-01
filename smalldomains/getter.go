package smalldomains

import (
	"encoding/json"
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

type Client struct {
	SmallDomainsGetterUrl string
	Log logrus.Entry
}

type SmallDomain struct {
	SmallDomain string `json:"smallDomain"`
	LargeDomain string `json:"largeDomain"`
	CreatedAt   uint64 `json:"createdAt"`
	ExpiringAt  uint64 `json:"expiringAt"`
}

type SmallDomainRetrievalError struct{}

func (e SmallDomainRetrievalError) Error() string {
	return "Error retrieving SmallDomain"
}

func (this Client) GetSmallDomain(smallDomain string) (SmallDomain, error) {
	this.Log.Infof("Getting SmallDomain with alias: %v", smallDomain)
	resp, err := http.Get(this.SmallDomainsGetterUrl + "/" + smallDomain)

	if err != nil {
		this.Log.Errorf("Error when retrieving SmallDomain (%v): %v", smallDomain, err)
		return SmallDomain{}, err
	}

	defer resp.Body.Close()

	if !isSuccessStatusCode(resp.StatusCode) {
		this.Log.Errorf("Received error HTTP code (%v) when retrieving SmallDomain (%v)", resp.StatusCode, smallDomain)
		return SmallDomain{}, SmallDomainRetrievalError{}
	}

	this.Log.Infof("Successfully retrieved SmallDomain (%v)", smallDomain)
	var toReturn SmallDomain
	json.NewDecoder(resp.Body).Decode(&toReturn)
	return toReturn, nil
}

func isSuccessStatusCode(statusCode int) bool {
	return 200 <= statusCode && statusCode < 300
}
