
# Akash Monitoring System

This repository contains the source code for a monitoring system designed to manage and monitor a network of nodes using a decentralized architecture. Nodes in the system register themselves in a global routing table, are organized into token rings, and communicate via a gossip protocol for data synchronization and health checks.  
This application is designed to integrate seamlessly with any of your other apps deployed on Akash. Simply add it as a binary, open required port and specify the IP address of any deployed node. It will continuously monitor the health and reachability of all nodes, and promptly notify you if any node becomes unavailable

## Features

- **Node Registration**: Nodes register their IP address and launch time to a global routing table.
- **Token Rings**: Nodes are divided into groups by token rings based on a robust hashing algorithm.
- **Health Checks**: Nodes perform health checks within their token ring and also to random nodes outside their token ring.
- **Failure Detection**: A consensus mechanism ensures that if a node fails health checks, corrective action is taken.
- **Notifications**: Owners are notified via email if their node is detected as unhealthy.

## Project Structure

- `/cmd` - Contains the main application entry point.
- `/pkg`
  - `/config` - Configuration management.
  - `/node` - Node definition and operations.
  - `/router` - HTTP routing for API endpoints.
  - `/store` - Data storage and synchronization logic.
  - `/health` - Health checking, failure detection and notification mechanisms.
- `/configs` - Example configuration files.

## Setup

### Prerequisites

- Go 1.15 or higher
- Access to an SMTP server for sending email notifications (optional)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/monitoring-system.git
   cd monitoring-system
   ```

2. Build the project:
   ```bash
   go build -o monitoring-system ./cmd/main
   ```

3. Configure the application:
   - Edit the configuration file in `/configs/config.json` or set the necessary environment variables (NODE_IP, NODE_EMAIL, etc.).

### Running the Application

To start the application, either use the configuration file:
```bash
./onePunchAkashMonitoring /path/to/configs.json
```
or ensure the environment variables are set and run:
```bash
./onePunchAkashMonitoring
```

## Usage

- **Health Check Endpoint**: `GET /health`
- **Update Endpoint**: `POST /update`

These endpoints can be accessed to perform health checks and send updates or commands to nodes respectively.
