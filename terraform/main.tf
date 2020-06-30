terraform {
  required_version = "= 0.12.20"
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_key_pair" "vensder" {
  key_name   = "vensder-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDfKSiUnU/u8WSn1VWab7p4xJuZgOnN9ydbKUGAGGU5jCSMH2njWsU4KRtdvrWGcSkFSNiufzFs6cXqIOi1fhM6t5weJGLsqL9cpRQgS7oUSMOeo6vii45xksNUU4tLHjz6RCCOnKmleUwPmFffVrWiCgnzwQcWBQQ7wvLoUt2qC58b9dFuh2ueNizanWaCTXCL2675on6kvB9rfbQXJ1cZy+y/z0Tsxp74wv3AuzclZDUSoX5xcEwJr7VN6HPfjAm0ZpJxHPgjPXoceqzWVs896pEfOqxAKH0eK4FW5FBM8EQSjspjpehYM4pWT4h5VmhEznYcD0kBaAopvdxSwg4P dmitrii@dmitrii-laptop"
}

resource "aws_security_group" "allow_ssh" {
  name        = "allow_ssh"
  description = "Allow SSH"

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Terraform removes the default rule
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "allow_ssh"
  }
}

resource "aws_security_group" "allow_http" {
  name        = "allow_http"
  description = "Allow HTTP"

  ingress {
    description = "SSH"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Terraform removes the default rule
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "allow_http"
  }
}

resource "aws_instance" "amazon-linux" {
  key_name        = "${aws_key_pair.vensder.key_name}"
  ami             = "ami-0e9089763828757e1"
  instance_type   = "t3.nano"
  security_groups = ["${aws_security_group.allow_ssh.name}", "${aws_security_group.allow_http.name}"]
  user_data = <<EOF
    #!/usr/bin/env bash
    sudo yum update -y
    sudo yum install docker -y
    sudo service docker start
    sudo docker run -d -p 80:8080 --rm --name ec2-testing vensder/ec2-testing
  EOF
  tags = {
    Name = "amazon-linux"
  }
}
