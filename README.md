# Local setup with Minikube

```sh
  minikube start
  minikube addons enable ingress
  kubectl apply -f test-site.yaml
  kubectl apply -f ingress.yaml
```

Add the ip from `minikube service test-site --url` to the `/etc/hosts` file

kubernetes-test-site should be running at `mollers.se`
