resource "null_resource" "ansible" {
  depends_on = [module.vm]

  provisioner "local-exec" {
    command = <<-EOT
      echo "[production]" > ../../ansible/inventory/hosts.ini
      echo "${module.vm.public_ip} ansible_user=test ansible_ssh_private_key_file=~/.ssh/id_ed25519" >> ../../ansible/inventory/hosts.ini
      
      sleep 30  # ждём загрузки VM
      
      cd ../../ansible
      ansible-playbook -i inventory/hosts.ini site.yml
    EOT
  }
}
