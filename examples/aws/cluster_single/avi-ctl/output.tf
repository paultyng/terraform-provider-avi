output Controller_PublicIP {
  value = aws_instance.avi-controller[*].public_ip
}
output Controller_PrivateIP {
  value = aws_instance.avi-controller[*].private_ip
}
output Msg {
  value = "It will take 15 min for the controller to be ready"
}
