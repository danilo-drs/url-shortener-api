resource "aws_security_group" "vpc_link" {
  name   = "vpc-link"
  vpc_id = aws_vpc.meli_main.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_apigatewayv2_vpc_link" "eks" {
  name               = "eks"
  security_group_ids = [aws_security_group.vpc_link.id]
  subnet_ids = [
    aws_subnet.private-us-east-2a.id,
    aws_subnet.private-us-east-2b.id
  ]
}

resource "aws_apigatewayv2_integration" "eks" {
  api_id = aws_apigatewayv2_api.meli_api.id

  integration_uri    = "arn:aws:elasticloadbalancing:us-east-2:339713144439:listener/net/a14ee6952610944e1a322aa63a6032cd/9fb898d5188c9c5e/05b30dc2c0963bd7"
  integration_type   = "HTTP_PROXY"
  integration_method = "ANY"
  connection_type    = "VPC_LINK"
  connection_id      = aws_apigatewayv2_vpc_link.eks.id
}

resource "aws_apigatewayv2_route" "get_echo" {
  api_id = aws_apigatewayv2_api.meli_api.id

  route_key = "GET /healthcheck"
  target    = "integrations/${aws_apigatewayv2_integration.eks.id}"
}

output "meli_base_url" {
  value = "${aws_apigatewayv2_stage.dev.invoke_url}/healthcheck"
}
