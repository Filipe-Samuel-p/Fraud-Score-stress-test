resource "google_container_cluster" "fraud_socore_cluster" {
  name     = var.gke_cluster_name
  project  = var.google_project_id
  location = var.region

  enable_autopilot    = true
  deletion_protection = false
}