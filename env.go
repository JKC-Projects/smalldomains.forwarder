package main

import (
	"fmt"
	"os"
)

type EnvVars struct {
	SmallDomainsGetterUrl string
	Environment           string
}

func getEnvVars() EnvVars {
	return EnvVars{
		SmallDomainsGetterUrl: os.Getenv("smallDomainsGetterUrl"),
		Environment:           getCurrEnvironment(),
	}
}

func getCurrEnvironment() string {
	envVar := os.Getenv("environment")

	if envVar != "dev" && envVar != "prod" {
		panic(fmt.Sprintf("Environment Variable \"environment\" was set to an invalid value: %v", envVar))
	}

	return envVar
}
