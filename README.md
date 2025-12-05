# origoss-task

<div align="center">
  <img src="https://origoss.com/images/logo.svg" alt="Logo" width="160" height="80">

  <h3 align="center">Origoss Task</h3>
  <p align="center">An HTTP server and its infrastructure</p>
</div>

## About This Project
This project demonstrates a complete end-to-end workflow for deploying a simple Go HTTP server with modern infrastructure practices.
## Features:
- A **Go HTTP server** that responds with `Hello, World!` at `/`
- A **Dockerfile** for server containerization
- A **GitHub CI workflow** for building and pushing the server changes to Docker Hub
- **Kubernetes manifests** for deployment, service, and ingress
- **Terraform infrastructure** replicating the Kubernetes resources plus some additional enhancements
- Deployment is designed for a **Minikube** cluster for local development


## Languages and Technologies

- [![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#) Server Programming Language
- [![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff)](#) Containerization
- [![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=fff)](#) Orchestration
- [![Terraform](https://img.shields.io/badge/Terraform-844FBA?logo=terraform&logoColor=fff)](#) Infrastructure-as-Code
- [![Minikube](https://img.shields.io/badge/Minikube-000000?logo=minikube&logoColor=fff)](#) Local Cluster
- [![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?logo=github-actions&logoColor=white)](#) CI


## Set Up
### Prerequisites

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Minikube](https://minikube.sigs.k8s.io/)
- [kubectl](https://kubernetes.io/)
- [Terraform](https://developer.hashicorp.com/terraform/)


### Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/origoss-task.git
cd origoss-task
```

2. **Start Minikube**

```bash
minikube start
```

3. **Add local domain to host list**

Aquire Minikube IP and add to `/etc/hosts`
```bash
echo "$(minikube ip) origoss-task.local" | sudo tee -a /etc/hosts
```


## Usage

There are a two ways to use the Minikube cluster. There are kubernetes manifest YAML files and also a terraform infrastructure.

### Via Kubernetes Manifest

1. **Apply the manifests**

```bash
cd kubernetes
kubectl apply -f deployment.yaml -f service.yaml -f ingress.yaml
```

2. **Access the server**

```bash
curl origoss-task.local
```


### Via Terraform

1. **Initialize Terraform**

```bash
cd minikube-infra
terraform init
```

2. **Apply the infrastructure**

```bash
terraform apply
```

3. **Access the server**

```bash
curl origoss-task.local
```


## Utilising CI
To use CI capabilities for custom development, follow these steps:

### Fork The Repository
Follow the official [Forking - GitHub documentation](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo)
### Edit Workflow File
1. Change the image reference to a custom one
2. Generate Docker Hub Personal Access Token (`Settings` > `Personal access tokens` > `Generate new token`)
3. Set The PAT for GitHub as `DOCKER_HUB_TOKEN` (`Settings` > `Secrets and variables`> `Actions` > `New Repository Secret`)
### Change Image
#### Kubernetes Manifest
1. Change the image reference in  deployment.yaml
#### Terraform
1. Create a `terraform.tfvars` file with the updated image reference or change the default image value in `variables.tf`
### Trigger GitHub Action
1. Navigate to `Actions`
2. Find and choose the action from left side menu
3. Find the `This workflow has a workflow_dispatch event trigger.` field and choose `Run workflow`
4. Continue with the [usage steps](#usage)



## Notes
* Feel free to change local domain from `origoss-task.local` to any custom one
* Feel free to change replica count according to custom needs
* This approach aims to satisfy the main requirements: simplicity and efficiency. If more complexity is acceptable, a few improvements could be added:
 - AWS EKS cluster usage (possibly with additional features)
 - additional Go endpoints for Kubernetes probes
 - etc.

