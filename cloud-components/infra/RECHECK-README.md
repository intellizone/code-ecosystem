# 🚀 Create a Kind Cluster with Preinstalled Istio and ArgoCD

This project helps you set up a local Kubernetes cluster using [Kind](https://kind.sigs.k8s.io/) with **Istio** and **ArgoCD** preinstalled. It's a great environment for testing GitOps workflows and service meshes locally.

---

## 📋 Prerequisites

Ensure you have the following installed on your machine:

- [Docker](https://docs.docker.com/engine/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [git](https://git-scm.com/)

---

## 📂 Clone the Repository

```bash

git clone https://github.com/intellizone/Certified-Argo-Project-Associate-CAPA.git

cd Certified-Argo-Project-Associate-CAPA
```

## ⚙️ Create the Cluster

Run the following script to create a Kind cluster with the name capa-lfs256. You can modify the script for custom configurations.

```sh

bash ./infra/scripts/init.sh create_cluster

# Set kubeconfig to access the new cluster
export KUBECONFIG=$(pwd)/.kube/config-capa-lfs256

# Add ArgoCD hostname to /etc/hosts for local access
echo "127.0.0.1	argocd.localhost" | sudo tee -a /etc/hosts

```

## 🔐 Access ArgoCD UI

Once the cluster is up and running, access ArgoCD at 
### 👉 http://argocd.localhost:31120/

### 🎫 Login Credentials

**Username**: admin

**Password**: Stored in infra/admin-pass.txt

To retrieve the login password:

```sh

bash ./infra/scripts/init.sh get_argocd_login

```

## 🧱 Application Deployment Structure (ArgoCD + Kustomize + Istio)

This repository demonstrates a GitOps-friendly structure using **ArgoCD**, **Kustomize**, and **Istio**.

You can use it as a **reference** for structuring and deploying your own applications.

### 📂 Folder Overview
```plaintext
.
├── applications/ # ArgoCD Application CRs (what to deploy)
└── kustomize/ # App manifests and Kustomize overlays (how to deploy)
```

---

## 🚀 Deploying Your Own App

### 1. Add App Resources

Create a new folder under `kustomize/` for your app:

```plaintext
kustomize/
    └──my-app/
          ├── deployment.yaml
          ├── service.yaml
          ├── virtualservice.yaml
          ├── gateway.yaml
          └── kustomization.yaml
```

Make sure your `VirtualService` includes:

```yaml
spec:
  hosts:
    - my-app.localhost
# And your Service points to the correct ports.
```
### 2. Define an ArgoCD Application
Create a file: applications/my-app.yaml

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/<username>/Certified-Argo-Project-Associate-CAPA.git
    path: kustomize/my-app
    targetRevision: HEAD
  destination:
    server: https://kubernetes.default.svc
    namespace: my-app
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```
Then apply it:

```bash
bash ./infra/scripts/init.sh install_app my-app
```
### 🌐 Accessing Your App via DNS
After deployment, your app will be accessible at:

```cpp
http://<app-name>.localhost:31120
```
### 📌 Example:
If your app name is my-app, visit:

```cpp
http://my-app.localhost:31120
```
#### This works because:

Istio's IngressGateway listens on port 31120

Traffic is routed based on the hostname (Host: my-app.localhost)

Your VirtualService and Gateway are configured accordingly


## 🧹 Cleanup

To delete the cluster and clean up resources:

```sh

bash ./infra/scripts/init.sh cleanup

```

## 📝 Notes

The scripts and configuration files are located in the infra/ directory.

Make sure Docker is running before creating the cluster.

The setup includes Istio and ArgoCD preconfigured for quick experimentation.

## 📧 Support

For issues or questions, feel free to open an issue on the GitHub repository.


---

### How to Save:

1. Open a terminal.
2. Run: `nano README.md` (or use your preferred editor).
3. Paste the above content.
4. Save and exit.

Let me know if you want to add diagrams, badges, a license, or contributor info.