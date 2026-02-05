resource "yandex_compute_instance" "app" {
  name        = "myapp-server"
  platform_id = "standard-v3"  # Intel Ice Lake
  
  resources {
    cores         = 8
    memory        = 12
    core_fraction = 100
  }

  boot_disk {
    initialize_params {
      image_id = "fd8..."  # Ubuntu 22.04 LTS (найдём через CLI)
      size     = 100
      type     = "network-ssd"
    }
  }

  network_interface {
    subnet_id = var.subnet_id
    nat       = true  # Публичный IP
  }

  metadata = {
    ssh-keys = "ubuntu:${file(var.ssh_public_key)}"
    user-data = templatefile("${path.module}/cloud-init.yml", {
      username = "test"
    })
  }
}
