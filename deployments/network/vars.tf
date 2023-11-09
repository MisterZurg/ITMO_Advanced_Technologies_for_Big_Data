variable "openstack_username" {}
variable "openstack_password" {}

# ID of Cloud Project : "Sus-Project"
variable "project_id" {
    default = "b82119825e2745cc852741ad01b98ece"
}

variable "domain_name" {
    default = 280331
}

variable "auth_url" {
  default = "https://api.selvpc.ru/identity/v3"
}

variable "region" {
    default = "ru-9"
}

variable "router_external_net_name" {
  default = "external-network"
}

variable "router_name" {
  default = "sus-router_1"
}

variable "network_name" {
  default = "sus-network_1"
}

variable "subnet_cidr" {
  default = "192.168.0.0/24"
}