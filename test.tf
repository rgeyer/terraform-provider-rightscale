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
    description = "Description is the best!"
  }
}

resource "rightscale_server" "terraform-server-test" {
  server {
    name = "Test from Terraform you hoe"
#    deployment_href = "${rightscale_deployment.terraform-test.id}"
    deployment_href = "/api/deployments/104278"
    instance {
      server_template_href = "/api/server_templates/366194004"
      cloud_href = "/api/clouds/1"
    }
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
