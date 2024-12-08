output "bucket_id" {
  value = module.s3_bucket.bucket_id
}

output "bucket_arn" {
  value = module.s3_bucket.bucket_arn
}

output "bucket_name" {
  value = module.s3_bucket.bucket_name
}

output "versioning_enabled" {
  value = module.s3_bucket.aws_s3_bucket_versioning.this.enabled
}
