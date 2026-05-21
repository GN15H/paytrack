resource "aws_db_subnet_group" "main" {
  name       = "paytrack-db-subnet-group"
  subnet_ids = [aws_subnet.public.id, aws_subnet.private.id]

  tags = { Name = "paytrack-db-subnet-group" }
}

resource "aws_db_instance" "postgres" {
  identifier        = "paytrack-db"
  engine            = "postgres"
  engine_version    = "16"
  instance_class    = "db.t3.micro"
  allocated_storage = 20

  db_name  = "paytrack"
  username = "paytrack"
  password = var.db_password

  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.rds.id]

  skip_final_snapshot = true
  publicly_accessible = false

  tags = { Name = "paytrack-db" }
}

output "rds_endpoint" {
  value = aws_db_instance.postgres.endpoint
}
