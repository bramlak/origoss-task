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
