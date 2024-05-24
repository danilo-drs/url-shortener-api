resource "aws_apigatewayv2_api" "meli_api" {
  name          = "meli_api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "dev" {
  api_id = aws_apigatewayv2_api.meli_api.id

  name        = "dev"
  auto_deploy = true
}
