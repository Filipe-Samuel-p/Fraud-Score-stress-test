resource "google_artifact_registry_repository" "fraud-score-artifact" {

  location      = var.region
  format        = "DOCKER"
  repository_id = "repo1"
  description   = "Artifact repo para a imagem docker da aplicação"

  depends_on = [google_project_service.enabled]
}