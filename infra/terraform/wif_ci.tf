data "google_project" "current" {
  project_id = var.google_project_id
}

resource "google_iam_workload_identity_pool" "wif_pool" { # Workload Identity Pool: repositório lógico que agrupa as identidades externa
  workload_identity_pool_id = "fraud-score-pool"
  display_name              = "fraud-score-pool"
  disabled                  = false
  mode                      = "FEDERATION_ONLY"
}


resource "google_iam_workload_identity_pool_provider" "github_provider" { # Workload Identity Provider: conexão OIDC com o emissor do GitHub Actions
  workload_identity_pool_id          = google_iam_workload_identity_pool.wif_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  display_name                       = "github-provider"


  attribute_mapping = { #   # Mapeia os claims do token OIDC do GitHub para atributos do GCP
    "google.subject"             = "assertion.sub"
    "attribute.repository"       = "assertion.repository"
    "attribute.repository_owner" = "assertion.repository_owner"
  }

  attribute_condition = "assertion.repository_owner == '${var.github_owner}'"

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}


resource "google_service_account_iam_member" "github_wif_user" {
  service_account_id = google_service_account.github_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/projects/${data.google_project.current.number}/locations/global/workloadIdentityPools/${google_iam_workload_identity_pool.wif_pool.workload_identity_pool_id}/attribute.repository/${var.github_repository}"
}

resource "google_project_iam_member" "github_sa_artifact_writer" {
  project = var.google_project_id
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.github_sa.email}"
}
