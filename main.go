package main

import (
	"github.com/JKC-Project/smalldomains.forwarder/smalldomains"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

var envVars = getEnvVars()
var client = smalldomains.Client{
	SmallDomainsGetterUrl: getEnvVars().SmallDomainsGetterUrl,
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx lambdacontext.LambdaContext, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	path := request.Path

	smallDomain, err := client.GetSmallDomain(path)

	if err == nil {
		return constructRedirectResponse(smallDomain.LargeDomain), nil
	} else {
		return constructFailResponse(), nil
	}
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

func constructFailResponse() events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        404,
		StatusDescription: "404: No SmallDomains Found.",
	}
}
