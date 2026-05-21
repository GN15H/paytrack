resource "aws_ecr_repository" "paytrack" {
  name                 = "paytrack"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = { Name = "paytrack-ecr" }
}

output "ecr_repository_url" {
  value = aws_ecr_repository.paytrack.repository_url
}
