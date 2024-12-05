variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
}

variable "tags" {
  description = "Tags to assign to the bucket"
  type        = map(string)
  default     = {}
}

variable "force_destroy" {
  description = "Allow force deletion of bucket (including objects)"
  type        = bool
  default     = false
}

variable "enable_object_lock" {
  description = "Enable S3 Object Lock on the bucket"
  type        = bool
  default     = false
}


variable "acl" {
  description = "Access Control List (ACL) for the bucket"
  type        = string
  default     = "private"
}

variable "versioning_enabled" {
  description = "Enable versioning for the bucket"
  type        = bool
  default     = false
}

variable "versioning_mfa_delete" {
  description = "Enable MFA delete for versioning"
  type        = bool
  default     = false
}



variable "enable_logging" {
  description = "Enable server access logging for the bucket"
  type        = bool
  default     = false
}

variable "logging_target_bucket" {
  description = "Target bucket for access logs"
  type        = string
  default     = ""
}

variable "logging_target_prefix" {
  description = "Prefix for access log objects"
  type        = string
  default     = ""
}

variable "enable_policy" {
  description = "Enable bucket policy"
  type        = bool
  default     = false
}

variable "bucket_policy" {
  description = "JSON bucket policy"
  type        = string
  default     = ""
}

variable "block_public_acls" {
  description = "Block public ACLs"
  type        = bool
  default     = true
}

variable "block_public_policy" {
  description = "Block public bucket policies"
  type        = bool
  default     = true
}

variable "ignore_public_acls" {
  description = "Ignore public ACLs"
  type        = bool
  default     = true
}

variable "restrict_public_buckets" {
  description = "Restrict public bucket policies"
  type        = bool
  default     = true
}

variable "enable_versioning" {
  description = "Enable versioning for the bucket"
  type        = bool
  default     = false
}
