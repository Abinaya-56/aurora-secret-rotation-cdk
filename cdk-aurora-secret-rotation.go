package main

import (
    "github.com/aws/aws-cdk-go/awscdk/v2"
    "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
    "github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
    "github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)

type CdkAuroraSecretRotationStackProps struct {
        awscdk.StackProps
}

func NewCdkAuroraSecretRotationStack(scope constructs.Construct, id string, props *CdkAuroraSecretRotationStackProps) awscdk.Stack {
        var sprops awscdk.StackProps
        if props != nil {
                sprops = props.StackProps
        }
        stack := awscdk.NewStack(scope, &id, &sprops)

        // The code that defines your stack goes here
	
	vpc := awsec2.NewVpc(stack, jsii.String("MyVpc"), &awsec2.VpcProps{
    	MaxAzs: jsii.Number(2),
	})

	vpc.AddInterfaceEndpoint(jsii.String("SecretsEndpoint"), &awsec2.InterfaceVpcEndpointOptions{
    	Service: awsec2.InterfaceVpcEndpointAwsService_SECRETS_MANAGER(),
	})

	dbSecret := awssecretsmanager.NewSecret(stack, jsii.String("DBSecret"), &awssecretsmanager.SecretProps{
    	GenerateSecretString: &awssecretsmanager.SecretStringGenerator{
        SecretStringTemplate: jsii.String(`{"username":"postgres"}`),
        GenerateStringKey: jsii.String("password"),
        ExcludeCharacters: jsii.String(`/@" `),
    	},
	})

	cluster := awsrds.NewDatabaseCluster(stack, jsii.String("AuroraCluster"), &awsrds.DatabaseClusterProps{
    	Engine: awsrds.DatabaseClusterEngine_AuroraPostgres(&awsrds.AuroraPostgresClusterEngineProps{
        Version: awsrds.AuroraPostgresEngineVersion_VER_14_6(),
    	}),
    	Credentials: awsrds.Credentials_FromSecret(dbSecret, nil),
    	Writer: awsrds.ClusterInstance_ServerlessV2(jsii.String("writer"), &awsrds.ServerlessV2ClusterInstanceProps{}),
    	Vpc: vpc,
	})

	lambdaFn := awslambda.NewFunction(stack, jsii.String("DBLambda"), &awslambda.FunctionProps{
    	Runtime: awslambda.Runtime_PYTHON_3_9(),
    	Handler: jsii.String("lambda.handler"),
    	Code: awslambda.Code_FromAsset(jsii.String("lambda"), nil),
    	Vpc: vpc,
    	Environment: &map[string]*string{
        "DB_HOST": cluster.ClusterEndpoint().Hostname(),
        "DB_PORT": jsii.String("5432"),
    	},
	})

	cluster.Connections().AllowDefaultPortFrom(lambdaFn, jsii.String("Allow Lambda access"))

	
	dbSecret.GrantRead(lambdaFn, nil)
	
	_ = vpc
	_ = dbSecret
	_ = cluster
        // example resource
        // queue := awssqs.NewQueue(stack, jsii.String("CdkAuroraSecretRotationQueue"), &awssqs.QueueProps{
        //      VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
        // })

        return stack
}

func main() {
        defer jsii.Close()

        app := awscdk.NewApp(nil)

        NewCdkAuroraSecretRotationStack(app, "CdkAuroraSecretRotationStack", &CdkAuroraSecretRotationStackProps{
                awscdk.StackProps{
                        Env: env(),
                },
        })

        app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
        // If unspecified, this stack will be "environment-agnostic".
        // Account/Region-dependent features and context lookups will not work, but a
        // single synthesized template can be deployed anywhere.
        //---------------------------------------------------------------------------
        return nil

        // Uncomment if you know exactly what account and region you want to deploy
        // the stack to. This is the recommendation for production stacks.
        //---------------------------------------------------------------------------
        // return &awscdk.Environment{
        //  Account: jsii.String("123456789012"),
        //  Region:  jsii.String("us-east-1"),
        // }

        // Uncomment to specialize this stack for the AWS Account and Region that are
        // implied by the current CLI configuration. This is recommended for dev
        // stacks.
        //---------------------------------------------------------------------------
        // return &awscdk.Environment{
        //  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
        //  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
        // }
}
