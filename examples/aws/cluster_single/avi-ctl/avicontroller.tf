data "template_file" "userdata" {
  template                  = file("files/userdata.json")
  vars                      = {
    avi_controller_password = var.admin_password
    dns                     = var.dns_ip[0]
    dns1                    = var.dns_ip[1]
    search_domain           = var.search_default_domain
    welcome_banner          = var.welcome_banner
    ntp_server1             = var.ntp_servers[0]
    ntp_server2             = var.ntp_servers[1]
    ntp_server3             = var.ntp_servers[2]
    ntp_server4             = var.ntp_servers[3]
  }
}
resource "aws_instance" "avi-controller" {
  count                       = var.controller_count
  user_data                   = count.index == 0 ? data.template_file.userdata.rendered : ""
  ami                         = lookup(var.ami-image, var.myregion)
  associate_public_ip_address = var.public_ip
  instance_type               = var.image-size
  subnet_id                   = var.se_subnet_id
  key_name                    = aws_key_pair.generated.key_name
  iam_instance_profile        = var.iam_profile
  vpc_security_group_ids = [
    aws_security_group.remo_sg.id,
  ]
  tags = {
    Name            = "ctl-avi-tf-${count.index + 1}"
    dept            = var.department_name
    shutdown_policy = var.shutdown_rules
    owner           = var.owner
  }
}

