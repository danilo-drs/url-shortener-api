terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "us-east-2"
}

variable "cluster_name" {
  default = "meli"
}

variable "cluster_version" {
  default = "1.29"
}

variable "account_id" {
  default = "339713144439"
}
