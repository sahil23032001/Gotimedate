🚀 Product Manager Technical Assignment – Sahil Jain
Welcome to the submission of my Product Manager technical assignment. This repository contains the implementation and deliverables for all three problem statements:

📌 Problem Statement 1: Product Requirements & Wireframes for Container Security
🎯 Objective
Build a security product feature that scans container images for known vulnerabilities and allows users to assess and take action based on severity.

✅ Deliverables
Product Requirements Document:
📝 View PRD Document

Low-Fidelity Wireframes:
Included in the PRD document.

Bonus – Dev Action Items:

Integrate with vulnerability scanning tools (e.g., Trivy, Clair).

Build an API to fetch image scan data.

Create filters for severity, fix status, etc.

Build an alerting engine for high/critical CVEs.

🛡️ Problem Statement 2: Kubernetes Security Scan
🔍 Objective
Scan a local Kubernetes cluster using a security tool like Kubescape and share the vulnerability report.

✅ Deliverables
Result.json contains findings from the K8s security scan.
📄 View JSON Output

⏱️ Problem Statement 3: Deploying Go App via Docker & Kubernetes
🔧 Step-by-Step Implementation
1️⃣ GoLang App – Date & Time
A simple Go program that shows the current date and time.

2️⃣ Dockerize the App
Dockerfile is included to build the container image.

3️⃣ Push Image to Docker Hub
Docker image built and used locally via tar (due to private registry constraints).

4️⃣ Kubernetes Deployment
Used k3s for lightweight local deployment.

Deployed using declarative YAML with 2 replicas.

Exposed via Kubernetes Service.

✅ Files Included
main.go: Go source code.

Dockerfile: For building the image.

deployment.yaml: K8s Deployment spec.

service.yaml: K8s Service spec.

datetimeapp.tar: Prebuilt image archive for local loading.

🌐 Sample Output


