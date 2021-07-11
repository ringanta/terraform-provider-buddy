# Example Usage Buddy Provider

The following Terraform project demonstrates usage of resources from the Buddy Terraform provider

## Prerequisites

You need the following to run this project:

- [tfenv](https://github.com/tfutils/tfenv)

## Setup

Follow the following steps to run the project:

1. Install required terraform version
   ```shell
   tfenv install min-required
   tfenv use min-required
   ```
1. Export Buddy credentials
   ```shell
    export BUDDY_URL=<Buddy API URL>
    export BUDDY_TOKEN=<Buddy personal access token>
   ```
1. Initialize terraform project
   ```shell
   terraform init
   ```
1. Generate and review execution plan
   ```shell
   terraform plan -out=tfplan.out
   ```
1. Apply the execution plan
   ```shell
   terraform apply tfplan.out
   ```

## Clean up

Should you want to clean up all resources created in this project, run
1. Generate and review execution plan to destroy the resources
   ```shell
   terraform plan -destroy -out=tfplan.out
   ```
1. Apply the execution plan
   ```shell
   terraform apply tfplan.out
   ```
