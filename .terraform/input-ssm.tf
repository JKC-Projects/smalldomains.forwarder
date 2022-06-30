data "aws_ssm_parameter" "forwarder-target-group-arn" {
  name = "/elb/target-groups/smalldomains/forwarder"
}