resource "google_storage_bucket" "fraud-score-bucket" {
  name     = "fraud-score-backend"
  location = "us-central1"
}