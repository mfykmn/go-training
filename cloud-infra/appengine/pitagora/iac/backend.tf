#####################################
# Management of tfstate
#####################################
terraform {
  required_version = "0.12.0"

  backend "gcs" {
    bucket = "tf-state-pitagora"
    prefix = "pitagora/state"
  }
}