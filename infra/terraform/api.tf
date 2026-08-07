
locals {
  gcp_services = [
    "compute.googleapis.com",              
    "container.googleapis.com",           
    "artifactregistry.googleapis.com",     
    "iam.googleapis.com",                  
    "iamcredentials.googleapis.com",       
    "sts.googleapis.com",                 
    "cloudresourcemanager.googleapis.com", 
    "storage.googleapis.com",             
  ]
}

resource "google_project_service" "enabled" {
  for_each = toset(local.gcp_services)

  project = var.google_project_id
  service = each.value

  disable_on_destroy = false
}
