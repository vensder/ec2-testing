# Simple Golang Web App in Docker for AWS EC2/ECS/Elastic Beanstalk testing

![Go](https://github.com/vensder/ec2-testing/workflows/Go/badge.svg) [![Build Status](https://travis-ci.org/vensder/ec2-testing.svg?branch=master)](https://travis-ci.org/vensder/ec2-testing) ![Docker](https://img.shields.io/docker/build/vensder/ec2-testing)

If you run this container in AWS Environment, the Web application will show you page with some meta-information about the current environment and environment variables inside the docker container.

![alt text](./img/screenshot.png?raw=true)

## How to run pre-built docker container on AWS

```bash
git clone https://github.com/vensder/ec2-testing.git
cd ec2-testing
```

Export your AWS Access keys:

```bash
export AWS_ACCESS_KEY_ID="XXXXXXXXXXXXXXXXXXXX"
export AWS_SECRET_ACCESS_KEY='YYzYYzzYYYYzYYYYYYYzYYYYzzzYzYzYYzYzzzzY'
```

Run helper script to configure S3 backend to store the Terraform state file:

```bash
cd terraform
./s3-terraform-state-helper.sh
```

Run Terraform plan and apply:

```bash
terraform plan
terraform apply
```

Copy the IP address from the terraform output and open it in the browser, for example:

```bash
Outputs:

instance_ip_addr = 35.172.233.117
```

Open http://35.172.233.117

## How to run pre-built docker container locally

```bash
docker run -d -p 80:8080 --rm --name ec2-test vensder/ec2-testing
```

## How to build image and run docker container locally

```bash
git clone https://github.com/vensder/ec2-testing.git
cd ec2-testing
docker build . -t ec2-testing
docker run -d -p 80:8080 --rm --name ec2-test ec2-testing
```

Open [http://localhost/Hi%20there!](http://localhost/Hi%20there!) or any other random path after slash in your web browser.

View logs:

```bash
docker logs ec2-testing
```

Stop container:

```bash
docker stop ec2-testing
```

## How to run pre-built docker container in Elastic Beanstalk environment

You can use this container for testing of AWS Elastic Beanstalk Blue/Green deployment (see Dockerrun.aws.json) - just set environment variable "color" inappropriate color for each environment. You can create an Elastic Beanstalk application and environment, using Elastic Beanstalk command-line interface:

```bash
eb init -p docker ec2-testing
eb create ec2-testing-blue  --instance_type t3.nano --region us-east-1 --envvars color=blue
eb create ec2-testing-green --instance_type t3.nano --region us-east-1 --envvars color=green
eb swap ec2-testing-blue --destination_name ec2-testing-green
```

Terminate environment:

```bash
eb terminate ec2-testing-blue
```
