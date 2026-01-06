# ArgoCD Deployment Guide

## Prerequisites
- OrbStack with Kubernetes cluster running
- ArgoCD installed on the cluster
- kubectl configured for your cluster

## Deployment Steps

### 1. Test Local Deployment (Optional)
```bash
./test-deploy.sh
```

### 2. Deploy with ArgoCD
```bash
./deploy.sh
```

### 3. Access the API
```bash
# Port forward to access the service
kubectl port-forward -n demo-api svc/demo-api-service 3000:80

# Test endpoints
curl http://localhost:3000/
curl http://localhost:3000/health
curl http://localhost:3000/docs
```

### 4. Monitor in ArgoCD
- Access ArgoCD UI
- Find the `demo-api` application
- Monitor sync status and health

### 5. Cleanup
```bash
./cleanup.sh
```

## Troubleshooting

### Check pod status
```bash
kubectl get pods -n demo-api
kubectl logs -n demo-api -l app=demo-api
```

### Check ArgoCD sync
```bash
kubectl get application demo-api -n argocd
```