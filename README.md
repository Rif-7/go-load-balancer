# Simple HTTP Load Balancer in Go

This project implements a lightweight HTTP load balancer in Go using only the standard library. It supports round-robin request distribution, health checks, and reverse proxying to backend servers.

## Features

- Round-robin load balancing across multiple backend servers
- TCP-based health checks for backend availability
- Reverse proxying to forward client requests
- Concurrency-safe design

## Getting Started

### Prerequisites

- Go installed on your system

### Installation

1. Clone the repository

2. Start backend servers:

   In separate terminals or background processes, run:

   ```bash
   go run backend/backend.go -port 8081
   go run backend/backend.go -port 8082
   go run backend/backend.go -port 8083
   ```

3. Start the load balancer:

   ```bash
   go run loadbalancer.go
   ```

   The load balancer will listen on port 8080 by default.

4. Test the load balancer:

   ```bash
   curl http://localhost:8080/hello
   ```

   You should see responses from different backend ports in a round-robin fashion.

## Project Structure

- `loadbalancer.go` – Load balancer logic
- `backend/backend.go` – Simple backend server that returns its port in the response

## Notes

- Backend availability is checked periodically using TCP connections.
- If a backend is unreachable, it is skipped until it becomes reachable again.
