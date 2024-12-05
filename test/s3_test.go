package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	// Define Terraform options
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/s3", // Update path to your Terraform module
		Vars: map[string]interface{}{
			"bucket_name": "my-terratest-bucket66", // Match the bucket name used in your apply
		},
	}

	// Clean up resources after the test
	defer terraform.Destroy(t, terraformOptions)

	// Run Terraform init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Retrieve Terraform outputs
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	bucketARN := terraform.Output(t, terraformOptions, "bucket_arn")

	// Validate the bucket exists in AWS
	awsRegion := "us-east-1" // Replace with your AWS region
	assert.True(t, aws.S3BucketExists(t, awsRegion, bucketID))

	// Validate the bucket ARN matches the expected format
	expectedARN := "arn:aws:s3:::" + bucketID
	assert.Equal(t, expectedARN, bucketARN)
}
