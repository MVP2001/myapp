resource "yandex_vpc_network" "main" {
  name = "myapp-network"
}

resource "yandex_vpc_subnet" "main" {
  name           = "myapp-subnet"
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.main.id
  v4_cidr_blocks = ["10.0.0.0/24"]
}
