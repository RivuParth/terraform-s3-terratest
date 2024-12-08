package test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	// Define Terraform options
	terraformOptions := &terraform.Options{
		// Path to the Terraform module
		TerraformDir: "../s3",
		// Variables to pass to Terraform
		Vars: map[string]interface{}{
			"bucket_name":      "partha-terratest-505",
			"enable_versioning": true,
			"tags": map[string]string{
				"Environment": "Test",
			},
		},
		// Disable color output for readability in logs
		NoColor: true,
	}

	// Ensure cleanup after test
	defer terraform.Destroy(t, terraformOptions)

	// Initialize and apply the Terraform configuration
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	bucketId := terraform.Output(t, terraformOptions, "bucket_id")
	assert.NotEmpty(t, bucketId, "Bucket ID should not be empty")

	bucketArn := terraform.Output(t, terraformOptions, "bucket_arn")
	assert.Contains(t, bucketArn, "arn:aws:s3:::"+bucketId, "Bucket ARN should contain the correct bucket ID")

	// Optionally, verify bucket location using AWS SDK
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	assert.NoError(t, err, "Failed to create AWS session")

	s3Client := s3.New(sess)

	// Check bucket location
	_, err = s3Client.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(bucketId),
	})
	assert.NoError(t, err, "Error checking S3 bucket location")
}
