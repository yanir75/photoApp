output "instance_id" {
  description = "Instance ID of the deployed instance"
  value = aws_instance.instance.id
}

output "instance_public_ip" {
  description = "Public IP of the instance"
  value = aws_instance.instance.public_ip
}