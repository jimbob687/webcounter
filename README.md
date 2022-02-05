1. Build the image `docker build --tag jimbob687/webcounter:latest .`

2. Push image to registry `docker push jimbob687/webcounter:latest`

3. Create namespace `kubectl create namespace webcounter`

4. Install the helm chart `helm install webcounter . -n webcounter`

5. Expose node port `minikube service --url webcounter -n webcounter`  
