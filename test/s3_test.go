package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	// Define Terraform options
	options := &terraform.Options{
		// Path to your Terraform code
		TerraformDir: "../modules/s3",
		NoColor:      true, // Disable colored output for cleaner logs
	}

	// Initialize and validate the Terraform configuration
	terraform.Init(t, options)
	terraform.Validate(t, options)

	// Use the expected values for validation
	expectedBucketName := "my-terratest-bucket-66666"
	expectedRegion := "us-east-1"

	// Fetch outputs from Terraform
	bucketName := terraform.Output(t, options, "bucket_name")
	assert.NotEmpty(t, bucketName, "Bucket name should not be empty")
	
	bucketRegion := terraform.Output(t, options, "bucket_region")
	assert.Equal(t, "us-east-1", bucketRegion, "Bucket region does not match expected value")
	
	bucketArn := terraform.Output(t, options, "bucket_arn")
	assert.Contains(t, bucketArn, bucketName, "Bucket ARN should contain the bucket name")
	
	bucketId := terraform.Output(t, options, "bucket_id")
	assert.NotEmpty(t, bucketId, "Bucket ID should not be empty")
	
}
