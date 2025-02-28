# IAM Server

## Overview
This project is a scalable **Identity and Access Management (IAM) service** built with **Go** and deployed using **Docker and Kubernetes**. It provides authentication and authorization capabilities using **JWT tokens**, **Role-Based Access Control (RBAC)**, and **PostgreSQL for user management**.

## **Architecture**
The IAM service is designed using a **microservices-inspired modular architecture**, ensuring **scalability, security, and high availability**. The system consists of multiple components working together to provide authentication, authorization, and secure access management.

### **Architecture Diagram**
```mermaid
graph TD;
    A[API Gateway (Gorilla Mux)] -->|Routes Requests| B[Auth Service (JWT, Bcrypt)];
    A -->|Routes Requests| C[Policy Engine (RBAC)];
    B -->|Stores Users| D[PostgreSQL (User Storage)];
    B -->|Caches Sessions| E[Redis (Session Cache)];
    C -->|Authorization Checks| D;
```

### **Components**
#### **1️⃣ API Gateway & Routing**
- **Gorilla Mux** acts as the main entry point for all requests.
- Routes requests to **authentication** and **authorization** services.
- Handles **rate limiting, logging, and monitoring**.

#### **2️⃣ Authentication Service**
- Responsible for user **registration, login, and session management**.
- Uses **bcrypt** for password hashing.
- Issues **JWT tokens** for stateless authentication.

#### **3️⃣ Authorization Layer (RBAC)**
- Implements **Role-Based Access Control (RBAC)**.
- Uses **Casbin or Open Policy Agent (OPA)** for authorization.
- Checks permissions before allowing access to protected resources.

#### **4️⃣ Database Layer**
- **PostgreSQL** stores **user identities, roles, permissions, and logs**.
- **Redis** caches **user sessions** to improve performance.

#### **5️⃣ Deployment & Orchestration**
- **Docker & Kubernetes** provide a containerized, scalable environment.
- Kubernetes **liveness & readiness probes** ensure high availability.
- Horizontal scaling supported using **Kubernetes replicasets**.

## **Getting Started**

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/iam-server.git
cd iam-server
```

### 2️⃣ Build and Run with Docker Compose
```sh
docker compose up --build -d
```

### 3️⃣ Run Database Migrations
```sh
docker compose exec iam-server go run cmd/migrate/main.go
```

### 4️⃣ Test Health Check
```sh
curl -X GET http://localhost:8080/health
```

## **Deployment with Kubernetes**
### 1️⃣ Apply Kubernetes Manifests
```sh
kubectl apply -f deployments/k8s/
```

### 2️⃣ Verify Deployment
```sh
kubectl get pods
```

## **Environment Variables**
| Variable | Description |
|----------|------------|
| `PORT` | The port on which the IAM server runs (default `8080`) |
| `DATABASE_URL` | PostgreSQL connection string |
| `JWT_SECRET` | Secret key for signing JWT tokens |

## **Contributing**
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## **License**
This project is licensed under the MIT License - see the LICENSE file for details.

## **Contact**
For any questions or issues, please reach out to `chiram.littleton@gmail.com` or open an issue on GitHub.

