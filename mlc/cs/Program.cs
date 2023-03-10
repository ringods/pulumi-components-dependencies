using Pulumi;
using Pulumi.Aws.S3;
using Pulumi.Aws.Rds;
using System.Collections.Generic;
using Pulumi.AwsQuickStartVpc;
using Pulumi.AwsQuickStartVpc.Inputs;

return await Deployment.RunAsync(() =>
{
    var vpc = new Vpc("quickvpc", new VpcArgs{
        CidrBlock = "10.0.0.0/16",
        AvailabilityZoneConfig = {
            new AvailabilityZoneArgs{
                AvailabilityZone = "eu-west-1a",
                PublicSubnetCidr = "10.0.128.0/20",
                PrivateSubnetACidr = "10.0.32.0/19",
            }, 
            new AvailabilityZoneArgs{
                AvailabilityZone = "eu-west-1b",
                PrivateSubnetACidr = "10.0.64.0/19",
            }
        }
    });

    // Create an AWS resource (S3 Bucket)
    var bucket = new Bucket("my-bucket", null, new CustomResourceOptions { DependsOn = { vpc } });

    // Export the name of the bucket
    return new Dictionary<string, object?> { ["bucketName"] = bucket.Id };
});
