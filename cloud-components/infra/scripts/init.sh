#!/bin/bash

export KUBECONFIG=.kube/config-iws

create_cluster(){
  echo "creating cluster...."
  kubectl cluster-info --context iws

  if [[ $? == 1 ]]; then
    if [[ $2 == "k3d" ]]; then
      echo "creating k3d cluster..."
      k3d cluster create --config infra/k3d-cluster-config.yaml
    elif [[ $2 == "kind" ]]; then
      echo "creating kind cluster..."
      kind create cluster --config infra/kind-cluster-config.yaml
    else
      echo "No valid cluster type provided, defaulting to kind..."
      k3d cluster create --config infra/k3d-cluster-config.yaml
    fi
  fi

  if [[ $? == 0 ]]; then
    echo "local cluster creation successful"
    kubectl create ns argocd
#    helm upgrade -f infra/values.yaml -n argocd argocd argo/argo-cd --install
    kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
    kubectl wait --for=create secret/argocd-initial-admin-secret -n argocd --timeout=180s
    kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d > infra/admin-pass.txt
    # disable tls
    kubectl patch cm -n argocd argocd-cmd-params-cm -p '{"data": {"server.insecure": "true"}}'
    kubectl rollout restart deployment -n argocd argocd-server
    kubectl rollout status deployment -n argocd argocd-server
    echo "Argocd UI login creds:"
    echo "---------------------------------------------"
    echo 
    echo "bash ./infra/scripts/init.sh get_argocd_login"
    echo
    echo "---------------------------------------------"
    echo
    install_app git-server
    install_app istio
    kubectl wait applications istio -n argocd --for=jsonpath='{.status.sync.status}=Synced' --timeout=180s
    kubectl wait applications istio -n argocd --for=jsonpath='{.status.health.status}=Healthy' --timeout=180s
    install_app argo-ingress
  else
    echo "Something went wrong please debug..."
    exit 1
  fi
  exit 0
}

cleanup(){
  echo "deleting cluster..."
  k3d cluster delete iws
  kind delete cluster --name iws
}

install_app(){
  kubectl apply -f applications/$1.yaml
}

install_all_apps(){
  kubectl apply -f applications/
}

get_argocd_login(){
    echo "username: admin"
    echo "password: $(cat infra/admin-pass.txt)"
}

"$@"