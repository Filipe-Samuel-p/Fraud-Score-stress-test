# Workload Identity do GKE: GSA usada pelos pods para LER imagens do Artifact Registry

# Service Account que o Kubernetes ServiceAccount irá impersonar
resource "google_service_account" "workload_identity" {
  account_id   = "fraud-score-gke"
  display_name = "workload_identity"
  project      = var.google_project_id
}

# Permissão apenas de leitura no Artifact Registry (pull da imagem docker)
resource "google_project_iam_member" "gke_artifact_reader" {
  project = var.google_project_id
  role    = "roles/artifactregistry.reader"
  member  = "serviceAccount:${google_service_account.workload_identity.email}"
}

# Binding entre o Kubernetes ServiceAccount e a GSA (Workload Identity do GKE)
resource "google_service_account_iam_member" "gke_wif_user" {
  service_account_id = google_service_account.workload_identity.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${var.google_project_id}.svc.id.goog[${var.k8s_namespace}/${var.k8s_service_account}]"
}
