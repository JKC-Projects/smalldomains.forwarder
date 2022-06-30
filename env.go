package main

import (
	"os"
)

type EnvVars struct {
	SmallDomainsGetterUrl string
}

func getEnvVars() EnvVars {
	return EnvVars{
		SmallDomainsGetterUrl: os.Getenv("smallDomainsGetterUrl"),
	}
}