resource "aws_vpc" "meli_main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "meli_main"
  }
}
