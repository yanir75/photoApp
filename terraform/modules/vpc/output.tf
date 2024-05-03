output "vpc_id" {
  description = "VPC id of module"
  value = aws_vpc.main.id
}

output "subnets_ids" {
  description = "Subnets ids of module"
  value = aws_subnet.main[*].id
}