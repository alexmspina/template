# install nginx ingress

## Standard install for generic clouds
### Step 1
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml

### Step 2
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml

## For KIND clusters
### Step 3 - run as a NodePort service to work with socat
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/baremetal/service-nodeport.yaml

### Step 4 - run nginx socat script to expose ingress NodePort service
./scripts/k8s/socat/socat_nginx.fish