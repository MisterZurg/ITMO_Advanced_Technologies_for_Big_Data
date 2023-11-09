# Server deployment config on Selectel Cloud Platform
> [!IMPORTANT]
> Don't forget to add ./vars.tfvars
> ```sh
> # Server secrets vars.tfvars
> openstack_username = "..."
> openstack_password = "..." 
> ```

## Deploy server infrastructure
```sh
# Check if evth is ok 
terraform plan --var-file="vars.tfvars"
terraform apply --var-file="vars.tfvars"

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes
```

## Remove deployed server infrastructure
```sh
terraform destroy --var-file="vars.tfvars"
```