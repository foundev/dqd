# DQD Quick Start Guide

Get up and running with the new DQD architecture in 5 minutes.

## Prerequisites

- **Go** 1.21+ ([install](https://go.dev/doc/install))
- **Node.js** 18+ ([install](https://nodejs.org/))
- **Git** (to clone the repo)

## Option 1: Run New Stack (Go + React)

### 1. Start Go Backend

```bash
cd backend-go
go mod download
go run ./cmd/server/main.go
```

Server starts on **http://localhost:8080**

### 2. Start React Frontend (New Terminal)

```bash
cd frontend
npm install
npm run dev
```

Frontend starts on **http://localhost:3000**

### 3. Open Browser

Visit **http://localhost:3000**

You should see the new DQD interface!

## Option 2: Run Java Backend (Legacy)

If you need the full functionality while Go backend is being developed:

```bash
./mvnw clean package
java -jar target/dqd-0.12.3.jar server -p 8080
```

Then configure the React frontend to use Java backend:

```bash
cd frontend
echo "VITE_API_URL=http://localhost:8080" > .env
npm run dev
```

## Option 3: Run Both Backends

Run both Java and Go backends simultaneously:

**Terminal 1 - Java Backend:**
```bash
./mvnw clean package
java -jar target/dqd-0.12.3.jar server -p 9090
```

**Terminal 2 - Go Backend:**
```bash
cd backend-go
PORT=8080 go run ./cmd/server/main.go
```

**Terminal 3 - React Frontend:**
```bash
cd frontend
npm install
npm run dev
```

The frontend will use Go backend (port 8080) by default, but you can switch to Java (port 9090) by changing `VITE_API_URL`.

## Testing the API

### Test Go Backend

```bash
# Get version
curl http://localhost:8080/api/about.json

# Health check
curl http://localhost:8080/health
```

### Test Java Backend

```bash
# Get version
curl http://localhost:9090/about.json

# Health check (if implemented)
curl http://localhost:9090/
```

## What's Implemented?

### ✅ Working Now (Go Backend)

- Version endpoint (`/api/about.json`)
- Health check endpoint
- CORS support
- Basic routing

### ✅ Working Now (React Frontend)

- Modern UI with Tailwind CSS
- Navigation and routing
- File upload component
- Simple profile analysis page (UI only)

### 🔄 Coming Soon

- Profile JSON analysis
- Queries JSON analysis
- Profile comparison
- Schema generation
- IOStat analysis
- Thread top analysis

## Common Issues

### Port Already in Use

If port 8080 is taken:

```bash
# Go backend
PORT=8081 go run ./cmd/server/main.go

# Update frontend proxy in vite.config.ts
```

### Module Not Found (Go)

```bash
cd backend-go
go mod tidy
```

### Module Not Found (Node)

```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
```

## Project Structure

```
dqd/
├── backend-go/          # New Go REST API
│   ├── cmd/server/      # Main entry point
│   └── internal/        # Go packages
├── frontend/            # New React app
│   └── src/             # React components
└── src/                 # Original Java code
```

## Development Workflow

1. **Make changes** to Go backend or React frontend
2. **Hot reload** happens automatically (Vite HMR)
3. **Test** via browser or curl
4. **Commit** when feature is working

## Next Steps

- Read [MIGRATION.md](./MIGRATION.md) for detailed migration plan
- Read [backend-go/README.md](./backend-go/README.md) for Go backend details
- Read [frontend/README.md](./frontend/README.md) for React frontend details
- Check [GitHub Issues](https://github.com/rsvihladremio/dqd/issues) for tasks

## Need Help?

- Check the README files in each directory
- Open an issue on GitHub
- Ask in #dqd-migration Slack channel

## Building for Production

### Build Go Backend

```bash
cd backend-go
make build
./bin/dqd-server
```

### Build React Frontend

```bash
cd frontend
npm run build
# Output in frontend/dist/
```

Serve with any static file server or integrate into Go binary.

---

**Happy coding! 🚀**
