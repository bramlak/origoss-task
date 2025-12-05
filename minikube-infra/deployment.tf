resource "kubernetes_deployment" "origoss_task_server" {
  metadata {
    name      = "origoss-task-server"
    namespace = kubernetes_namespace.origoss_task.metadata[0].name
    labels = {
      app = "origoss-task"
    }
  }

  spec {
    replicas = var.replicas

    selector {
      match_labels = {
        app = "origoss-task"
      }
    }

    template {
      metadata {
        labels = {
          app = "origoss-task"
        }
      }

      spec {
        container {
          image = var.image
          name  = "server"
          image_pull_policy = "Always"
          env {
            name  = "PORT"
            value = "8080"
          }

          port {
            container_port = 8080
          }

          resources {
            requests = {
              cpu    = "100m"
              memory = "128Mi"
            }
            limits = {
              cpu    = "300"
              memory = "256Mi"
            }
          }

          liveness_probe {
            http_get {
              path = "/healthz"
              port = 8080
            }
            initial_delay_seconds = 10
            period_seconds        = 10
          }
        }
      }
    }
  }
}