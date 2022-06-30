package main

import (
	"testing"
)

func TestRedirectResponseConstruction(t *testing.T) {
	const redirectTo = "https://google.com"
	resp := constructRedirectResponse(redirectTo)

	if resp.Headers["Location"] != redirectTo {
		t.Errorf("The generated response does not redirect to %v: %v", redirectTo, resp)
	}
}

func TestExtractsLastPathComponent(t *testing.T) {
	const path = "/first/second"
	extractedSmallDomain := extractSmallDomainAliasFromPath(path)

	if extractedSmallDomain != "second" {
		t.Errorf("SmallDomain of second incorrectly extracted from path: %v", path)
	}
}
