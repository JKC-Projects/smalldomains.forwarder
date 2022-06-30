resource "aws_cloudwatch_log_group" "forwarder_lambda" {
  name              = "/aws/lambda/${var.lambda_function_name}"
  retention_in_days = 14
}