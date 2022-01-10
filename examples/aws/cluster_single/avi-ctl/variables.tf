variable "old_password" {}

variable "iam_profile" {
  default = "remo-avi-controller"
}
variable "vpc_id_fse" {
  description = "VPC ID"
}
variable "admin_username" {
  description = "User Admin"
}
variable "admin_password" {
  description = "Admin Password"
}
variable "se_subnet_id" {
  description = "Subnet ID"
}
# variable "ami-image" {
#   type        = "string"
#   description = "Image to use"
# }
variable "image-size" {
  description = "Image size"
}
variable "controller_name" {
  description = "Controller Name"
}
variable "public_ip" {
  description = "If you want a Public IP"
}
variable "shutdown_rules" {
  description = "Add noshut if you want to keep it up"
}
variable "department_name" {
  description = "Department Name"
}
variable "myregion" {
  description = "Region"
}
variable "remo_sg_name" {
  description = "Name of the Sec Groups"
}
variable "profile" {
  description = "Profile to use "
}
variable "owner" {
  description = "owner"
}
variable "ami-image" {
  default = {
    us-west-1 = "ami-03baad459ee4a3980"
    us-west-2 =  "ami-04d08b852b47a5876"
   }
  }


# You could spin up more than 1 but use the other repo for cluster
variable "controller_count" {
  default = 1
}
variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "api_version" {
  default = "21.1.1"
}
variable "dns_ip" {
  default = "8.8.8.8"
}
