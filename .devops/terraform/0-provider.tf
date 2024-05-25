terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.51"
    }
  }
}

provider "aws" {
  region = "us-east-2"
}

variable "cluster_name" {
  default = "meli_cluster"
}

variable "cluster_version" {
  default = "1.30"
}

variable "account_id" {
  default = "339713144439"
}
