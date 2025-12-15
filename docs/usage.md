# Usage Guide

## 🔧 Configuration

### Environment Variables
The Go Log Generator uses an `.env` file located in `go-log-local/.env`. You can modify this file to change where logs are written inside the container (though you shouldn't need to for standard usage).

### Promtail Config
Located at `promtail/promtail.yaml`.
- **Scrape Configs**: Defines how Promtail finds log files.
- **Labels**: Important for querying in Grafana. Each job is tagged.
    - `job: frontend` matches `/var/log/frontend/*.log`
    - `job: backend` matches `/var/log/backend/*.log`
    - etc.

## 📊 Querying Logs in Grafana

1. **Log in to Grafana**: [http://localhost:3000](http://localhost:3000)
2. **Go to Explore**: Click the compass icon in the left sidebar ("Explore").
3. **Select Datasource**: Ensure **Loki** is selected in the top-left dropdown.

### Basic LogQL Queries

To view logs for a specific service, use the **Log Browser** or type the query directly:

**Filter by Job (Service):**
```logql
{job="backend"}
```

**Search for Errors:**
```logql
{job="backend"} |= "ERROR"
```

**Count Errors over time (Rate):**
```logql
rate({job="database"} |= "ERROR" [1m])
```

**JSON Parsing (if logs are JSON):**
Currently, logs are in syslog-like format. If you switch to JSON format in the Go app, you can use:
```logql
{job="auth"} | json
```

## 🐞 Troubleshooting

- **No Logs in Grafana?**
    - Check if the containers are running: `docker-compose ps`
    - Check Promtail logs: `docker logs promtail`
    - Verify log files are being created:
      ```bash
      ls -R logs/
      ```
    - Ensure time synchronization between containers (Loki is sensitive to timestamps).

- **Permission Issues**:
    - If Promtail cannot read logs, ensure file permissions allow reading by the Promtail user. The Docker Compose setup mounts `./logs` which maps to the container's log directory.
