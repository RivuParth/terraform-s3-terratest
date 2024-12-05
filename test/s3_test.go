package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Path to the Terraform code
		TerraformDir: "../s3", 
		// Variables for the Terraform module
		Vars: map[string]interface{}{
			"bucket_name":      "your-terratest-004",  
			"enable_versioning": true,               
			"tags": map[string]string{
				"Environment": "Test",
			},
		},
		// Ensure Terraform doesn't ask for user input during apply
		NoColor: true,
	}
	

	// Clean up resources after the test
	defer terraform.Destroy(t, terraformOptions)

	// Run Terraform init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Get the bucket ID and ARN from the Terraform output
	bucketId := terraform.Output(t, terraformOptions, "bucket_id")
	assert.NotEmpty(t, bucketId, "Bucket ID should not be empty")

	bucketArn := terraform.Output(t, terraformOptions, "bucket_arn")
	assert.Contains(t, bucketArn, "arn:aws:s3:::"+bucketId, "Bucket ARN should contain the correct bucket ID")

	// Verify the bucket versioning status
	// versioningStatus := terraform.Output(t, terraformOptions, "bucket_versioning_status")
	// assert.Equal(t, "Enabled", versioningStatus, "Versioning should be enabled")

	// Optionally, use AWS SDK to verify the S3 bucket's location
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create an S3 service client
	s3Client := s3.New(sess)

	// Get bucket location
	_, err = s3Client.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(bucketId),
	})
	assert.NoError(t, err, "Error checking S3 bucket location")
}
