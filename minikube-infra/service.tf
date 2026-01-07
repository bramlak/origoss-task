resource "kubernetes_service" "origoss_task_service" {
  metadata {
    name      = "origoss-task-service"
    namespace = kubernetes_namespace.origoss_task.metadata[0].name

    annotations = {
      "prometheus.io/scrape" = "true"
      "prometheus.io/port"   = "8080"
      "prometheus.io/path"   = "/metrics"
    }
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