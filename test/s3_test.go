package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketModule(t *testing.T) {
	t.Parallel()

	// Define Terraform options with initial variables
	terraformOptions := &terraform.Options{
		TerraformDir: "../s3", // Path to the Terraform code
		// Vars: map[string]interface{}{
		// 	"bucket_name":        "partha-terratest-505",
		// 	"enable_versioning":  true,
		// 	"tags": map[string]string{
		// 		"Environment": "Test",
		// 	},
		// },
		NoColor:      true,                // Disable color in Terraform commands
		PlanFilePath: "./terraform-plan",  // Path to save the Terraform plan file
	}

	// Ensure resources are destroyed at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Initialize and plan Terraform
	terraform.InitAndPlan(t, terraformOptions)

	// Apply the saved plan file
	terraform.Apply(t, &terraform.Options{
		TerraformDir:   "../s3",
		PlanFilePath:   "./terraform-plan", // Apply the saved plan
	})

	// Run validations
	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	assert.Equal(t, "partha-terratest-505", bucketName)

	versioningEnabled := terraform.Output(t, terraformOptions, "versioning_enabled")
	assert.Equal(t, "true", versioningEnabled)

	tags := terraform.OutputMap(t, terraformOptions, "tags")
	assert.Equal(t, "Test", tags["Environment"])
}
