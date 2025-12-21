# Cloud Components - GitOps Infrastructure

Infrastructure components and configurations for the CODE Ecosystem project, designed for learning **Certified Argo Project Associate (CAPA)** concepts.

---

## 🚀 Quick Start

To set up a local Kubernetes cluster (K3d or Kind) with Istio and ArgoCD, follow the step-by-step guide:

👉 **[Setup Guide → infra/README.md](infra/README.md)**

This includes:

- K3d/Kind cluster creation
- Istio & ArgoCD installation
- Git server deployment
- Accessing the ArgoCD UI
- Cleanup steps

---

## 📚 Resources

- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
- [Argo Workflows Documentation](https://argoproj.github.io/argo-workflows/)
- [Argo Events Documentation](https://argoproj.github.io/argo-events/)
- [Argo Rollouts Documentation](https://argo-rollouts.readthedocs.io/)
- [Istio Documentation](https://istio.io/)
- [K3d Documentation](https://k3d.io/)
- [Kind Documentation](https://kind.sigs.k8s.io/)
- [Kustomize Documentation](https://kustomize.io/)

---

# 📖 Learning Guide: Argo Projects & GitOps

This section covers the theory and concepts needed for the CAPA certification.

---

## What Is GitOps?
- Backbone of modern Software delivery practice

### Five key aspects:
- Declarative configuration
- Immutable storage
- Automation
- Software agents
- Closed loop

## What Is Argo?
- Kubernetes-native tools with workflow management capabilities of Kubernetes.
- includes 
    - Argo Continuous Delivery (CD)
    - Argo Workflows for running complex jobs
    - Argo Events for event-based dependency management 
    - Argo Rollouts for progressive delivery

### Argo Continuous Delivery (CD)
- declarative, GitOps Continuous Delivery tool for Kubernetes

features:
- blue-green and canary update strategies
- integrates with service meshes
- ingress controllers to shape traffic
- automates promotion
- rollback based on analysis
- safely deploy artifacts into Production

### Argo Workflows
- Kubernetes-native workflow engine for orchestrating parallel jobs
- Container-native workflow engine (similar to Jenkins but Kubernetes-native)
- Supports DAGs (Directed Acyclic Graphs) and step-based workflows
- Used for CI/CD, batch processing, and machine learning pipelines

### Argo Events
- Event-driven workflow automation framework
- Triggers Kubernetes objects, Argo Workflows, serverless workloads, and other processes in response to events
- Supports 20+ event sources: webhooks, S3, schedules, messaging queues, GCP PubSub, and more
- Allows customization of business-level constraint logic for workflow automation

**Key Components:**
- **Event Sources**: Define where events come from (webhooks, Git, queues, etc.)
- **Sensors**: Define what happens when events are received (triggers)
- **EventBus**: Message bus for event communication
- **Triggers**: Actions to execute (create workflows, K8s resources, etc.)

### Argo Rollouts
- Progressive delivery controller for Kubernetes
- Extends Kubernetes Deployment capabilities with advanced deployment strategies
- Native Kubernetes doesn't support canary/blue-green deployments out of the box

**Features:**
- Blue-green and canary update strategies
- Integrates with service meshes and ingress controllers to shape traffic
- Automated promotion and rollback based on analysis
- Integration with metric providers (Prometheus, Datadog, New Relic, etc.)
- Analysis runs for automated deployment validation

Quiz:

2.1 What is Argo’s role in Kubernetes workflow management?

a. Argo is a database management tool \
b. Argo provides version-controlled application deployment and workflow automation \
c. Argo is a containerization platform \
d. Argo is a networking solution for Kubernetes 

Ans: b

2.2 What is the primary role of Kubernetes in the context of Argo’s functionalities?

a. Kubernetes is a version control system \
b. Kubernetes provides essential infrastructure for container orchestration \
c. Kubernetes is a continuous delivery tool \
d. Kubernetes is a database management system 

Ans: b

2.3 How does Argo align with modern development practices, ensuring a single source of truth for both code and infrastructure?

a. Utilizing Docker Compose files \
b. Enforcing Microservices architecture \
c. Removing outdated practices from Kubernetes \
d. Following the principles of GitOps

Ans: D

## Argo CD
- declarative, GitOps continuous delivery tool

### Advancements
- GitOps
- Continuous Delivery
- Rollbacks
- Multi-environment management
- UI and API

### Vocabulary
- Configuration
    - Application
    - Application source type: helm/kustomize
- States
    - Target state
    - Live state
- Statuses
    - Sync status
    - Sync operation status
    - Health status 
- Actions
    - Sync
    - Refresh

### Core Components

**Reference:** https://argo-cd.readthedocs.io/en/stable/operator-manual/architecture/

<img src="https://argo-cd.readthedocs.io/en/stable/assets/argocd_architecture.png" 
     onerror="this.onerror=null; this.src='static/argocd_architecture.png';" 
     alt="argocd_architecture" />

**Main Components:**
- **API Server**: Exposes the API consumed by Web UI, CLI, and CI/CD systems
- **Repository Server**: Internal service that maintains a local cache of Git repositories
- **Application Controller**: Kubernetes controller that continuously monitors running applications and compares current state against desired state

![argocd_architecture_simple](static/argocd_architecture_simple.png)


### Argo CD reconciliation loop
![Argo CD reconciliation loop](static/argo_cd_reconciliation_loop.png)

### Synchronization Principles

**Resource Hooks** (using `kind: Job`):
- `PreSync`: Runs before the sync operation
- `Sync`: Runs during the sync operation
- `PostSync`: Runs after all resources are synced
- `Skip`: Skips the resource during sync
- `SyncFail`: Runs when sync operation fails

**Reference:** https://argo-cd.readthedocs.io/en/stable/user-guide/resource_hooks/

**Sync Waves:**
- Controls the order of resource synchronization
- Each wave is delayed by 2 seconds (default)
- Modify delay with: `ARGOCD_SYNC_WAVE_DELAY` environment variable
- Lower numbered waves are synced first
### Simplifying Application Management
- Application
  ```yaml
  apiVersion: argoproj.io/v1alpha1
  kind: Application
  metadata:
    name: guestbook
    namespace: argocd
  spec:
    project: default
    source:
      repoURL: 'htt‌ps://github.com/argoproj/argocd-example-apps.git'
      targetRevision: HEAD
      path: guestbook
  destination:
    server: 'htt‌ps://kubernetes.default.svc'
    namespace: guestbook
  ```

- AppProject

  ```yaml
  apiVersion: argoproj.io/v1alpha1
  kind: AppProject
  metadata:
    name: production
    namespace: argocd
  spec:
    description: Production applications
    sourceRepos:
      - '*'
    destinations:
      - namespace: production
        server: 'htt‌ps://kubernetes.default.svc'
    clusterResourceWhitelist:
      - group: '*'
        kind: '*'
  ```

- Repository credentials

  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: private-repo-creds
    namespace: argocd
    labels:
      argocd.argoproj.io/secret-type: repository
  stringData:
    url: 'htt‌ps://github.com/private/repo.git'
    username: <username>
    password: <password>
  ```

- Cluster credentials

  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: external-cluster-creds
    labels:
      argocd.argoproj.io/secret-type: cluster
  type: Opaque
  stringData:
    name: external-cluster
    server: 'ht‌tps://external-cluster-api.com'
    config: |
      {
        "bearerToken": "<token>",
        "tlsClientConfig": {
          "insecure": false,
          "caData": "<certificate encoded in base64>"
        }
      }
  ```

### Argo CD Extensions & Integrations
- Plugins --> we can manage it with configmap

  ```yaml
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: argocd-notifications-cm
  data:
    context: |
      region: east
      environmentName: staging

    template.a-slack-template-with-context: |
      message: "Something happened in {{ .context.environmentName }} in the {{ .context.region }} data center!"
  ```

### Securing Argo CD
- Use RBAC
- Manage Secrets Securely
- Regularly Update Argo CD



---

## Argo Workflows Deep Dive

### Workflow Structure

A Workflow spec has two core parts:

- Entrypoint: Specifies the name of the template that serves as the entrypoint for the workflow. It defines the starting point of the workflow execution.
- Templates: A template represents a step or task in the workflow that should be executed. There are six types of templates that we will introduce next.

### Template Types
- Template Definitions
  - ***Container***: A Container is the most common template type and represents a step in the workflow that runs a container. It is suitable for executing containerized applications or scripts. Example:

    ```yaml
    - name: whalesay
      container:
        image: docker/whalesay
        command: [cowsay]
        args: ["hello world"]
    ```

  - ***Resource***: A Resource represents a template for creating, modifying, or deleting Kubernetes resources. It is useful for performing operations on Kubernetes objects. Example:

    ```yaml
      - name: k8s-owner-reference
        resource:
          action: create
          manifest: |
            apiVersion: v1
            kind: ConfigMap
            metadata:
              generateName: owned-eg-
            data:
              some: value
    ```

  - ***Script***: A Script is similar to the container template but allows specifying the script inline without referencing an external container image. It can be used for simple scripts or one-liners. Example:

    ```yaml
    - name: gen-random-int
      script:
        image: python:alpine3.6
        command: [python]
        source: |
          import random
          i = random.randint(1, 100)
          print(i)
    ```

  - ***Suspend***: A Suspend is a template that suspends execution, either for a duration or until it is resumed manually. It can be resumed using the CLI, the API endpoint, or the UI. Example:

    ```yaml
    - name: delay
      suspend:
        duration: "20s"
    ```

- Template Invocators
  - ***DAG***: Defining task as a Dependency Graph, used for complex dependencies and conditional execution

    ```yaml
    - name: diamond
      dag:
        tasks:
        - name: A
          template: echo
        - name: B
          dependencies: [A]
          template: echo
        - name: C
          dependencies: [A]
          template: echo
        - name: D
          dependencies: [B, C]
          template: echo
    ```

  - ***Steps***: Steps are defining multiple steps within a template as several steps need to be executed sequentially or in parallel.

    ```yaml
    - name: hello-hello-hello
      steps:
      - - name: step1
          template: prepare-data
      - - name: step2a
          template: run-data-first-half
        - name: step2b
          template: run-data-second-half
    ```


  ### Outputs
  - Defining Outputs
  - Accessing Outputs

  ```yaml
  apiVersion: argoproj.io/v1alpha1
  kind: Workflow
  metadata:
    generateName: artifact-passing-
  spec:
    entrypoint: artifact-example
    templates:
    - name: artifact-example
      steps:
      - - name: generate-artifact
          template: whalesay
      - - name: consume-artifact
          template: print-message
          arguments:
            artifacts:
            - name: message
              from: "{{steps.generate-artifact.outputs.artifacts.hello-art}}"

    - name: whalesay
      container:
        image: docker/whalesay:latest
        command: [sh, -c]
        args: ["cowsay hello world | tee /tmp/hello_world.txt"]
      outputs:
        artifacts:
      - name: hello-art
        path: /tmp/hello_world.txt

    - name: print-message
      inputs:
        artifacts:
        - name: message
          path: /tmp/message
      container:
        image: alpine:latest
        command: [sh, -c]
        args: ["cat /tmp/message"]
  ```
### WorkflowTemplate
  ```yaml
  apiVersion: argoproj.io/v1alpha1
  kind: WorkflowTemplate
  metadata:
    name: sample-template
  spec:
    templates:
    - name: hello-world
      inputs:
        parameters:
          - name: msg
            value: "Hello World!"
      container:
        image: docker/whalesay
        command: [cowsay]
        args: ["{{inputs.parameters.msg}}"]
  ```

  ```yaml
  apiVersion: argoproj.io/v1alpha1
  kind: Workflow
  metadata:
    generateName: hello-world-
  spec:
  entrypoint: whalesay
  templates:
    - name: whalesay
      steps:
        - - name: hello-world
            templateRef:
              name: sample-template
              template: hello-world
  ```
### Argo Workflows Architecture

**Components:**
- **Argo Server**: REST API server (similar to Kubernetes API server)
  - Provides API interface for workflow submission, monitoring, and management
  - Serves the Web UI
  - Handles authentication and authorization
- **Workflow Controller**: Kubernetes controller that orchestrates workflow execution
  - Watches for Workflow CRDs
  - Manages workflow state transitions
  - Schedules and monitors pods
- **Argo UI**: Web-based interface for visualizing and managing workflows

---

## Argo Rollouts Deep Dive

### Continuous Integration
- Frequent code commits
- Automated tests
- Immediate problem detection
- Reduced integration problems

### Continuous Delivery
- Automated release process
- Reliable deployments
- Rapid release cycles
- Close collaboration between teams

### Progressive Delivery
- **Canary releases**: Gradually roll out the change to a small subset of users before rolling it out to the entire user base.
- **Feature flags**: Control who gets to see what feature in the application, allowing for selective and targeted deployment.
- **Experiments & A/B testing**: Test different versions of a feature with different segments of the user base.
- **Phased rollouts**: Slowly roll out features to incrementally larger segments of the user base, monitoring and adjusting based on feedback.

### Deployment Strategies
- Recreate/fixed deployment
- Rolling update
- Blue-green deployment
- Canary deployment

![alt text](static/Strategies_for_Smooth_and_Reliable_Releases.png)

### Building Blocks of Argo Rollouts
![alt text](static/Building_Blocks_of_Argo_Rollouts.png)

### Argo Rollouts Components

**Core Resources:**
- **Rollouts Controller**: Watches for Rollout resources and manages progressive deployments
- **Rollout Resource**: Custom resource that extends Deployment capabilities
  - [Specification Reference](https://argo-rollouts.readthedocs.io/en/stable/features/specification/)
- **Ingress/Service**: Standard Kubernetes resources for traffic routing
  - [Ingress Docs](https://kubernetes.io/docs/concepts/services-networking/ingress/)
  - [Service Docs](https://kubernetes.io/docs/concepts/services-networking/service/)
- **ReplicaSet**: Manages different versions during rollout
  - [ReplicaSet Docs](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)
- **AnalysisTemplate & AnalysisRun**: Define and execute deployment validation metrics
- **Metric Providers**: External systems for analysis (Prometheus, Datadog, etc.)

### Argo Rollouts Functionalities
- Blue-green deployments
- Canary deployments
- Advanced traffic routing
- Integration with metric providers
- Automated decision making
---

## Argo Events Deep Dive

### Event-Driven Architecture

![Event-Driven Architecture](static/Event-Driven_Architecture.png)

**Architecture Components:**
- **Event Source**: Defines the source of events (webhooks, Git, S3, queues, etc.)
  - Creates an EventSource custom resource
  - Listens for events and publishes them to EventBus
- **EventBus**: Message transport layer (NATS-based by default)
  - Provides pub/sub messaging between EventSources and Sensors
- **Sensor**: Listens to events and triggers actions
  - Defines dependencies (which events to listen for)
  - Defines triggers (what to do when events match)
- **Trigger**: Action to execute when event conditions are met
  - Can create Kubernetes resources, Argo Workflows, HTTP requests, etc.

### Implementation in This Project

In the CODE Ecosystem, Argo Events orchestrates the CI/CD pipeline:

1. **EventSource** (`webhook.yaml`): Listens on port 12000 for Git push events
2. **Sensor** (`special-workflow-trigger-shortened.yaml`): Triggers build workflow when events received
3. **Workflow**: Builds Docker image, updates manifests, pushes to Git
4. **ArgoCD**: Detects manifest changes and syncs deployment

## 🧾 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

## 🙌 Contributing

Contributions are welcome! Please open issues or submit pull requests if you find bugs or have suggestions.

---