resource "kubernetes_namespace" "origoss_task" {
  metadata {
    name = "origoss-task"
  }
}