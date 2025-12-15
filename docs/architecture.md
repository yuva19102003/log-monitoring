# Architecture Overview

This document describes the architectural components of the log monitoring setup.

## High-Level Diagram

```mermaid
graph LR
    subgraph "Application Layer"
        GoApp[Go Log Generator] -->|Writes Logs| LogFiles[Log Files (/var/log/*)]
    end

    subgraph "Log Shipping"
        Promtail[Promtail] -->|Reads| LogFiles
        Promtail -->|Pushes Stream| Loki
    end

    subgraph "Storage & Query"
        Loki[Loki]
    end

    subgraph "Visualization"
        Grafana[Grafana] -->|Queries| Loki
    end
```

## Component Details

### 1. Go Log Generator
- **Location**: `go-log-local/`
- **Function**: A custom Golang application designed to simulate a microservices environment.
- **Behavior**: It writes simulated log entries to various files in `/var/log/` corresponding to different services:
    - `frontend`
    - `backend`
    - `database`
    - `auth`
    - `system`
- **Configuration**: Uses an `.env` file to determine log file paths.

### 2. Promtail
- **Role**: Log Collector & Shipper.
- **Configuration**: `promtail/promtail.yaml`
- **Mechanism**:
    - Mounts the host's log directory (which is shared with the Go app via Docker volumes).
    - Discovers log files using `static_configs`.
    - Tags logs with labels (e.g., `job=backend`, `job=frontend`) before sending them to Loki.

### 3. Loki
- **Role**: Log Aggregation System.
- **Configuration**: Default local config (mounted via `docker-compose`).
- **Function**: Receives log streams from Promtail, indexes the labels, and stores the log chunks. It does *not* index the text content of the logs, making it highly efficient.

### 4. Grafana
- **Role**: Visualization Dashboard.
- **Integration**: Pre-configured to use Loki as a datasource.
- **Features**: Allows users to query logs using LogQL (Loki Query Language) and visualize log volume and content.
