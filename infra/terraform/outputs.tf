output "gke_cluster_name" {
  value = google_container_cluster.fraud_socore_cluster.name
}

output "github_ci_service_account" {
  description = "Email da GSA sem chave usada pelo GitHub Actions (impersonation via WIF)"
  value       = google_service_account.github_sa.email
}

output "gke_workload_identity_service_account" {
  description = "Email da GSA vinculada ao Kubernetes ServiceAccount (Workload Identity do GKE)"
  value       = google_service_account.workload_identity.email
}

output "wif_provider_name" {
  description = "Nome completo do provider WIF para uso no google-github-actions/auth"
  value       = google_iam_workload_identity_pool_provider.github_provider.name
}
