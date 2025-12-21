# CODE Ecosystem

A complete GitOps-based CI/CD learning environment demonstrating Kubernetes-native automation with Argo Projects and Istio Service Mesh.

## 📖 Overview

This project showcases a production-like continuous delivery pipeline running entirely in a local Kubernetes cluster. It demonstrates:

- **GitOps principles** with ArgoCD
- **Event-driven automation** using Argo Events
- **Container orchestration workflows** with Argo Workflows
- **Service mesh routing** with Istio
- **In-cluster Git server** for complete isolation
- **Automated CI/CD pipeline** from code push to deployment

## 🎯 Learning Objectives

This repository is designed for learning the **Certified Argo Project Associate (CAPA)** certification topics:

- ✅ Declarative GitOps continuous delivery
- ✅ Kubernetes-native workflow automation
- ✅ Event-driven CI/CD pipelines
- ✅ Service mesh integration and traffic management
- ✅ Kustomize for environment-specific configurations
- ✅ Automated version management and deployment

## 🏗️ Architecture

### Components

```
        ┌────────────────────────────────────────────────────────────┐
        │                    Local K3d/Kind Cluster                  │
        │                                                            │
        │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
        │  │   ArgoCD     │  │ Argo Events  │  │Argo Workflows│      │
        │  │  (GitOps)    │  │  (Webhooks)  │  │   (CI/CD)    │      │
        │  └──────────────┘  └──────────────┘  └──────────────┘      │
        │         │                  │                  │            │
        │         │                  │                  │            │
        │  ┌──────▼──────────────────▼──────────────────▼─────────┐  │
        │  │              Istio Service Mesh                      │  │
        │  │         (Traffic Routing & Ingress Gateway)          │  │
        │  └──────────────────────────────────────────────────────┘  │
        │         │                                                  │
        │  ┌──────▼─────────┐         ┌────────────────┐             │
        │  │  Git Server    │         │  ReadingList   │             │
        │  │  (StatefulSet) │         │  Application   │             │
        │  └────────────────┘         └────────────────┘             │
        │                                                            │
        │  ┌──────────────────────────────────────────────────────┐  │
        │  │        Local Registry (registry.localhost:5000)      │  │
        │  └──────────────────────────────────────────────────────┘  │
        └────────────────────────────────────────────────────────────┘
```

### Repositories Structure

