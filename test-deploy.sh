#!/bin/bash

set -e

echo "🧪 Testing Kubernetes manifests locally"

# Build Docker image
echo "📦 Building Docker image..."
docker build -t demo-api:latest .

# Load image to OrbStack
echo "📤 Loading image to OrbStack..."
docker save demo-api:latest | docker load

# Apply manifests directly
echo "🚀 Applying manifests..."
kubectl apply -k k8s/overlays/dev

# Wait for deployment
echo "⏳ Waiting for pods to be ready..."
kubectl wait --for=condition=ready pod -l app=demo-api -n demo-api --timeout=60s

# Show status
echo "✅ Deployment status:"
kubectl get pods,svc -n demo-api

echo "🎉 Test deployment complete!"
echo "📍 Test the API:"
echo "   kubectl port-forward -n demo-api svc/demo-api-service 3000:80"