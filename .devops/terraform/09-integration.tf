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
    aws_subnet.private-us-east-2b.id,
    aws_subnet.public-us-east-2a.id,
    aws_subnet.public-us-east-2b.id
  ]
}
resource "aws_lb" "meli_load_balancer" {
  name               = "meli-load-balancer"
  load_balancer_type = "application"
  subnets = [
    aws_subnet.public-us-east-2a.id,
    aws_subnet.public-us-east-2b.id
  ]
}

resource "aws_lb_target_group" "meli_target_group" {
  name                 = "meli-target-group"
  port                 = 80
  protocol             = "HTTP"
  vpc_id               = aws_vpc.meli_main.id
  target_type          = "instance"
  deregistration_delay = 1
  lifecycle {
    create_before_destroy = true
  }

  health_check {
    path                = "/healthcheck"
    protocol            = "HTTP"
    port                = 8099
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2

  }
}

resource "aws_lb_listener" "meli_listener" {
  load_balancer_arn = aws_lb.meli_load_balancer.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.meli_target_group.arn
  }
}

resource "aws_alb_listener_rule" "rule-1" {
  action {
    target_group_arn = "arn:aws:elasticloadbalancing:us-east-2:339713144439:targetgroup/meli-target-group/0199d723bf60a04b"
    type             = "forward"
  }
  condition {
    http_request_method {
      values = ["GET"]
    }
  }

  listener_arn = "arn:aws:elasticloadbalancing:us-east-2:339713144439:listener/app/meli-load-balancer/7f7f9c8c11885d89/7320ec9e64ace630"
  priority     = 100
}

resource "aws_apigatewayv2_integration" "eks" {
  api_id          = aws_apigatewayv2_api.meli_api.id
  integration_uri = aws_lb_listener.meli_listener.arn
  # integration_uri    = "arn:aws:elasticloadbalancing:us-east-2:339713144439:listener/app/meli-load-balancer/7f7f9c8c11885d89/7320ec9e64ace630"
  integration_type   = "HTTP_PROXY"
  integration_method = "ANY"
  connection_type    = "VPC_LINK"
  connection_id      = aws_apigatewayv2_vpc_link.eks.id
}

resource "aws_apigatewayv2_route" "get_healthcheck" {
  api_id = aws_apigatewayv2_api.meli_api.id

  route_key = "GET /healthcheck"
  target    = "integrations/${aws_apigatewayv2_integration.eks.id}"
}

output "meli_base_url" {
  value = "${aws_apigatewayv2_stage.dev.invoke_url}/healthcheck"
}

