resource "google_service_account" "github_sa" {
  account_id   = "fraud-score-github"
  display_name = "github_sa"
  project      = var.google_project_id
}