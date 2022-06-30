resource "aws_lambda_function" "forwarder" {
  function_name = "smalldomains.forwarder"
  description   = "Responsible for making HTTP Redirects for SmallDomain users"
  role          = aws_iam_role.forwarder-lambda.arn
  handler       = "index.test"

  filename         = "deploy_artifact.zip"
  source_code_hash = filebase64sha256("deploy_artifact.zip")
  package_type     = "Zip"
  runtime          = "go1.x"

  memory_size                    = var.appconfig-memory_size
  publish                        = var.appconfig-publish_new_lambda_version
  reserved_concurrent_executions = var.appconfig-reserved_concurrent_executions

  environment {
    variables = {
      smallDomainsGetterUrl = var.appconfig-smallDomainsGetterUrl
    }
  }
}