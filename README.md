# Infokes Take-Home Test

## Project Overview
This project is a folder and file explorer application with a Vue.js frontend and a Go backend.

## Prerequisites
- Docker
- Docker Compose
- npm
- Go
- MySQL

## Getting Started

### Running the Application

#### Using Docker Compose

1. Clone the repository:
```bash
git clone https://github.com/yourusername/infokes-take-home-test.git
cd infokes-take-home-test
```

2. Build and start the application:
```bash
docker-compose up --build
```

3. Access the application:
- Frontend: `http://localhost`
- Backend API: `http://localhost:8080`

### Stopping the Application
```bash
docker-compose down
```

## Development

### Frontend (Vue.js)
- Located in `./frontend`
- Uses Vite
- Runs on port 80

### Backend (Go)
- Located in `./backend`
- Uses Gin framework
- Runs on port 8080

## Running Tests
- Frontend tests: `npm run test`
- Backend tests: `go test ./...`

## Technologies
- Frontend: Vue.js, TypeScript
- Backend: Go, Gin
- Database: MySQL
- Containerization: Docker, Docker Compose
