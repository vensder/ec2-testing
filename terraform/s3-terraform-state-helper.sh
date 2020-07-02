#!/usr/bin/env bash

set -eufx -o pipefail

if aws s3 ls | grep terraform-state; then
	echo "s3 bucket present, creating backend config file..."
	_terraform_backend_s3_bucket=$(aws s3 ls | grep -Eo 'terraform-state-[0-9]{12}')
	cat > backend.tf << EOF
terraform {
  backend "s3" {
    bucket = "$_terraform_backend_s3_bucket"
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