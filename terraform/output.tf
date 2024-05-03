output "vpc" {
  description = "VPC outputs"
  value = module.vpc
}

output "s3" {
  description = "S3 outputs"
  value = module.s3
}