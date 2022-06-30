resource "aws_cloudwatch_log_group" "forwarder_lambda" {
  name = "/aws/lambda/${local.lambda_function_name}"
  tags {
    logging_for = "/lambda/${local.lambda_function_name}"
  }
}