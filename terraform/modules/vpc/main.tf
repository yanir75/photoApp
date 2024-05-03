resource "aws_vpc" "main" {
  cidr_block = var.cidr
  enable_dns_hostnames = true
  enable_dns_support = true
      tags = {
    Name = var.vpc_name
  }
}

resource "aws_subnet" "main" {
    count = var.number_of_subnets
    vpc_id = aws_vpc.main.id
  cidr_block = cidrsubnet(var.cidr,8,count.index)
    tags = {
    Name = "${var.vpc_name} - ${count.index}"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.vpc_name} GW"
  }
}



resource "aws_route_table" "rt" {
  vpc_id = aws_vpc.main.id

  # since this is exactly the route AWS will create, the route will be adopted
  route {
    cidr_block = var.cidr
    gateway_id = "local"
  }
    tags = {
    Name = "${var.vpc_name} RT"
  }
}

resource "aws_route" "outbound" {
    route_table_id            = aws_route_table.rt.id
  destination_cidr_block    = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.gw.id
}


resource "aws_route_table_association" "rt" {
    count = var.number_of_subnets
  subnet_id      = aws_subnet.main[count.index].id
  route_table_id = aws_route_table.rt.id
}