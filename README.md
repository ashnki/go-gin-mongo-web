

---

```md
# Go + Gin + MongoDB Web Application (Dockerized)

This project is a full-stack web application built using **Go**, **Gin**, and **MongoDB**, with a server-side rendered **HTML webpage**.  
The entire stack is containerized using **Docker** and **Docker Compose**.

The goal of this project is to practice:
- Reading unfamiliar backend repositories
- Serving real webpages from backend services
- Connecting Go applications to MongoDB
- Debugging containerized applications
- Verifying live database data

---

## Project Structure

```

go-gin-mongo-web/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── .env
├── README.md
├── main.go
└── templates/
└── index.html

```

---

## Tech Stack

- Go (Gin Framework)
- MongoDB
- HTML (server-side templates)
- Docker
- Docker Compose

---

## Application Flow

1. Go application starts from `main.go`
2. Gin web server listens on port `8080`
3. HTML page is rendered using Go templates
4. User submits data via form
5. Data is stored in MongoDB
6. Stored data is fetched and displayed on the webpage

---

## Environment Variables

Create a `.env` file in the project root:

```

MONGO_URI=mongodb://db:27017

````

> Note: `db` is the MongoDB service name in Docker Compose.

---

## Running the Application

### Using Docker Compose

Build and start the application:

```bash
docker-compose up --build
````

---

## Accessing the Web Application

### Local Machine

```
http://localhost:8080
```

### Killercoda

1. Expose port **8080** in the Killercoda UI
2. Use the generated public URL

---

## Viewing Database Data

### Enter MongoDB Container

```bash
docker exec -it <mongo_container_name> mongosh
```

### Check Data

```js
use appdb
show collections
db.items.find().pretty()
```

---

## Live Data Monitoring (Change Streams)

To watch data arriving in real time:

```js
db.items.watch()
```

> Requires MongoDB replica set configuration.

---

## Common Issues

### 502 Bad Gateway

* App container is not running
* Check logs:

```bash
docker logs <app_container_name>
```

### Go Build Failure

* Go does not allow unused imports
* Remove unused imports from `main.go`

---

## Useful Commands

```bash
docker-compose up
docker-compose up --build
docker-compose down
docker ps
docker logs
```

---

## Key Learnings

* Go applications fail fast and require explicit error handling
* Docker Compose service names act as DNS hostnames
* MongoDB data persists using Docker volumes
* Real webpages can be served directly from backend services
* Container logs are the first place to debug issues

---


