package smalldomains

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	SmallDomainsGetterUrl string
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
	resp, err := http.Get(this.SmallDomainsGetterUrl)

	if err != nil {
		return SmallDomain{}, err
	}
	if isSuccessStatusCode(resp.StatusCode) {
		return SmallDomain{}, SmallDomainRetrievalError{}
	}

	defer resp.Body.Close()
	var toReturn *SmallDomain
	json.NewDecoder(resp.Body).Decode(toReturn)
	return *toReturn, nil
}

func isSuccessStatusCode(statusCode int) bool {
	return 200 <= statusCode && statusCode < 300
}