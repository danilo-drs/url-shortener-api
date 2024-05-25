resource "aws_eip" "meli_nat" {
  domain = "vpc"

  tags = {
    Name = "meli_nat"
  }
}

resource "aws_nat_gateway" "meli_nat" {
  allocation_id = aws_eip.meli_nat.id
  subnet_id     = aws_subnet.public-us-east-2a.id

  tags = {
    Name = "meli_nat"
  }

  depends_on = [aws_internet_gateway.meli_igw]
}
