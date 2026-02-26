# Kubernetes Deployment Guide

## Prerequisites

- Kubernetes cluster (1.24+)
- `kubectl` configured
- Docker registry access
- (Optional) Ingress controller (nginx)
- (Optional) cert-manager for TLS

## Build and Push Docker Image

```bash
# Build the image
docker build -t your-registry/taskmanager:v1.0.0 -f k8s/Dockerfile .

# Push to registry
docker push your-registry/taskmanager:v1.0.0
```

## Deploy to Kubernetes

### Option 1: Deploy All Resources

```bash
# Update image in deployment.yaml first
kubectl apply -f k8s/
```

### Option 2: Deploy Step-by-Step

```bash
# 1. Create ConfigMap
kubectl apply -f k8s/configmap.yaml

# 2. Create Deployment
kubectl apply -f k8s/deployment.yaml

# 3. Create Service
kubectl apply -f k8s/service.yaml

# 4. (Optional) Create Ingress
kubectl apply -f k8s/ingress.yaml

# 5. (Optional) Create HPA
kubectl apply -f k8s/hpa.yaml
```

## Verify Deployment

```bash
# Check pods
kubectl get pods -l app=taskmanager

# Check service
kubectl get svc taskmanager-service

# Check deployment
kubectl get deployment taskmanager-app

# View logs
kubectl logs -l app=taskmanager --tail=100 -f

# Describe pod (for troubleshooting)
kubectl describe pod <pod-name>
```

## Access the Application

### Via Port Forward (Local Testing)

```bash
kubectl port-forward svc/taskmanager-service 8080:80
# Access at http://localhost:8080
```

### Via Ingress (Production)

Update `k8s/ingress.yaml` with your domain and apply:

```bash
kubectl apply -f k8s/ingress.yaml
# Access at https://taskmanager.example.com
```

## Scaling

### Manual Scaling

```bash
kubectl scale deployment taskmanager-app --replicas=5
```

### Auto-Scaling (HPA)

```bash
kubectl apply -f k8s/hpa.yaml
kubectl get hpa taskmanager-hpa --watch
```

## Update Deployment

```bash
# Update image
kubectl set image deployment/taskmanager-app taskmanager=your-registry/taskmanager:v1.1.0

# Or edit deployment
kubectl edit deployment taskmanager-app

# Check rollout status
kubectl rollout status deployment/taskmanager-app

# Rollback if needed
kubectl rollout undo deployment/taskmanager-app
```

## Cleanup

```bash
# Delete all resources
kubectl delete -f k8s/

# Or delete individually
kubectl delete deployment taskmanager-app
kubectl delete service taskmanager-service
kubectl delete ingress taskmanager-ingress
kubectl delete hpa taskmanager-hpa
kubectl delete configmap taskmanager-config
```

## Resource Specifications

| Resource | CPU Request | CPU Limit | Memory Request | Memory Limit |
|---|---|---|---|---|
| Per Pod | 250m | 500m | 512Mi | 1Gi |

## Notes

- **Health Checks**: The deployment includes liveness and readiness probes. For this CLI app, you'll need to add Spring Boot Actuator or a simple HTTP endpoint.
- **Image**: Update `deployment.yaml` line 23 with your actual image registry and tag.
- **Domain**: Update `ingress.yaml` line 10 with your actual domain.
- **TLS**: Requires cert-manager installed in the cluster for automatic certificate provisioning.
