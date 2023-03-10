"""An AWS Python Pulumi program"""

import pulumi
from pulumi_aws import s3, rds

class ComponentDatabase(pulumi.ComponentResource):
    def __init__(self, name, opts = None):
        super().__init__('single:index:ComponentDatabase', name, None, opts)
        db = rds.Instance(
            name,
            allocated_storage=10,
            db_name="mydb",
            engine="mysql",
            engine_version="5.7",
            instance_class="db.t3.micro",
            parameter_group_name="default.mysql5.7",
            password="foobarbaz",
            skip_final_snapshot=True,
            username="foo",
            opts=pulumi.ResourceOptions(parent=self)
        )
        self.register_outputs({})

compDb = ComponentDatabase("compdb")

# Create an AWS resource (S3 Bucket)
bucket = s3.Bucket('my-bucket',opts=pulumi.ResourceOptions(depends_on=[compDb]))

# Export the name of the bucket
pulumi.export('bucket_name', bucket.id)
