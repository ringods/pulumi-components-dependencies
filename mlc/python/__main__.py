"""An AWS Python Pulumi program"""

import pulumi
from pulumi_aws import s3, rds
import pulumi_aws_quickstart_vpc

vpc = pulumi_aws_quickstart_vpc.Vpc(
    "vpc",
    cidr_block = "10.0.0.0/16",
    availability_zone_config = [
        pulumi_aws_quickstart_vpc.AvailabilityZoneArgs(
            availability_zone = "eu-west-1a",
            public_subnet_cidr = "10.0.128.0/20",
            private_subnet_a_cidr = "10.0.32.0/19",
        ), 
        pulumi_aws_quickstart_vpc.AvailabilityZoneArgs(
            availability_zone = "eu-west-1b",
            private_subnet_a_cidr = "10.0.64.0/19",
        )
    ]
)

# Create an AWS resource (S3 Bucket)
bucket = s3.Bucket('my-bucket',opts=pulumi.ResourceOptions(depends_on=[vpc]))

# Export the name of the bucket
pulumi.export('bucket_name', bucket.id)
