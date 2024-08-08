# DevOpsify the go web application

The main goal of this project is to implement DevOps practices in the Go web application. The project is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

DevOps practices include the following:

- Creating Dockerfile (Multi-stage build)
- Containerization
- Continuous Integration (CI)
- Continuous Deployment (CD)

## Summary Diagram
![image](https://github.com/user-attachments/assets/45f4ef12-c5b5-4247-9d43-356b5dfb671b)


## Creating Dockerfile (Multi-stage build)

The Dockerfile is used to build a Docker image. The Docker image contains the Go web application and its dependencies. The Docker image is then used to create a Docker container.

We will use a Multi-stage build to create the Docker image. The Multi-stage build is a feature of Docker that allows you to use multiple build stages in a single Dockerfile. This will reduce the size of the final Docker image and also secure the image by removing unnecessary files and packages.

## Containerization

Containerization is the process of packaging an application and its dependencies into a container. The container is then run on a container platform such as Docker. Containerization allows you to run the application in a consistent environment, regardless of the underlying infrastructure.

We will use Docker to containerize the Go web application. Docker is a container platform that allows you to build, ship, and run containers.

Commands to build the Docker container:

```bash
docker build -t <your-docker-username>/go-web-app .
```

Command to run the Docker container:

```bash
docker run -p 8080:8080 <your-docker-username>/go-web-app
```

Command to push the Docker container to Docker Hub:

```bash
docker push <your-docker-username>/go-web-app
```

## Continuous Integration (CI)

Continuous Integration (CI) is the practice of automating the integration of code changes into a shared repository. CI helps to catch bugs early in the development process and ensures that the code is always in a deployable state.

We will use GitHub Actions to implement CI for the Go web application. GitHub Actions is a feature of GitHub that allows you to automate workflows, such as building, testing, and deploying code.

The GitHub Actions workflow will run the following steps:

- Checkout the code from the repository
- Build the Docker image
- Run the Docker container
- Run tests

## Continuous Deployment (CD)

Continuous Deployment (CD) is the practice of automatically deploying code changes to a production environment. CD helps to reduce the time between code changes and deployment, allowing you to deliver new features and fixes to users faster.

We will use Argo CD to implement CD for the Go web application. Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes. It allows you to deploy applications to Kubernetes clusters using Git as the source of truth.

The Argo CD application will deploy the Go web application to a Kubernetes cluster. The application will be automatically synced with the Git repository, ensuring that the application is always up to date.

## What tools and tech are we implementing in this project?


	1. Containerization: Docker
	2. Kubernetes:  including the creation of EKS cluster, deployment, service and ingress
	3. CI: GitHub Action (Done)
	4. CD: Argo CD 
	5. Helm Chart 
	6. Ingress Controller:  LB -> Exposed
	                                   |
                                   	  DNS



## Installation Instructions

### 1. Download and Install Go Version 1.22.5

```bash
wget https://golang.org/dl/go1.22.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```

### 2. Add Go to PATH

```bash
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile
```

### 3. Verify the Installation

```bash
go version
```

### 4. Remove the Current Version of Go (if needed)

```bash
sudo apt-get remove golang-go
sudo apt-get remove --auto-remove golang-go
sudo rm -rf /usr/local/go
```




# Go Lang Application Deployment on EKS using Helm and Argo CD

This guide outlines the steps to build, containerize, and deploy a Go Lang application on an EKS cluster using Helm and Argo CD.

## Prerequisites

- Go Lang installed
- Docker installed
- Kubernetes cluster (EKS)
- Helm installed
- Argo CD installed and configured
- Private Docker image repository
- GitHub account for CI/CD

## Steps

### 1. Build the Go Lang Application

```bash
go build -o main .
```

### 2. Execute the Go Binary

```bash
./main
```

### 3. Test the Application Locally

Test the application successfully on your local machine to ensure it works as expected.

### 4. Create a Multi-Stage Dockerfile

Create a multi-stage Dockerfile to containerize your application.

### 5. Build Docker Image

Build a Docker image from the Dockerfile and test the application via containerization.

```bash
docker build -t your-image-name .
```

### 6. Push Docker Image to Private Repository

Push the Docker image to your private Docker image repository.

```bash
docker push your-image-name
```

### 7. Create Kubernetes Manifests

Create Kubernetes manifests for deployment, service, and ingress.

### 8. Apply Kubernetes Manifests

Apply all the manifest files using `kubectl`.

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```

### 9. Configure Ingress

Create an ingress resource but note that the address is not defined yet. An ingress controller will assign an IP address. After assigning the IP address, map it to your domain name `go-web-app.local` in your `/etc/hosts` file.

### 10. Install Helm

Install Helm on your local machine if it's not already installed.

```bash
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
```

### 11. Create a Helm Chart

Create a Helm chart for your application.

```bash
helm create go-web-app
```

### 12. Copy Manifests to Helm Chart

Copy all your Kubernetes manifests files into the Helm chart `templates` directory.

### 13. Install Helm Chart

Install the Helm chart, which will deploy all your resources at once.

```bash
helm install go-web-app ./go-web-app
```

### 14. CI Stage: GitHub Actions

Set up GitHub Actions for CI with the following steps:
1. Build and Unit Testing
2. Static Code Analysis
3. Docker Image Creation and Push to Docker Hub
4. Update Helm with the newly created Docker image

### 15. CD Stage: Argo CD (GitOps)

Set up Argo CD for Continuous Deployment.

#### Install and Setup Argo CD

Follow the official Argo CD installation guide to install and configure Argo CD in the gitops folder.

#### Access Argo CD via NodePort Mode

Get the Argo CD secret:

```bash
kubectl get secret -n argocd
kubectl edit secret argocd-initial-admin-secret -n argocd
```

Decode the secret to get the password:

```bash
echo 'your base64 encode password' | base64 --decode
```

Login to Argo CD with the username `admin` and the decoded password.

### 16. Deploy with Argo CD

Pull the Helm chart and deploy it into the EKS cluster using Argo CD.

## Conclusion

Congratulations!!! You have now successfully built, containerized, and deployed your Go Lang application on an EKS cluster using Helm and Argo CD. The CI/CD pipeline ensures your application is continuously integrated and deployed with ease.




