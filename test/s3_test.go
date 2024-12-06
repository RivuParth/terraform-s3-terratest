package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/s3", // Path to your Terraform code
		NoColor:      true,
	}

	// Clean up resources after the test
	defer terraform.Destroy(t, terraformOptions)

	// Run Terraform init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Expected values
	expectedBucketName := "my-terratest-bucket-666666"
	expectedRegion := "us-east-1"

	// Validate bucket name
	actualBucketName := terraform.Output(t, terraformOptions, "bucket_name")
	assert.Equal(t, expectedBucketName, actualBucketName, "Bucket name does not match expected value")

	// Validate bucket region
	actualRegion := terraform.Output(t, terraformOptions, "bucket_region")
	assert.Equal(t, expectedRegion, actualRegion, "Bucket region does not match expected value")
}
