# Log Monitoring Setup

This repository contains a complete log monitoring stack using **Golang**, **Promtail**, **Loki**, and **Grafana** (PLG Stack). It demonstrates how to generate, ship, store, and visualize logs from a microservices-style application.

## 🚀 Quick Start

### Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running the Stack

1. **Clone the repository** (if you haven't already):
   ```bash
   git clone <repository-url>
   cd log-monitoring
   ```

2. **Start the services**:
   ```bash
   docker-compose up -d
   ```
   This will start the following containers:
   - `log-generator`: A Go application that simulates log generation.
   - `loki`: The log aggregation system.
   - `promtail`: The log shipper.
   - `grafana`: The visualization dashboard.

3. **Access Grafana**:
   Open your browser and navigate to [http://localhost:3000](http://localhost:3000).
   - **Default Credentials**: `admin` / `admin` (You may be asked to change the password).

## 📂 Documentation

- [Architecture Overview](docs/architecture.md): Understand the components and data flow.
- [Usage Guide](docs/usage.md): Detailed instructions on configuration and querying logs.

## 🛠️ Components

| Component | Role | URL / Port |
|-----------|------|------------|
| **Go Log Generator** | Simulates application logs (Frontend, Backend, DB, etc.) | N/A |
| **Promtail** | Ships logs from local files to Loki | N/A |
| **Loki** | Stores and indexes log streams | `http://localhost:3100` |
| **Grafana** | Visualizes logs and metrics | `http://localhost:3000` |

## 🧹 Cleanup

To stop and remove all containers, networks, and volumes:
```bash
docker-compose down -v
```
