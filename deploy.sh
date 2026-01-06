#!/bin/bash

set -e

echo "🚀 Deploying Demo API to Kubernetes with ArgoCD"

# Build and tag Docker image
echo "📦 Building Docker image..."
docker build -t demo-api:latest .

# Load image into OrbStack (if using local registry)
echo "📤 Loading image to OrbStack..."
docker save demo-api:latest | docker load

# Apply ArgoCD application
echo "🔄 Creating ArgoCD application..."
kubectl apply -f argocd/application.yaml

# Wait for sync
echo "⏳ Waiting for ArgoCD to sync..."
sleep 10

# Check deployment status
echo "✅ Checking deployment status..."
kubectl get pods -n demo-api

echo "🎉 Deployment complete!"
echo "📍 Access the API:"
echo "   kubectl port-forward -n demo-api svc/demo-api-service 3000:80"
echo "   Then visit: http://localhost:3000"