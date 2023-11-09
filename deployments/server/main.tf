# Description of created resourses
# Инициализация Terraform и хранения Terraform State
# TODO : Check actual versions
terraform {
required_version = ">= 0.14.0"
  required_providers {
    selectel = {
      source  = "selectel/selectel"
      version = "~> 4.0.1"
    }

    openstack = {
      source  = "terraform-provider-openstack/openstack"
      version = "~> 1.49.0"
    }
  }
}


# Configure the Selectel provider
# TODO
# provider "selectel" {
#   username    = var.selectel_username
#   password    = var.selectel_password
#   domain_name = var.selectel_domain_name
# }

# <project_id> — ID проекта облачной платформы;
# <user_name> — пользователь OpenStack, привязанный к проекту облачной платформы;
# <user_password> — пароль пользователя OpenStack;
# <pool> — пул, в котором будет развернута инфраструктура.

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

resource "random_string" "random_name" {
  length  = 5
  special = false
}

# Server cofig
data "openstack_images_image_v2" "image_1" {
  name        = var.server_image_name
  visibility  = "public"
  most_recent = true
}

resource "openstack_compute_flavor_v2" "flavor_1" {
  name      = "flavor-${random_string.random_name.result}"
  ram       = var.server_ram_mb
  vcpus     = var.server_vcpus
  disk      = var.server_root_disk_gb
  is_public = var.is_public

  lifecycle {
    create_before_destroy = true
  }
}

resource "openstack_compute_instance_v2" "instance_1" {
  name              = var.server_name
  image_id          = data.openstack_images_image_v2.image_1.id
  flavor_id         = openstack_compute_flavor_v2.flavor_1.id
  availability_zone = var.server_zone
  admin_pass 	      = "12345678"

  network {
    port = openstack_networking_port_v2.port_1.id
  }

  lifecycle {
    ignore_changes = [image_id]
  }

  vendor_options {
    ignore_resize_confirmation = true
  }
}

# Network
# Search 4 Network and Subnet
# data "openstack_networking_network_v2" "external_net" {
#   name = var.network_name
# }

# Запрос ID внешней сети по имени
data "openstack_networking_network_v2" "network_1" {
  name = "sus-network_1"
}

data "openstack_networking_subnet_v2" "subnet_1" {
  network_id = data.openstack_networking_network_v2.network_1.id
}

resource "openstack_networking_port_v2" "port_1" {
  name       = "${var.server_name}-eth0"
  network_id = data.openstack_networking_network_v2.network_1.id

  fixed_ip {
    subnet_id = data.openstack_networking_subnet_v2.subnet_1.id
  }
}


# Search datastore_type
# data "selectel_dbaas_datastore_type_v1" "dt" {
#   project_id = var.project_id
#   region     = var.region
#   filter {
#     engine  = "postgresql"
#     version = "15"
#   }
# }


# Creation af DB Cluster
# resource "selectel_dbaas_datastore_v1" "datastore_1" {
#   name       = "datastore-1"
#   project_id = var.project_id
#   region     = var.region
#   type_id    = data.selectel_dbaas_datastore_type_v1.dt.datastore_types[0].id
#   subnet_id  = data.openstack_networking_subnet_v2.my_subnet.id
#   node_count = 3
#   flavor {
#     vcpus = 4
#     ram   = 4096
#     disk  = 32
#   }

#   pooler {
#     mode = "transaction"
#     size = 50
#   }
# }


# Create DB User
# resource "selectel_dbaas_user_v1" "user_1" {
#   project_id   = var.project_id
#   region       = var.region
#   datastore_id = selectel_dbaas_datastore_v1.datastore_1.id
#   name         = "user"
#   password     = "secret"
# }


# Creation of DB
# resource "selectel_dbaas_database_v1" "database_1" {
#   project_id   = var.project_id
#   region       = var.region
#   datastore_id = selectel_dbaas_datastore_v1.datastore_1.id
#   owner_id     = selectel_dbaas_user_v1.user_1.id
#   name         = "db"
#   lc_ctype     = "ru_RU.utf8"
#   lc_collate   = "ru_RU.utf8"
# }


# Connect Extension
# data "selectel_dbaas_available_extension_v1" "ae" {
#   project_id = var.project_id
#   region     = var.region
#   filter {
#     name = "hstore"
#   }
# }

# resource "selectel_dbaas_extension_v1" "extension_1" {
#   project_id             = var.project_id
#   region                 = var.region
#   datastore_id           = selectel_dbaas_datastore_v1.datastore_1.id
#   database_id            = selectel_dbaas_database_v1.database_1.id
#   available_extension_id = data.selectel_dbaas_available_extension_v1.ae.available_extensions[0].id
# }