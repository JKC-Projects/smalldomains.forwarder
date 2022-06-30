locals {
  lambda_function_name = "smalldomains--forwarder"
}

resource "aws_lambda_function" "forwarder" {
  function_name = local.lambda_function_name
  description   = "Responsible for making HTTP Redirects for SmallDomain users"
  role          = aws_iam_role.forwarder-lambda.arn
  handler       = "LambdaHandler"

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

  depends_on = [
    aws_cloudwatch_log_group.forwarder_lambda
  ]
}

resource "aws_lb_target_group_attachment" "forwarder" {
  target_group_arn = data.aws_ssm_parameter.forwarder-target-group-arn.value
  target_id        = aws_lambda_function.forwarder.arn
  depends_on       = [aws_lambda_permission.with_lb]
}

resource "aws_lambda_permission" "with_lb" {
  statement_id  = "AllowExecutionFromlb"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.forwarder.arn
  principal     = "elasticloadbalancing.amazonaws.com"
  source_arn    = aws_lb_target_group.forwarder.arn
}