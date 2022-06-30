package main

import (
	"context"
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

func HandleLambdaEvent(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	if request.HTTPMethod != "GET" {
		return constructMethodNotAllowedResponse(), nil
	}

	smallDomainAlias := extractSmallDomainAliasFromPath(request.Path)
	smallDomain, err := client.GetSmallDomain(smallDomainAlias)

	if err == nil {
		return constructRedirectResponse(smallDomain.LargeDomain), nil
	} else {
		return constructNotFoundResponse(), nil
	}
}

func extractSmallDomainAliasFromPath(path string) string {
	regex := regexp.MustCompile("(?<=\\/)[a-zA-Z0-9\\-_]+$")
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

func constructNotFoundResponse() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        404,
		StatusDescription: "404: No SmallDomains Found.",
	}
}

func constructMethodNotAllowedResponse() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        405,
		StatusDescription: "404: No SmallDomains Found.",
		Headers: map[string]string{
			"Allow": "GET",
		},
	}
}
