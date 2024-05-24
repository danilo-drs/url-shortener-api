resource "aws_db_instance" "meli_db" {
  allocated_storage        = 20
  engine                   = "postgres"
  engine_version           = "16.1"
  identifier               = "meli"
  instance_class           = "db.t3.micro"
  storage_encrypted        = false
  publicly_accessible      = true
  delete_automated_backups = true
  skip_final_snapshot      = true
  db_name                  = "meli"
  username                 = "postgres"
  password                 = "jOFjqRTpKZ0QptqDVVln"
  apply_immediately        = true
  multi_az                 = false
}
