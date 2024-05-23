resource "aws_internet_gateway" "meli_igw" {
  vpc_id = aws_vpc.meli_main.id

  tags = {
    Name = "meli_igw"
  }
}
