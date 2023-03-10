import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";

class ComponentDatabase extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("single:index:ComponentDatabase", name, {}, opts);

        const db = new aws.rds.Instance(name,
            {
                allocatedStorage: 10,
                engine: "mysql",
                engineVersion: "5.7",
                instanceClass: "db.t3.micro",
                parameterGroupName: "default.mysql5.7",
                password: "foobarbaz",
                skipFinalSnapshot: true,
                username: "foo"
            },
            {
                parent: this
            }
        )

        this.registerOutputs();
    }
}

const compdb = new ComponentDatabase("compdb");

const compdb2 = new ComponentDatabase("compdb2",
    {
        dependsOn: compdb
    }
);

// Create an AWS resource (S3 Bucket)
const bucket = new aws.s3.Bucket("my-bucket",
    undefined,
    {
        dependsOn: compdb
    }
);

// Export the name of the bucket
export const bucketName = bucket.id;
