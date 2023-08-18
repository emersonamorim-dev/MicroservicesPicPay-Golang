provider "google" {
  credentials = file("<CAMINHO_PARA_SEU_ARQUIVO_DE_CREDENCIAIS>.json")
  project     = "<SEU_ID_DO_PROJETO>"
  region      = "us-central1"
}

resource "google_container_cluster" "primary" {
  name     = "my-gke-cluster"
  location = "us-central1"

  remove_default_node_pool = true
  initial_node_count       = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}
