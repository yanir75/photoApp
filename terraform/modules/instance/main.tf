data "aws_ami" "ubuntu" {
  most_recent = true

dynamic "filter" {
  for_each = var.image_filters
  content {
    name = filter.value["name"]
    values = filter.value["values"]
  }
}

  owners = var.owners
}

resource "aws_instance" "instance" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instace_type
  associate_public_ip_address = var.associate_public_ip_address
  root_block_device {
    delete_on_termination = var.delete_on_termination
    encrypted = var.encrypted
    iops = var.iops
    throughput = var.throughput
    volume_size = var.volume_size
    volume_type = var.volume_type
    tags = {
        Name = "${var.instance_name} - EBS"
    }
  }

  tags = {
    Name = var.instance_name
  }
  
  lifecycle {
    # The AMI ID must refer to an existing AMI that has the tag "nomad-server".
    postcondition {
      condition     =  contains(["io1", "io2", "gp3"], var.volume_type) || var.iops == null
      error_message = "IOPs can't be defined with volume type ${var.volume_type} volume type must be one of io1, io2, gp3."
    }
  }
}
