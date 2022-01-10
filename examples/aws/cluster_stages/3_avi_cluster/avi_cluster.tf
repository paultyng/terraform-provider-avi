
data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-1"]
  }
}

data "aws_instance" "avi_controller_2" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-2"]
  }
}

data "aws_instance" "avi_controller_3" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-3"]
  }
}

resource "avi_cluster" "aws_cluster" {
  name =  var.cluster_name
  nodes {
    ip {
      type = "V4"
      addr = data.aws_instance.avi_controller.private_ip
    }
    name = "node01"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.aws_instance.avi_controller_2.private_ip
    }
    name = "node02"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.aws_instance.avi_controller_3.private_ip
    }
    name = "node03"
  }
}

resource "avi_systemconfiguration" "avi_system" {
    common_criteria_mode      = false
    uuid                      = "default-uuid"
    welcome_workflow_complete = true
    }

resource "avi_backupconfiguration" "backup_config" {
  name                     = "Backup-Configuration"
  tenant_ref               =  "admin"
  #tenant_ref              = data.avi_tenant.default_tenant.id
  save_local               = true
  maximum_backups_stored   = 4
  backup_passphrase        =  var.avi_password
  configpb_attributes {
     version = 1
   }
}
