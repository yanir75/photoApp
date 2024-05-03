variable "cidr" {
    description = "Cidr of the vpc"
  type = string
  default = "10.0.0.0/16"
}

variable "number_of_subnets" {
  description = "How many subnets to create"
  type = number
  default = 1
}

variable "vpc_name" {
  description = "Name of the vpc"
  type = string
  default = "Main Vpc APP"
}