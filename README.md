The solution  
A golang webserver has been created and built into an image which is pushed to docker hub.  
Thhe image is deployed as a deployment behind a service. By default is set to be a single pod deployed, however the number can be increased.  
The golang server will connect to a redis instance behind a service, and will connect to it.  

The image has already been pushed to Docker hub. However, the steps to build the image were as follow
1. cd webimage
2. Build the image `docker build --tag jimbob687/webcounter:latest .`
3. Push image to registry `docker push jimbob687/webcounter:latest`

Steps to deploy the chart (which containes the golang webserver and redis) on Minikube 
1. cd webcounter  
2. Create namespace `kubectl create namespace webcounter`  
3. Install the helm chart `helm install webcounter . -n webcounter`  
4. Expose the node port `minikube service --url webcounter -n webcounter`  
5. Take the output url from step 4 and enter into a browser  
