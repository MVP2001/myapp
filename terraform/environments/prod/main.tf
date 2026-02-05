terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "~> 0.99"
    }
  }
  backend "local" {
    path = "terraform.tfstate"
  }
}

provider "yandex" {
  token     = var.yc_token
  cloud_id  = var.yc_cloud_id
  folder_id = var.yc_folder_id
  zone      = "ru-central1-a"
}

module "vpc" {
  source = "../../modules/vpc"
}

module "vm" {
  source = "../../modules/vm"
  
  subnet_id      = module.vpc.subnet_id
  ssh_public_key = var.ssh_public_key
}

output "server_ip" {
  value = module.vm.public_ip
}
