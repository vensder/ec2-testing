output "instance_ip_addr" {
  value = "${aws_instance.amazon-linux.public_ip}"
}