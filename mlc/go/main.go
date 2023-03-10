package main

import (
	awsqvpc "github.com/pulumi/pulumi-aws-quickstart-vpc/sdk/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		publicSubnet1Cidr := "10.0.128.0/20"
		privateSubnet1ACidr := "10.0.32.0/19"

		publicSubnet2Cidr := "10.0.64.0/19"

		vpc, err := awsqvpc.NewVpc(ctx, "quickvpc", &awsqvpc.VpcArgs{
			CidrBlock: "10.0.0.0/16",
			AvailabilityZoneConfig: []awsqvpc.AvailabilityZoneArgs{
				awsqvpc.AvailabilityZoneArgs{
					AvailabilityZone:   "eu-west-1a",
					PublicSubnetCidr:   &publicSubnet1Cidr,
					PrivateSubnetACidr: &privateSubnet1ACidr,
				},
				awsqvpc.AvailabilityZoneArgs{
					AvailabilityZone:   "eu-west-1b",
					PrivateSubnetACidr: &publicSubnet2Cidr,
				},
			},
		})
		if err != nil {
			return err
		}

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil, pulumi.DependsOn([]pulumi.Resource{vpc}))
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
