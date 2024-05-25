resource "aws_apigatewayv2_api" "meli_api" {
  name          = "meli_api"
  protocol_type = "HTTP"
}

resource "aws_cloudwatch_log_group" "api_log_group" {
  name = "/aws/apigateway/meli_api"
}
resource "aws_apigatewayv2_stage" "dev" {
  api_id = aws_apigatewayv2_api.meli_api.id

  name        = "dev"
  auto_deploy = true
  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_log_group.arn
    format = jsonencode({
      requestId = "$context.requestId",
      ip        = "$context.identity.sourceIp",
      request   = "$context.httpMethod $context.routeKey",
      status    = "$context.status",
      response  = "$context.responseLength"
    })
  }
}
