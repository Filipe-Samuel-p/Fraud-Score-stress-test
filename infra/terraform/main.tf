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
  }
}

provider "google" {
  project = var.google_project_id
  region  = "us-central1"
}