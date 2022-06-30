resource "aws_iam_role" "forwarder-lambda" {
  name = "Role_SmallDomains_Forwarder"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.forwarder-lambda.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

data "aws_iam_policy" "lambda_logging" {
  name = "AWSLambdaBasicExecutionRole"
}