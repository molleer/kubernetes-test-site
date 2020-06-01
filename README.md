# Local setup with Minikube

```sh
  minikube start
  minikube addons enable ingress
  kubectl apply -f test-site.yaml
  kubectl apply -f ingress.yaml
```

Add the ip from `minikube service test-site --url` to the `/etc/hosts` file

kubernetes-test-site should be running at `mollers.se`

# Installing Traefik

```sh
  helm repo add traefik https://containous.github.io/traefik-helm-chart
  helm repo update
  helm install --values traefik-values.yaml traefik traefik/traefik
```

Traefik will work as you default ingress controller
