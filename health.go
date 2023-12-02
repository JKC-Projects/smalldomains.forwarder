package main

import (
	"encoding/json"

	"github.com/JKC-Project/smalldomains.forwarder/smalldomains"
)

type healthCheckPayload struct {
	IsSmallDomainsClientHealthy bool `json:"canAccessSmallDomainsGetter"`
}

func (this healthCheckPayload) areAllHealthChecksOk() bool {
	return this.IsSmallDomainsClientHealthy
}

func isAppHealthy(client smalldomains.Client) (isHealthy bool, healthJsonSummary string) {
	healthChecks := healthCheckPayload{
		IsSmallDomainsClientHealthy: client.IsHealthy(),
	}

	healthCheckResponseBodyBytes, parseError := json.MarshalIndent(healthChecks, "", "  ")

  isHealthy = healthChecks.areAllHealthChecksOk()

  if (parseError == nil) {
    healthJsonSummary = string(healthCheckResponseBodyBytes)
  } else {
    healthJsonSummary = `{
      "parseError" : "Error marshalling JSON response to health check. This does not affect the actual health check"
    }`
  }

  return
}
