cat << 'EOF' > README.md
# LifeScience

Welcome to **LifeScience**, a *Golang-based REST API* designed for logging and retrieving life sciences video annotations with geospatial data integration—this project simulates a data engineering workflow, tracking events like video recordings of biological experiments, enriched with location data from **OpenStreetMap (OSM)** and stored in a **PostgreSQL database**, all containerized with **Docker** for easy deployment, showcasing *concurrency*, *database interactions*, and *external API usage*.

## Project Overview

*LifeScience* serves as a proof-of-concept for a life sciences data engineering pipeline, developed over **23 commits** from *March 6, 2022, to March 6, 2025*, demonstrating incremental feature growth like PostgreSQL persistence, OSM geospatial lookups, and Docker support, making it perfect for learning about RESTful services, database integration, and containerization in a scientific context, with ✨ **key features** including REST API endpoints (`POST /api/v1/events` to create events with metadata and `GET /api/v1/events` to retrieve them by timestamp), structured storage with PostgreSQL and `sqlx`, OSM integration with retry logic, asynchronous operations via goroutines, Dockerized API and database services, and dependency management via Go modules.

## 🚀 Prerequisites

To get started, ensure you have:  
**Go** version 1.21+ (check with `go version`),  
**Docker** with Compose (verify with `docker --version` and `docker-compose --version`),  
**Git** for repo management (run `git --version`),  
and **internet access** for OSM API calls and Docker image pulls.

## 📦 Installation

Installation begins with cloning the project to your local machine using `git clone https://github.com/davidgin/LifeScience.git` followed by `cd LifeScience`, then:

### Step 2: Build and Run with Docker

Set up and run the project using Docker—build the application with `docker-compose build` (*builds the Go app image using the `Dockerfile`*),  
start all services with `docker-compose up` (*runs the API at `http://localhost:8080` and PostgreSQL at `localhost:5432`, use `-d` for background mode*),  
and stop services with `docker-compose down` (*stops containers; add `--volumes` to reset PostgreSQL data*),

### Step 3: Run Locally Without Docker

or run natively by installing Go dependencies with `go mod download` (*downloads `gin-gonic/gin`, `jmoiron/sqlx`, and `lib/pq`*),  
setting up PostgreSQL by installing it (e.g., macOS: `brew install postgresql`, Ubuntu: `apt install postgresql`),  
starting the service with `pg_ctl start` or `systemctl start postgresql`,  
creating a user and database with `psql -c "CREATE USER life_user WITH PASSWORD 'life_pass123';"` and `psql -c "CREATE DATABASE life_science_db OWNER life_user;"`,  
and launching the app with `go run main.go` (*starts the API at `http://localhost:8080`, stop with `Ctrl+C`*).

## 🎯 Usage

Interact with the API—create a new event with `curl -X POST http://localhost:8080/api/v1/events -H "Content-Type: application/json" -d '{"title": "Frog Dissection", "description": "Lab experiment video", "video_url": "https://example.com/frog.mp4"}'` (*submits an event; the API assigns an `id`, `timestamp`, and fetches a `location` from OSM, returning a sample response like `{"id": 1, "title": "Frog Dissection", "description": "Lab experiment video", "timestamp": "2025-03-06T12:00:00Z", "processed": false, "location": "Unknown Location", "video_url": "https://example.com/frog.mp4"}`*),  
retrieve all events with `curl http://localhost:8080/api/v1/events` (*returns events in descending timestamp order, e.g., `[{"id": 1, "title": "Frog Dissection", "description": "Lab experiment video", "timestamp": "2025-03-06T12:00:00Z", "processed": false, "location": "Unknown Location", "video_url": "https://example.com/frog.mp4"}]`*),  
and handle errors where invalid JSON yields `400 Bad Request` with `{"error": "parsing error"}` and database issues yield `500 Internal Server Error` with `{"error": "db error"}`.

## 🗂️ Project Structure

The project structure is organized as `LifeScience/` containing:  
`main.go` (🚀 *entry point: sets up Gin router and server*),  
`models.go` (📋 *defines `Event` struct with JSON/DB tags*),  
`handlers.go` (🎮 *API handlers: `CreateEvent` and `GetEvents` with goroutines*),  
`db.go` (🗄️ *PostgreSQL setup: connection and schema via `sqlx`*),  
`osm.go` (🌍 *OSM integration: client and retry logic for location fetching*),  
`Dockerfile` (🐳 *multi-stage build: Go 1.21 to Alpine*),  
`docker-compose.yml` (⚙️ *configures API and PostgreSQL services*),  
`go.mod` (📦 *lists dependencies: `gin`, `sqlx`, `pq`*),  
and `go.sum` (🔒 *dependency checksums for integrity*).

## 🤝 Contributing

Contributing to *LifeScience* is encouraged—fork the repo using the "Fork" button on GitHub or `git clone https://github.com/<your-username>/LifeScience.git`,  
create a branch with `git checkout -b feature/your-feature-name` (*e.g., `feature/add-authentication`*),  
make changes and test locally with `go run main.go` or `docker-compose up`,  
commit changes with `git add .` and `git commit -m "Implement your feature with details"`,  
push to your fork with `git push origin feature/your-feature-name`,  
and submit a pull request by visiting `https://github.com/davidgin/LifeScience`, clicking "Pull Requests," then "New Pull Request,"  
following guidelines to use consistent coding style (run `go fmt`), add tests if feasible (future goal), and provide clear PR descriptions.

## 📜 License

This project is under the **MIT License** (to be added as `LICENSE` file).

## 📬 Contact

For questions or support, contact **[davidgin](https://github.com/davidgin)** via GitHub Issues—**Happy coding with LifeScience!** 🎉
EOF
