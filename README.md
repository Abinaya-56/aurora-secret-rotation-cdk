# AWS Aurora Secret Rotation with AWS CDK (Go)

## Overview

This project demonstrates how to provision a secure, production-grade AWS infrastructure using AWS CDK (Go) that includes:

* A custom VPC with public and private subnets
* AWS Secrets Manager for secure credential storage
* Amazon Aurora PostgreSQL Serverless v2 cluster
* Automated integration between database and secrets

The goal is to eliminate hardcoded credentials and implement secure, scalable database provisioning using Infrastructure as Code (IaC).

---

## Architecture

The infrastructure created includes:

* **VPC**

  * Public and private subnets across multiple AZs
  * NAT Gateways for outbound internet access from private subnets

* **Secrets Manager**

  * Stores database credentials securely
  * Automatically generates strong passwords

* **Aurora PostgreSQL (Serverless v2)**

  * Runs inside private subnets
  * Uses credentials directly from Secrets Manager
  * Auto-scales based on workload

---

## Tech Stack

* AWS CDK v2
* Go (Golang)
* AWS Services:

  * Amazon VPC
  * AWS Secrets Manager
  * Amazon RDS (Aurora PostgreSQL Serverless v2)
  * AWS IAM
  * AWS CloudFormation

---

## Project Structure

```
cdk-aurora-secret-rotation/
│
├── cdk-aurora-secret-rotation.go   # Main CDK stack definition
├── go.mod                          # Go module dependencies
├── go.sum
├── cdk.json                        # CDK configuration
├── README.md
└── cdk.out/                        # Synthesized CloudFormation (ignored)
```

---

## Prerequisites

Ensure the following are installed and configured:

* Go (>= 1.20)
* Node.js (>= 18)
* AWS CLI (configured with credentials)
* AWS CDK v2

Verify installation:

```
go version
node -v
npm -v
aws --version
cdk --version
```

---

## Setup Instructions

### 1. Clone the repository

```
git clone https://github.com/your-username/aurora-secret-rotation-cdk.git
cd aurora-secret-rotation-cdk
```

---

### 2. Install dependencies

```
go mod tidy
```

---

### 3. Bootstrap CDK (one-time per account/region)

```
cdk bootstrap aws://<account-id>/<region>
```

Example:

```
cdk bootstrap aws://123456789012/ap-south-1
```

---

### 4. Synthesize CloudFormation template

```
cdk synth
```

---

### 5. Deploy infrastructure

```
cdk deploy
```

---

## Key Features

### Secure Credential Management

* Passwords are automatically generated
* No hardcoded credentials in code
* Credentials stored securely in AWS Secrets Manager

### Serverless Database

* Aurora Serverless v2 scales automatically
* Pay only for actual usage
* No manual capacity planning required

### Production-Ready Networking

* Database deployed in private subnets
* NAT Gateway for controlled outbound access
* Security groups restrict unnecessary traffic

---

## Important Notes

* Ensure your AWS region supports Aurora PostgreSQL Serverless v2
* Engine versions must match region availability
* Secrets Manager password generation must exclude invalid characters for RDS

---

## Common Issues and Fixes

### CDK Bootstrap Error

If you see:

```
SSM parameter not found
```

Run:

```
cdk bootstrap
```

---

### Engine Version Errors

If deployment fails due to version issues:

* Use a supported Aurora PostgreSQL version for your region
* Avoid hardcoding unsupported versions

---

### Password Validation Errors

If you encounter password issues:

* Exclude invalid characters in secret generation
* Ensure compatibility with RDS password constraints

---

## Future Improvements

* Add automatic secret rotation using Lambda
* Integrate CloudWatch monitoring and alerts
* Add CI/CD pipeline (GitHub Actions)
* Add application layer to consume the database

---

## Conclusion

This project demonstrates a secure and scalable way to deploy cloud infrastructure using AWS CDK with Go. It focuses on best practices such as:

* Infrastructure as Code
* Secure credential management
* Serverless architecture
* Production-ready networking

---

## Author

Abinaya

---
