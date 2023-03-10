import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsqvpc from "@pulumi/aws-quickstart-vpc";


const vpc = new awsqvpc.Vpc("compdb", {
    cidrBlock: "10.0.0.0/16",
    availabilityZoneConfig: [
        {
            availabilityZone: "eu-west-1a",
            publicSubnetCidr: "10.0.128.0/20",
            privateSubnetACidr: "10.0.32.0/19",
        }, 
        {
            availabilityZone: "eu-west-1b",
            privateSubnetACidr: "10.0.64.0/19",
        }
    ]
});

// Create an AWS resource (S3 Bucket)
const bucket = new aws.s3.Bucket("my-bucket",
    undefined,
    {
        dependsOn: vpc
    }
);

// Export the name of the bucket
export const bucketName = bucket.id;