- **code-ecosystem** (this repo): Infrastructure, ArgoCD apps, and application source
- **cloud-components/**: Infrastructure components and Kubernetes manifests
- **cloud-deploy/**: Deployment configurations (simulates separate GitOps repo)
- **readinglist/**: Sample Go application source code

## 🚀 Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [k3d](https://k3d.io/) or [kind](https://kind.sigs.k8s.io/)
- [git](https://git-scm.com/)

### Setup

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd code-ecosystem
   ```

2. **Create the cluster:**
   ```bash
   cd cloud-components
   bash ./infra/scripts/init.sh create_cluster k3d
   ```

3. **Set kubeconfig:**
   ```bash
   export KUBECONFIG=$(pwd)/.kube/config-iws
   ```

4. **Add hostname to /etc/hosts:**
   ```bash
   echo "127.0.0.1 argocd.localhost readinglist-dev.localhost" | sudo tee -a /etc/hosts
   ```

5. **Get ArgoCD credentials:**
   ```bash
   bash ./infra/scripts/init.sh get_argocd_login
   ```

### Access Points

- **ArgoCD UI:** http://argocd.localhost:31120
  - Username: `admin`
  - Password: Check `infra/admin-pass.txt`

- **ReadingList App:** http://readinglist-dev.localhost:31120/api/v1/healthcheck

## 🔄 CI/CD Workflow

The complete automation pipeline:

```
  1. Developer pushes code to Git Server
          ↓
  2. post-receive hook triggers webhook
          ↓
  3. Argo Events EventSource receives webhook
          ↓
  4. Sensor triggers Argo Workflow
          ↓
  5. Workflow executes:
    - Clone readinglist repo
    - Build Docker image (version: vYY.MM-<commit-hash>)
    - Push to local registry
    - Clone cloud-deploy repo
    - Update Kustomize manifest with new tag
    - Commit and push manifest change
          ↓
  6. ArgoCD detects Git change
          ↓
  7. ArgoCD syncs new deployment
          ↓
  8. Application updated automatically
```

## 📁 Project Structure

```
code-ecosystem/
├── README.md
├── cloud-components/              # Infrastructure components
│   ├── applications/              # ArgoCD Application CRDs
│   │   ├── argo-events.yaml
│   │   ├── argo-workflow.yaml
│   │   ├── git-actions.yaml       # CI automation trigger
│   │   ├── git-server.yaml
│   │   ├── istio.yaml
│   │   └── readinglist-dev.yaml
│   ├── infra/
│   │   ├── scripts/init.sh        # Cluster setup automation
│   │   ├── k3d-cluster-config.yaml
│   │   └── kind-cluster-config.yaml
│   └── kustomize/                 # Base Kubernetes manifests
│       ├── argo-events/
│       ├── argo-workflow/
│       ├── argocd/
│       ├── git-server/
│       └── istio/
├── cloud-deploy/                  # Deployment configurations
│   └── kustomize/
│       ├── git-actions/           # Webhook EventSource & Sensor
│       │   └── base/
│       │       ├── webhook.yaml
│       │       └── special-workflow-trigger-shortened.yaml
│       └── readinglist/           # App deployment manifests
│           ├── base/              # Base resources
│           │   ├── deployment.yaml
│           │   ├── service.yaml
│           │   ├── configmap.yaml
│           │   ├── gateway.yaml
│           │   └── vs.yaml
│           └── dev/               # Dev environment overlay
│               └── kustomization.yaml
└── readinglist/                   # Go application
    ├── cmd/api/main.go
    ├── Dockerfile
    ├── go.mod
    └── scripts/build.sh
```

## 🛠️ Technologies Used

|        Component        |       Technology      | Purpose |
|-------------------------|-----------------------|------------------------------|
| **GitOps CD**           |        ArgoCD         | Continuous delivery and sync |
| **Workflow Engine**     |     Argo Workflows    | CI pipeline execution        |
| **Event Automation**    |      Argo Events      | Webhook-based triggers       |
| **Service Mesh**        |         Istio         | Traffic routing & ingress    |
| **Manifest Management** |       Kustomize       | Environment-specific configs |
| **Application**         |      Go (Golang)      | Sample REST API service      |
| **Container Runtime**   |        Docker         | Image builds and registry    |
| **Orchestration**       | Kubernetes (K3d/Kind) | Container orchestration      |
|-------------------------|-----------------------|------------------------------|
## 📋 Available Applications

### ReadingList
- **Type:** Go REST API
- **Endpoint:** `/api/v1/healthcheck`
- **Response:** Status, environment, version
- **Image:** `registry.localhost:5000/readinglist`
- **Access:** http://readinglist-dev.localhost:31120

### Git Server
- **Type:** In-cluster Git repository
- **Protocol:** SSH
- **Repositories:**
  - `code-ecosystem/readinglist.git`
  - `code-ecosystem/cloud-deploy.git`
- **Hook:** post-receive triggers CI/CD on push to main/master

## 🎓 Key GitOps Patterns Demonstrated

### 1. Declarative Configuration
All infrastructure and applications defined as Kubernetes manifests in Git.

### 2. Immutable Versioning
Automatic version tagging using date + commit hash (`vYY.MM-<hash>`).

### 3. Automation
Complete pipeline from code push to deployment without manual intervention.

### 4. Software Agents
ArgoCD continuously monitors Git and reconciles cluster state.

### 5. Closed Loop
Self-healing deployments with automated sync policies.

## 🧹 Cleanup

To delete the cluster and all resources:

```bash
cd cloud-components
bash ./infra/scripts/init.sh cleanup
```

## 📚 Resources

- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
- [Argo Workflows Documentation](https://argoproj.github.io/argo-workflows/)
- [Argo Events Documentation](https://argoproj.github.io/argo-events/)
- [Istio Documentation](https://istio.io/)
- [Kustomize Documentation](https://kustomize.io/)
- [K3d Documentation](https://k3d.io/)

## 📝 License

See [LICENSE](cloud-components/LICENSE) file for details.

---

**Note:** This is a learning environment designed for local development. For production deployments, additional security hardening, observability, and reliability features should be implemented.