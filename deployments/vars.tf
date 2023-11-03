# # Пул Облачной платформы 
# variable "region" {
#   default = "ru-9"
# }

# # Значение SSH-ключа для доступа к облачному серверу
# variable "public_key" {
#   default = "key_value"
# }

# # Сегмент пула
# variable "az_zone" {
#   default = "ru-3b"
# }

# # Тип сетевого диска, из которого создается сервер
# variable "volume_type" {
#   default = "fast.ru-3b"
# }

# # CIDR подсети
# variable "subnet_cidr" {
#   default = "10.10.0.0/24"
# }



variable "username" {}

variable "password" {}

variable "domain_name" {}

variable "project_id" {}

variable "region" {}

variable "auth_url" {
  default = "https://api.selvpc.ru/identity/v3"
}

variable "network_name" {
  default = "nat"
}