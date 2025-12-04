variable "replicas" {
  description = "Number of replicas of server pods"
  type = number
  default = 1
}

variable "image" {
  description = "The name of the server image"
  type = string
  default = "bramlak/origoss-task-server:latest"
}

variable "kubernetes_config_path" {
  description = "Path to the kubeconfig file used by the Kubernetes provider"
  type = string
  default = "~/.kube/config"
}

variable "host_domain" {
  description = "The host domain used to access the server"
  type = string
  default = "origoss-task.local"
}

