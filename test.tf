variable "account_id" { }
variable "api_host"   { }
variable "refresh_token" {}

provider "rightscale" {
  account_id    = "${var.account_id}"
  api_host      = "${var.api_host}"
  refresh_token = "${var.refresh_token}"
}

resource "rightscale_deployment" "terraform-test" {
  deployment {
    name        = "Test Created from Terraform"
    description = "Description"
  }
}

output "rightscale_deployment.terraform-test.href" {
  value = "${rightscale_deployment.terraform-test.id}"
}

output "rightscale_deployment.terraform-test.name" {
  value = "${rightscale_deployment.terraform-test.name}"
}

output "rightscale_deployment.terraform-test.description" {
  value = "${rightscale_deployment.terraform-test.description}"
}
