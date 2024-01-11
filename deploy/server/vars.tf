# ————————————————————————————
#         Slectel Vars
# ————————————————————————————
variable "selectel_username" {}
variable "selectel_password" {}
variable "selectel_domain_name" {}

# ————————————————————————————
#        OpenStack Vars
# ————————————————————————————
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

# Пул Облачной платформы
variable "region" {
  default = "ru-9"
}


#        Server Stats Vars
# ————————————————————————————
variable "server_image_name" {
  default = "Ubuntu 22.04 LTS 64-bit"
}

variable "server_name" {
  default = "tf_Aboba_server"
}

variable "server_vcpus" {
  default = 1
}

variable "server_ram_mb" {
  default = 8192
}

variable "server_root_disk_gb" {
  default = 8
}

variable "is_public" {
  default = false
}

variable "server_zone" {
  default = "ru-9a"
}

# # Значение SSH-ключа для доступа к облачному серверу
# variable "public_key" {
#   default = "key_value"
# }


# Network Config
variable "network_name" {
  default = "aboba-network"
}