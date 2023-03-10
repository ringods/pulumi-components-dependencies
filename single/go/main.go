package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComponentDatabase struct {
	pulumi.ResourceState
}

func NewComponentDatabase(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*ComponentDatabase, error) {
	myComponentDatabase := &ComponentDatabase{}
	err := ctx.RegisterComponentResource("single:index:ComponentDatabase", name, myComponentDatabase, opts...)
	if err != nil {
		return nil, err
	}
	_, err = rds.NewInstance(ctx, name, &rds.InstanceArgs{
		AllocatedStorage:   pulumi.Int(10),
		DbName:             pulumi.String("mydb"),
		Engine:             pulumi.String("mysql"),
		EngineVersion:      pulumi.String("5.7"),
		InstanceClass:      pulumi.String("db.t3.micro"),
		ParameterGroupName: pulumi.String("default.mysql5.7"),
		Password:           pulumi.String("foobarbaz"),
		SkipFinalSnapshot:  pulumi.Bool(true),
		Username:           pulumi.String("foo"),
	}, pulumi.Parent(myComponentDatabase))
	if err != nil {
		return nil, err
	}

	ctx.RegisterResourceOutputs(myComponentDatabase, pulumi.Map{})

	return myComponentDatabase, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		compDb, err := NewComponentDatabase(ctx, "compdb")
		if err != nil {
			return err
		}

		_, err = NewComponentDatabase(ctx, "compdb2", nil, pulumi.DependsOn([]pulumi.Resource{compDb}))
		if err != nil {
			return err
		}

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil, pulumi.DependsOn([]pulumi.Resource{compDb}))
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
