# NK Celestin - Go Web Application

A modern, professional web application built with Go, showcasing technology courses and educational resources. This application demonstrates clean architecture, responsive design, and modern web development practices.

## 🌟 Brand Identity

**NK Celestin** - Your trusted partner in technology education and professional development.

### Color Palette
- **Primary**: #2563eb (Blue) - Headers, footers, and primary elements
- **Secondary**: #7c3aed (Purple) - Accents and highlights  
- **Accent**: #06b6d4 (Cyan) - Links and interactive elements

## 🚀 Features

- Clean, modern UI with professional branding
- Responsive design that works on all devices
- Multiple pages: Home, About, Courses, Contact
- Built with Go's native `net/http` package
- Containerized with Docker for easy deployment
- Kubernetes-ready with Helm charts
- CI/CD pipeline with GitHub Actions

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Package Manager**: Helm
- **CI/CD**: GitHub Actions

## 🏃 Running the Application

### Local Development

To run the server locally, execute:

```bash
go run main.go
```

The server will start on port 8080. Access the application at:
- Home: `http://localhost:8080/home`
- About: `http://localhost:8080/about`
- Courses: `http://localhost:8080/courses`
- Contact: `http://localhost:8080/contact`

### Docker Deployment

```bash
# Build the Docker image
docker build -t nk-celestin-web-app .

# Run the container
docker run -p 8080:8080 nk-celestin-web-app
```

### Kubernetes Deployment

```bash
# Using kubectl
kubectl apply -f k8s/manifests/

# Using Helm
helm install nk-celestin-app helm/go-web-app-chart/
```

## 📁 Project Structure

```
.
├── static/              # Static HTML files and assets
│   ├── home.html
│   ├── about.html
│   ├── courses.html
│   ├── contact.html
│   └── images/
├── helm/                # Helm charts for Kubernetes
├── k8s/                 # Kubernetes manifests
├── main.go              # Main application entry point
├── main_test.go         # Unit tests
├── Dockerfile           # Container configuration
└── README.md            # This file
```

## 📚 Documentation

For detailed DevOps and deployment documentation, see:
- [DevOps Guide](README-DevOps.md)
- [EKS Setup](eks/)
- [GitOps with ArgoCD](gitops/argocd/)
- [Ingress Controller](ingress-controller/nginx/)

## ✅ Testing

```bash
# Run unit tests
go test -v

# Run with coverage
go test -cover
```

## 🚀 CI/CD

This project includes automated CI/CD pipeline using GitHub Actions. The pipeline:
- Runs tests on every push
- Builds Docker images
- Deploys to Kubernetes clusters

See `.github/workflows/cicd.yaml` for configuration details.

## 📝 License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## ©️ Copyright

Copyright © 2026 NK Celestin. All rights reserved.

---

**NK Celestin** - Empowering through technology education