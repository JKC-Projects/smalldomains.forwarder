package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/JKC-Project/smalldomains.forwarder/smalldomains"
	"github.com/sirupsen/logrus"
)

var envVars = getEnvVars()
var client, log = initialiseDependencies()

func main() {
	http.HandleFunc("/actuator/health", withPanicRecoveryMiddleware(handleHealthCheckHttpRequest))
	http.HandleFunc("/", withPanicRecoveryMiddleware(handleSmallDomainsRedirectHttpRequest))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initialiseDependencies() (client smalldomains.Client, log *logrus.Logger) {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info("Logger initialised")

	client = smalldomains.Client{
		SmallDomainsGetterUrl: getEnvVars().SmallDomainsGetterUrl,
		Log:                   *log,
	}
	log.Info("smallDomains.Client initialised")

	return
}

func withPanicRecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(resposeWriter http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Internal Server Error: %+v", r)
				writeInternalServerErrorHttpResponse(resposeWriter)
			}
		}()

		next(resposeWriter, request)
	}
}

func handleSmallDomainsRedirectHttpRequest(responseWriter http.ResponseWriter, request *http.Request) {
	smallDomainAlias := extractSmallDomainAliasFromPath(request.URL.Path)
	smallDomain, err := client.GetSmallDomain(smallDomainAlias)

	if err == nil {
		log.Infof("Successfully found a SmallDomain: %+v", smallDomain)
		writeRedirectHttpResponse(responseWriter, smallDomain.LargeDomain)
	} else {
		log.Errorf("Could not find a SmallDomain with alias: %v", smallDomainAlias)
		writeNotFoundHttpResponse(responseWriter, smallDomainAlias)
	}
}

func handleHealthCheckHttpRequest(responseWriter http.ResponseWriter, _ *http.Request) {
	isAppHealthy, healthJsonSummary := isAppHealthy(client)

	responseWriter.Header().Add("Content-Type", "application/json")

	if isAppHealthy {
		responseWriter.WriteHeader(200)
	} else {
		responseWriter.WriteHeader(503)
	}

	responseWriter.Write([]byte(healthJsonSummary))
}

func extractSmallDomainAliasFromPath(path string) string {
	regex := regexp.MustCompile("([a-zA-Z0-9-_]+)$")
	return regex.FindString(path)
}

func writeRedirectHttpResponse(responseWriter http.ResponseWriter, urlToRedirectTo string) {
	responseWriter.Header().Add("Location", urlToRedirectTo)
	responseWriter.WriteHeader(302)
	responseWriter.Write([]byte{})
}

func writeNotFoundHttpResponse(responseWriter http.ResponseWriter, desiredSmallDomain string) {
	httpResponseBody := fmt.Sprintf(`{
    "message" : "404 No SmallDomains found for %v"
  }
  `, desiredSmallDomain)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(404)
	responseWriter.Write([]byte(httpResponseBody))
}

func writeInternalServerErrorHttpResponse(responseWriter http.ResponseWriter) {
	httpResponseBody := `{
    "message" : "500 Internal Server Error."
  }`

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(500)
	responseWriter.Write([]byte(httpResponseBody))
}
