# Terraform Provider Buddy

Run the following command to build the provider

```shell
make build
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to apply configuration

```shell
terraform init
terraform plan -out=tfplan.out
terraform apply tfplan.out
```
