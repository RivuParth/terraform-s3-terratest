
module "s3_bucket" {
  source           = "../modules/s3"
  bucket_name      = var.bucket_name
  enable_versioning = var.enable_versioning
  tags             = var.tags
  
}

