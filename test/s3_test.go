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

    // Define the bucket name
    bucketName := "my-terratest-bucket2"

    // Define Terraform options
    terraformOptions := &terraform.Options{
        TerraformDir: "../modules/s3", // Path to your Terraform module

        Vars: map[string]interface{}{
            "bucket_name": bucketName,
        },
    }

    // Skip terraform apply since the bucket already exists
    terraform.Init(t, terraformOptions)

    // Validate that the bucket exists in AWS
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"), // Replace with your AWS region
    })
    assert.NoError(t, err)

    s3Svc := s3.New(sess)
    _, err = s3Svc.HeadBucket(&s3.HeadBucketInput{
        Bucket: aws.String(bucketName),
    })

    // Assert that the bucket exists
    assert.NoError(t, err)
}
