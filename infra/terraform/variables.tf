variable "google_project_id" {
  type        = string
  description = "Variável que possui o id do projeto"
}

variable "region" {
  type        = string
  description = "Variável que possui a região do projeto"

}

variable "gke_cluster_name" {
  type    = string
  default = "fraud-score-cluster"
}

variable "k8s_namespace" {
  type    = string
  default = "default"
}

variable "k8s_service_account" {
  type        = string
  description = "Nome do Kubernetes ServiceAccount vinculado à GSA via Workload Identity"
  default     = "fraud-score-sa"
}

variable "github_owner" {
  type        = string
  description = "Owner/organização do repositório no GitHub"
  default     = "Filipe-Samuel-p"
}

variable "github_repository" {
  type        = string
  description = "Repositório no formato owner/repo autorizado a impersonar a GSA do CI"
  default     = "https://github.com/Filipe-Samuel-p/Fraud-Score-stress-test.git"
}

