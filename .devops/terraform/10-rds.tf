resource "aws_db_parameter_group" "pg_meli" {
  name   = "pgmeli"
  family = "postgres16"
  parameter {
    name  = "log_connections"
    value = "1"
  }
}
resource "aws_db_instance" "melidb" {
  identifier             = "melidb"
  instance_class         = "db.t3.micro"
  allocated_storage      = 5
  engine                 = "postgres"
  engine_version         = "16.1"
  db_name                = "meli"
  username               = "postgres"
  password               = "jOFjqRTpKZ0QptqDVVln"
  db_subnet_group_name   = aws_db_subnet_group.subnet_meli.name
  vpc_security_group_ids = [aws_security_group.vpc_link.id]
  parameter_group_name   = aws_db_parameter_group.pg_meli.name
  publicly_accessible    = true
  skip_final_snapshot    = true
}


output "melidb_host" {
  value = aws_db_instance.melidb.address
}

output "melidb_port" {
  value = aws_db_instance.melidb.port
}

output "melidb_name" {
  value = aws_db_instance.melidb.db_name
}

output "melidb_username" {
  value = aws_db_instance.melidb.username
}

output "melidb_password" {
  value     = aws_db_instance.melidb.password
  sensitive = true
}
