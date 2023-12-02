package main

import (
	"testing"
)

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
