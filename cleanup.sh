#!/bin/bash

echo "🧹 Cleaning up demo-api deployment"

# Delete ArgoCD application
echo "🗑️ Deleting ArgoCD application..."
kubectl delete -f argocd/application.yaml --ignore-not-found=true

# Delete namespace and resources
echo "🗑️ Deleting namespace and resources..."
kubectl delete namespace demo-api --ignore-not-found=true

echo "✅ Cleanup complete!"