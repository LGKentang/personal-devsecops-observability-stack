# coffee-stack Helm chart

This chart deploys the coffee microservices (`coffee-catalog`, `coffee-orders`, `gateway`) plus a Postgres database and a one-shot migration Job.

Quick install (for a cluster where you can pull `lgkentang/*` images or after you push your images):

```bash
# from repository root
helm install my-coffee services/k8s/helm/coffee-stack
```

Apply migrations (the chart includes a Job that will run migrations on install if `migrations.enabled` is true):

```bash
# optionally check job status
kubectl get jobs
kubectl logs job/apply-migrations
```

Secrets / credentials
- The chart no longer stores a plaintext Postgres password in `values.yaml`.
- For local/dev use you can create a `values.secret.yaml` (see `values.secret.example.yaml`) and install with:

```bash
helm install my-coffee services/k8s/helm/coffee-stack -f services/k8s/helm/coffee-stack/values.secret.yaml
```

- Or create a Kubernetes Secret directly and enable `postgres.createSecret`:

```bash
kubectl create secret generic postgres-secret \
	--from-literal=username=postgres \
	--from-literal=password='REPLACE_ME' \
	--from-literal=database=coffee -n coffee

# then install with createSecret enabled
helm install my-coffee services/k8s/helm/coffee-stack --set postgres.createSecret=true
```

Do NOT commit the secret values file into git; add it to `.gitignore`.

Customization: edit `values.yaml` to set image repositories/tags. Provide secrets via `values.secret.yaml` or native Kubernetes Secrets.
