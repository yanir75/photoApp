module "s3" {
  source = "./modules/s3"
  bucket_prefix = var.bucket_prefix
}

module "vpc" {
  source = "./modules/vpc"
}