# 🚀 Go Peregerine Boilerplate

A production-ready **Go Fiber boilerplate** with built-in integrations for **AI services, Redis, background jobs, dependency injection (DI), structured logging, middleware, and clean architecture folder structure**.

This project is designed to help you bootstrap scalable backend services quickly with best practices in mind.

---

## ✨ Features

* ⚡ **Fiber Framework** – Fast HTTP web framework
* 🤖 **AI Integration** – Ready-to-use AI service layer
* 🧠 **Redis Support** – Caching & pub/sub support
* ⏱️ **Background Jobs** – Async task processing
* 🧩 **Dependency Injection (DI)** – Clean and testable architecture
* 🪵 **Structured Logger** – Centralized logging system
* 🛡️ **Middleware Ready** – Auth, logging, recovery, etc.
* 🗂️ **Clean Folder Structure** – Scalable and maintainable
* 🌱 **Environment Config** – Easy configuration management

---

## 📁 Project Structure

```
├── app/                 # Application entry points
│   ├── controllers      # Application logic
│   ├── jobs             # Background process
│   ├── libraries        # Main libraries
│   ├── middlwares       # Middleware for controller and route
│   ├── requests         # application request for controller
├── configs/             # App configuration (env, setup)
├── DI/                  # Dependency Injection
├── routes/              # Route definitions
├── systems/             # Helper scripts
│   ├── conf/            # Main app configuration
│   ├── services/        # Main services
│   ├── types/           # Main Types
│   ├── middleware/      # Custom middleware
│   ├── ai/              # AI integration layer
├── .env                 # Environment variables
├── go.mod
├── Dockerfile
├── makefile
├── compose.yml
└── main.go
```

---

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/NovalHendriyanto13/go-peregerine.git
cd go-peregerine
```


### 2. Install dependencies

```bash
go mod tidy
```
### 3. Setup environment

Copy `.env.example`:

```bash
cp .env.example .env
```

Edit the config as needed.

---

## ▶️ Running the App

```bash
go run main.go
```

Server will start at:

```
http://localhost:8000
```
---

## 🔧 Configuration

Example `.env`:

```
APP_NAME=go-fiber-boilerplate
APP_ENV=development
PORT=3000

REDIS_HOST=localhost
REDIS_PORT=6379

AI_API_KEY=your_api_key
```

---

## 🤖 AI Integration

AI services are abstracted inside:

```
systems/ai/
```

You can:

* Swap providers easily
* Centralize prompt logic
* Reuse AI clients across services

---

## 🧠 Redis Usage
Supports:

* Caching
* Distributed locking
* Pub/Sub messaging

---

## ⏱️ Background Jobs

Jobs are handled in:

```
jobs/
```

Features:

* Async processing
* Retry mechanism (optional)
* Queue-based execution

---

## 🧩 Dependency Injection (DI)

All dependencies are initialized in:

```
DI/
```

Benefits:

* Loose coupling
* Easier testing
* Better scalability

---

## 🪵 Logger

Structured logging available via:

```
DI/
```

Features:

* JSON logging
* Request tracing
* Error tracking

---

## 🛡️ Middleware

Custom middleware examples:

* Request logging
* Authentication
* Panic recovery
* Rate limiting (optional)

Located in:

```
systems/middleware/
```

---

## 📌 Example Route

```go
app.Get("/health", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status": "ok",
    })
})
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📦 Build

```bash
go build -o app 
```

---

## 🐳 Docker (Optional)

```bash
make dev
```

---

## 🚀 Roadmap

* [ ] Add authentication module (JWT/OAuth)
* [ ] Add database ORM integration
* [ ] Add API documentation (Swagger)

---

## 🤝 Contributing

Contributions are welcome!

1. Fork the repo
2. Create a new branch
3. Commit your changes
4. Open a pull request

---

## 📄 License

MIT License

---

## 💡 Notes

This boilerplate is designed for:

* Scalable backend systems
* Microservices architecture
* AI-powered applications

Feel free to customize based on your project needs.

---

Happy coding 🚀
