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
    name        = "001 - Test Created from Terraform v2"
    description = "Description is the best!"
  }
}

resource "rightscale_server" "terraform-server-test" {
  server {
    name = "Test from Terraform"
    deployment_href = "${rightscale_deployment.terraform-test.id}"
#    deployment_href = "/api/deployments/104278"
    instance {
      server_template_href = "/api/server_templates/366194004"
      cloud_href = "/api/clouds/1"

    }
  }
  count = 2
}

output "rightscale_deployment.terraform-test.href" {
  value = "${rightscale_deployment.terraform-test.id}"
}
