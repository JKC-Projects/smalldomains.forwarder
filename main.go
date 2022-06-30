package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/JKC-Project/smalldomains.forwarder/smalldomains"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var envVars = getEnvVars()
var client = smalldomains.Client{
	SmallDomainsGetterUrl: getEnvVars().SmallDomainsGetterUrl,
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx context.Context, request events.ALBTargetGroupRequest) (resp events.ALBTargetGroupResponse, error error) {
	defer func() {
		if r := recover(); r != nil {
			resp = constructInternalServerError()
		}
	}()

	if request.HTTPMethod != "GET" {
		return constructMethodNotAllowedResponse(), nil
	}

	smallDomainAlias := extractSmallDomainAliasFromPath(request.Path)
	smallDomain, err := client.GetSmallDomain(smallDomainAlias)

	if err == nil {
		return constructRedirectResponse(smallDomain.LargeDomain), nil
	} else {
		return constructNotFoundResponse(smallDomainAlias), nil
	}
}

func extractSmallDomainAliasFromPath(path string) string {
	regex := regexp.MustCompile("([a-zA-Z0-9\\-_]+)$")
	return regex.FindString(path)
}

func constructRedirectResponse(url string) events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        302,
		StatusDescription: "URL Shortner: Redirecting to aliased location.",
		Headers: map[string]string{
			"Location": url,
		},
	}
}

func constructNotFoundResponse(desiredSmallDomain string) events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode: 404,
		Body:       fmt.Sprintf("404: No SmallDomains Found for %v", desiredSmallDomain),
	}
}

func constructMethodNotAllowedResponse() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        405,
		StatusDescription: "405: HTTP Method Not Allowed.",
		Headers: map[string]string{
			"Allow": "GET",
		},
	}
}

func constructInternalServerError() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        500,
		StatusDescription: "500: Internal Server Error.",
		Body:              "500: Internal Server Error.",
	}
}
