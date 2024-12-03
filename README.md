# INFO 7255 Big Data Indexing

---

## Overview
This project implements a RESTful API for managing and indexing structured JSON objects. It supports CRUD operations, JSON schema validation, and caching with ETag headers. Indexed data is searchable through Elasticsearch and visualized using Kibana.

---

## Tech Stack
- **Programming Language:** Go (Gin-Gonic framework)
- **Cache:** Redis
- **Search Engine:** Elasticsearch
- **Message Queue:** RabbitMQ

---

## Features
- **OAuth 2.0 Authentication:** Secure API access using Google Cloud Platform OAuth2.0.
- **JSON Validation:** Validate incoming requests against a defined JSON schema.
- **Response Caching:** Cache server responses and validate cache consistency using ETag.
- **Comprehensive API Methods:** Support for `POST`, `PUT`, `PATCH`, `GET`, and `DELETE` HTTP methods.
- **Data Persistence:** Store JSON objects in Redis as key-value pairs.
- **Search Capabilities:** Index JSON objects in Elasticsearch for advanced search.
- **Message Queue Integration:** Use RabbitMQ to manage indexing requests asynchronously.

---

## Data Flow
1. **Authentication:** Generate an OAuth token through the authorization workflow.
2. **Validation:** Validate API requests using the received ID token.
3. **Create Data:** Create a JSON object using the `POST` method.
4. **JSON Schema Validation:** Validate incoming JSON objects against the defined schema.
5. **Redis Storage:** Deconstruct hierarchical JSON objects into key-value pairs for storage in Redis.
6. **Queue Indexing Requests:** Add objects to a RabbitMQ queue for indexing.
7. **Index Data:** Consume messages from RabbitMQ and index data into Elasticsearch.
8. **Search:** Use Kibana to query and retrieve indexed data.

---

## API Endpoints
### Plan Management
- **POST** `/v1/plan`: Create a new plan from the request body.
- **PUT** `/v1/plan/{id}`: Update an existing plan by ID (requires a valid ETag in `If-Match` header).
- **PATCH** `/v1/plan/{id}`: Partially update an existing plan by ID (requires a valid ETag in `If-Match` header).
- **GET** `/v1/plan/{id}`: Retrieve a plan by ID. Optionally provide an ETag in `If-None-Match` header for caching.
- **DELETE** `/v1/plan/{id}`: Delete a plan by ID (requires a valid ETag in `If-Match` header).

---

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```
3. Install [`redis-stack-server`](https://redis.io/docs/latest/operate/oss_and_stack/install/install-stack/) and start the service:
   ```bash
   redis-stack-server
   ```
4. Use Docker Compose to deploy dependencies, including Elasticsearch, Kibana, and RabbitMQ:
   ```bash
   docker compose up -d
   ```
5. Run the Go application:
   ```bash
   go run main.go
   ```
6. Start the RabbitMQ consumer:
   ```bash
   go run consumer/main.go
   ```

---

## Dependencies and URLs
- **Elasticsearch:** [http://localhost:9200](http://localhost:9200)
- **Kibana:** [http://localhost:5610](http://localhost:5610)
- **RabbitMQ:** [http://localhost:15672](http://localhost:15672)
  - **Default RabbitMQ Credentials:**
    - Username: `guest`
    - Password: `guest`
