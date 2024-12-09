package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	// Set up Terraform options
	terraformOptions := &terraform.Options{
		TerraformDir: "../s3",  // Path to your Terraform module
		NoColor:      true,     // Disable colored output
		PlanFilePath: "./terraform-plan", // Path to save the plan file
	}

	// Ensure resources are destroyed at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Initialize and create a Terraform plan
	terraform.InitAndPlan(t, terraformOptions)

	// Apply the saved plan file
	terraform.Apply(t, &terraform.Options{
		TerraformDir: "../s3",
		PlanFilePath: "./terraform-plan",
	})

	// Run validations for outputs
	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	assert.Equal(t, "parthaa-terratest-000", bucketName)

	bucketArn := terraform.Output(t, terraformOptions, "bucket_arn")
	assert.Contains(t, bucketArn, "arn:aws:s3:::parthaa-terratest-000")

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	assert.Equal(t, "parthaa-terratest-000", bucketID) // Adjust expected value as needed
}
