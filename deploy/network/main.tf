# Description of created resourses
# Инициализация Terraform и хранения Terraform State
# Можно вынести в отдельный файлик versions.tf
terraform {
required_version = ">= 0.14.0"
  required_providers {
    openstack = {
      source  = "terraform-provider-openstack/openstack"
      version = "~> 1.49.0"
    }
  }
}

# Configure the OpenStack provider
provider "openstack" {
  user_name           = var.openstack_username    # — Сервисный (или неочень) пользователя OpenStack;
  password            = var.openstack_password    # — пароль пользователя OpenStack;
  tenant_id           = var.project_id            # <project_id> — ID проекта облачной платформы;
  project_domain_name = var.domain_name           # <selectel_account> — номер аккаунта Selectel (номер договора). Можно посмотреть в панели управления в правом верхнем углу;
  user_domain_name    = var.domain_name           # <selectel_account> — ...
  auth_url            = var.auth_url
  region              = var.region
}

# Searching for selectel connect idonno
data "openstack_networking_network_v2" "external_net" {
  name     = var.router_external_net_name
  external = true
}

# Создание сети
resource "openstack_networking_router_v2" "router_1" {
  name                = var.router_name
  external_network_id = data.openstack_networking_network_v2.external_net.id
}

resource "openstack_networking_network_v2" "network_1" {
  name = var.network_name
}

resource "openstack_networking_subnet_v2" "subnet_1" {
  network_id      = openstack_networking_network_v2.network_1.id
  name            = var.subnet_cidr
  cidr            = var.subnet_cidr
}

resource "openstack_networking_router_interface_v2" "router_interface_1" {
  router_id = openstack_networking_router_v2.router_1.id
  subnet_id = openstack_networking_subnet_v2.subnet_1.id
}

// TODO : add floating ip
resource "openstack_networking_floatingip_v2" "floatingip_1" {
  pool = "external-network"
}