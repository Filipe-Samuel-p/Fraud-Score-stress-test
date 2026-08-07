terraform {
  backend "gcs" {
    bucket = "fraud-score-backend"
    prefix = "terraform/state"
  }

  required_version = "1.15.8"

  required_providers {
    google = {
      source = "hashicorp/google"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 3.0"
    }
  }
}

provider "google" {
  project = var.google_project_id
  region  = var.region
}

data "google_client_config" "default" {}

provider "helm" {
  kubernetes = {
    host                   = "https://${google_container_cluster.fraud_socore_cluster.endpoint}"
    token                  = data.google_client_config.default.access_token
    cluster_ca_certificate = base64decode(google_container_cluster.fraud_socore_cluster.master_auth[0].cluster_ca_certificate)
  }
}