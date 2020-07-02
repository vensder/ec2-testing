# Simple Golang Web App in Docker for AWS EC2/ECS testing

![Go](https://github.com/vensder/ec2-testing/workflows/Go/badge.svg) [![Build Status](https://travis-ci.org/vensder/ec2-testing.svg?branch=master)](https://travis-ci.org/vensder/ec2-testing) ![Docker](https://img.shields.io/docker/build/vensder/ec2-testing)

If you run this container in AWS Environment, the Web application will show you page with some meta-information about the current environment:

* ami-id
* hostname
* instance-id
* instance-type
* local-hostname
* local-ipv4
* public-ipv4
* public-keys
* security-groups

And environment variables inside the docker container too.


### How to run:

``
docker run -d -p 8080:8080 --rm --name ec2-testing vensder/ec2-testing
``

Open [http://localhost:8080/Hi%20there!](http://localhost:8080/Hi%20there!) or any other random path after slash in your web browser.

View logs:

``
docker logs ec2-testing
``

Stop container:

``
docker stop ec2-testing
``

You can use this container for testing of AWS Elastic Beanstalk Blue/Green deployment (see Dockerrun.aws.json) - just set environment variable "color" in appropriate color for each environment.

