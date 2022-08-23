package main

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("smallDomainsGetterUrl", "https://api.dev.small.domains/smalldomains")
	os.Setenv("environment", "dev")
}

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

func TestSmallDomainsAllowsHyphens(t *testing.T) {
	const path = "test-this-small-domain"
	extractedSmallDomain := extractSmallDomainAliasFromPath(path)

	if extractedSmallDomain != path {
		t.Errorf("SmallDomain incorrectly disallows hyphens: %v", path)
	}
}

func TestSmallDomainsAllowsUnderscores(t *testing.T) {
	const path = "test_this_small_domain"
	extractedSmallDomain := extractSmallDomainAliasFromPath(path)

	if extractedSmallDomain != path {
		t.Errorf("SmallDomain incorrectly disallows underscores: %v", path)
	}
}
