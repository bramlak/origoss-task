resource "kubernetes_ingress_v1" "origoss_task_ingress" {
  metadata {
    name      = "origoss-task-ingress"
    namespace = kubernetes_namespace.origoss_task.metadata[0].name
  }

  spec {
    rule {
      host = var.host_domain

      http {
        path {
          path      = "/"
          path_type = "Prefix"

          backend {
            service {
              name = kubernetes_service.origoss_task_service.metadata[0].name
              port {
                number = 80
              }
            }
          }
        }
      }
    }
  }
}