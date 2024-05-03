output "bucket_name" {
  description = "Name of the newly created bucket"
  value = aws_s3_bucket.s3.id
}