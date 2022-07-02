variable "environment" {
  type = string

  validation {
    condition     = contains(["dev", "prod"], var.environment)
    error_message = "Valid values for var.environment are [\"dev\",\"prod\"]."
  }
}

variable "appconfig-smallDomainsGetterUrl" {
  type = string
  validation {
    condition     = length(regexall("/$", var.appconfig-smallDomainsGetterUrl)) == 0
    error_message = "Configuration cannot end in a '/'"
  }
}

variable "appconfig-publish_new_lambda_version" {
  type = bool
}

variable "appconfig-memory_size" {
  type = number
  validation {
    condition     = var.appconfig-memory_size >= 128
    error_message = "Memory size must be greater than 128 MB."
  }
}

variable "appconfig-reserved_concurrent_executions" {
  type = number
}