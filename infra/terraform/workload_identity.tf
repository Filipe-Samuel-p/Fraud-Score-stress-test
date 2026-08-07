

resource "google_service_account" "workload_identity" { # Service Account que o Kubernetes ServiceAccount irá impersonar
  account_id   = "fraud-score-gke"
  display_name = "workload_identity"
  project      = var.google_project_id
}


resource "google_project_iam_member" "gke_artifact_reader" { # Permissão apenas de leitura no Artifact Registry (pull da imagem docker)
  project = var.google_project_id
  role    = "roles/artifactregistry.reader"
  member  = "serviceAccount:${google_service_account.workload_identity.email}"
}


resource "google_service_account_iam_member" "gke_wif_user" { # Binding entre o Kubernetes ServiceAccount e a GSA (Workload Identity do GKE)
  service_account_id = google_service_account.workload_identity.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${var.google_project_id}.svc.id.goog[${var.k8s_namespace}/${var.k8s_service_account}]"

  
  depends_on = [google_container_cluster.fraud_socore_cluster]
}
