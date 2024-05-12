module "s3" {
  source = "./modules/s3"
  bucket_prefix = var.bucket_prefix
}

module "vpc" {
  source = "./modules/vpc"
}

module "instance" {
  source = "./modules/instance"
  subnet_id = module.vpc.subnets_ids[0]
  instance_name = var.instance_name
}