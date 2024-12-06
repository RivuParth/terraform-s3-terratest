output "bucket_name" {
  value = aws_s3_bucket.this.bucket
}

output "bucket_region" {
  value = aws_s3_bucket.this.region
}

output "bucket_arn" {
  value = aws_s3_bucket.this.arn
}

output "bucket_id" {
  value = aws_s3_bucket.this.id
}
