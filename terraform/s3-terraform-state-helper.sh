#!/usr/bin/env bash

set -ex

if aws s3 ls | grep terraform-state; then
	echo "s3 bucket present, creating backend config file..."
	TERRAFORM_BACKEND_S3_BUCKET=$(aws s3 ls | grep -Eo 'terraform-state-[0-9]{12}')
	cat > backend.tf << EOF
terraform {
  backend "s3" {
    bucket = "$TERRAFORM_BACKEND_S3_BUCKET"
    key = "terraform_main.tfstate"
    region = "us-east-1"
  }
}
EOF
else
	cd ./terraform_backend/ || echo "terraform_backend dir not exists" && exit 1
	terraform init
	terraform apply
fi