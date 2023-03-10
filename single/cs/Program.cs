using Pulumi;
using Pulumi.Aws.S3;
using Pulumi.Aws.Rds;
using System.Collections.Generic;

return await Deployment.RunAsync(() =>
{
   var compDb = new SingleLanguage.ComponentDatabase("compdb", new ComponentResourceOptions {});

    // Create an AWS resource (S3 Bucket)
    var bucket = new Bucket("my-bucket", null, new CustomResourceOptions { DependsOn = { compDb } });

    // Export the name of the bucket
    return new Dictionary<string, object?> { ["bucketName"] = bucket.Id };
});

namespace SingleLanguage
{
    class ComponentDatabase : Pulumi.ComponentResource
    {
        public ComponentDatabase(string name, ComponentResourceOptions opts)
            : base("single:index:ComponentDatabase", name, opts)
        {
            // initialization logic.
            var @default = new Instance(
                name,
                new()
                {
                    AllocatedStorage = 10,
                    Engine = "mysql",
                    EngineVersion = "5.7",
                    InstanceClass = "db.t3.micro",
                    ParameterGroupName = "default.mysql5.7",
                    Password = "foobarbaz",
                    SkipFinalSnapshot = true,
                    Username = "foo",
                },
                new CustomResourceOptions { 
                  Parent = this 
               }
            );

            // Signal to the UI that this resource has completed construction.
            this.RegisterOutputs();
        }
    }
}
