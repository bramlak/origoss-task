resource "kubernetes_service" "origoss_task_service" {
  metadata {
    name      = "origoss-task-service"
    namespace = kubernetes_namespace.origoss_task.metadata[0].name
  }

  spec {
    selector = {
      app = "origoss-task"
    }

    port {
      port        = 80
      target_port = 8080
    }

    type = "ClusterIP"
  }
}