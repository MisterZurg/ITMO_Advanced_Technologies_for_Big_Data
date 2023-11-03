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

  backend "s3" {
    bucket                      = "<container_name>"
    endpoint                    = "s3.ru-1.storage.selcloud.ru"
    key                         = "<file_name>.tfstate"
    region                      = "ru-1"
    skip_region_validation      = true
    skip_credentials_validation = true
    access_key                  = "<access_key>"
    secret_key                  = "<secret_key>"
  }
}


# Configure the Selectel provider
# TODO
provider "selectel" {
  username    = var.username
  password    = var.password
  domain_name = var.domain_name
}


# Configure the OpenStack provider
# TODO
provider "openstack" {
  user_name           = var.username
  password            = var.password
  tenant_id           = var.project_id
  project_domain_name = var.domain_name
  user_domain_name    = var.domain_name
  auth_url            = var.auth_url
  region              = var.region
}


# Search 4 Network and Subnet
data "openstack_networking_network_v2" "my_net" {
  name = var.network_name
}

data "openstack_networking_subnet_v2" "my_subnet" {
  network_id = data.openstack_networking_network_v2.my_net.id
}


# Search datastore_type
data "selectel_dbaas_datastore_type_v1" "dt" {
  project_id = var.project_id
  region     = var.region
  filter {
    engine  = "postgresql"
    version = "15"
  }
}


# Creation af DB Cluster
resource "selectel_dbaas_datastore_v1" "datastore_1" {
  name       = "datastore-1"
  project_id = var.project_id
  region     = var.region
  type_id    = data.selectel_dbaas_datastore_type_v1.dt.datastore_types[0].id
  subnet_id  = data.openstack_networking_subnet_v2.my_subnet.id
  node_count = 3
  flavor {
    vcpus = 4
    ram   = 4096
    disk  = 32
  }

  pooler {
    mode = "transaction"
    size = 50
  }
}


# Create DB User
resource "selectel_dbaas_user_v1" "user_1" {
  project_id   = var.project_id
  region       = var.region
  datastore_id = selectel_dbaas_datastore_v1.datastore_1.id
  name         = "user"
  password     = "secret"
}


# Creation of DB
resource "selectel_dbaas_database_v1" "database_1" {
  project_id   = var.project_id
  region       = var.region
  datastore_id = selectel_dbaas_datastore_v1.datastore_1.id
  owner_id     = selectel_dbaas_user_v1.user_1.id
  name         = "db"
  lc_ctype     = "ru_RU.utf8"
  lc_collate   = "ru_RU.utf8"
}


# Connect Extension
data "selectel_dbaas_available_extension_v1" "ae" {
  project_id = var.project_id
  region     = var.region
  filter {
    name = "hstore"
  }
}

resource "selectel_dbaas_extension_v1" "extension_1" {
  project_id             = var.project_id
  region                 = var.region
  datastore_id           = selectel_dbaas_datastore_v1.datastore_1.id
  database_id            = selectel_dbaas_database_v1.database_1.id
  available_extension_id = data.selectel_dbaas_available_extension_v1.ae.available_extensions[0].id
}