resource "helm_release" "argocd" {
  name             = "argocd"
  namespace        = "argocd"
  create_namespace = true

  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"
  version    = "7.7.11"

  depends_on = [google_container_cluster.fraud_socore_cluster]
}

resource "helm_release" "argocd_app" {
  name      = "fraud-score-app"
  namespace = "argocd"

  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argocd-apps"
  version    = "2.0.2"

  values = [
    yamlencode({
      applications = {
        "fraud-score" = {
          namespace = "argocd"
          finalizers = [
            "resources-finalizer.argocd.argoproj.io"
          ]
          project = "default"
          source = {
            repoURL        = "https://github.com/Filipe-Samuel-p/Fraud-Score-stress-test.git"
            targetRevision = "main"
            path           = "k8s/manifests"
          }
          destination = {
            server    = "https://kubernetes.default.svc"
            namespace = "default"
          }
          syncPolicy = {
            automated = {
              prune    = true
              selfHeal = true
            }
            syncOptions = [
              "CreateNamespace=true"
            ]
          }
        }
      }
    })
  ]

  depends_on = [helm_release.argocd]
}
