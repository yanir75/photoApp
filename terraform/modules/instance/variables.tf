variable "image_filters" {
  description = "Which image filters to use in order to search in aws AMI-Search"
  type = list(object({
    name  = string
    values = list(string)
    }))
    default = [ {
      name   = "name"
      values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
    },
    {
    name   = "virtualization-type"
    values = ["hvm"]
    } ]
}

variable "owners" {
  description = "The owner of the AMI"
  type = list(string)
    default = ["099720109477"]
}

variable "associate_public_ip_address" {
    description = "Whether to associate a public IP address with an instance in a VPC."
    type = bool
    default = true
}

variable "instace_type" {
  description = "The type of the instance"
  type = string
  default = "t3.micro"
}

variable "instance_name" {
  description = "The name of the instance"
  type = string
}

variable "delete_on_termination" {
  description = "Whether the volume should be destroyed on instance termination."
  type = bool
  default = true
}

variable "encrypted" {
  description = "Whether to enable volume encryption"
  type = bool
  default = true
}

variable "iops" {
    description = "Amount of provisioned IOPS. Only valid for volume_type of io1, io2 or gp3."
    type = number
    default = null
}

variable "throughput" {
    description = "Throughput to provision for a volume in mebibytes per second (MiB/s). This is only valid for volume_type of gp3."
    type = number
    default = null
}

variable "volume_size" {
    description = "Size of the volume in gibibytes (GiB)."
    type = number
    default = 20
}

variable "volume_type" {
    description = "Type of volume. Valid values include standard, gp2, gp3, io1, io2, sc1, or st1."
    type = string
    default = "gp3"
}
  
}
  
}
  
}