variable "aws_region" {
  default = "us-east-1"
}

variable "db_password" {
  sensitive = true
}

variable "jwt_secret" {
  sensitive = true
}

variable "ec2_key_name" {
  description = "Name of the EC2 key pair for SSH access"
}

variable "profile" {
  sensitive = true
}
