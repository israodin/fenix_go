### Build Docker Image

docker build -t israodin/fenix:v0.1.0-alpine  -f Dockerfile.alpine .

### Push Docker image to repository

docker push israodin/fenix:v0.1.0-alpine

### Forward ports from kubernetese to outside
kubectl port-forward svc/fenix 6767 -n fenix  

### Create secrets from text file
kubectl create secret generic db-creds --from-file=./credantials/username.txt --from-file=./credantials/password.txt -n fenix

### How to  get secrets to  output
kubectl  get secrets db-creds -o yaml -n fenix

