# QuikBid - Online Auction Platform

A scalable microservices-based online auction platform built with Go, featuring real-time bidding with WebSockets.

## 🏗️ Architecture

### Services
- **User Service** (Port 8081): Handles authentication, user registration, and user profiles
- **Auction Service** (Port 8082): Manages auction items, creation, and lifecycle *(Simplified in this implementation)*
- **Bid Service** (Port 8083): Handles bidding logic and real-time notifications via WebSockets

### Technology Stack
- **Language**: Go 1.21+
- **Web Framework**: Fiber v2
- **WebSockets**: Gorilla WebSocket
- **Databases**: 
  - PostgreSQL (User Service - Port 5432)
  - PostgreSQL (Bid Service - Port 5433)
  - MongoDB (Auction Service - Port 27017)
- **Inter-Service Communication**: gRPC
- **Authentication**: JWT tokens
- **ORM**: GORM (PostgreSQL)
- **Containerization**: Docker & Docker Compose

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL
- MongoDB

### 1. Clone and Setup
```bash
git clone <repository-url>
cd QuikBid
```

### 2. Start Infrastructure Services
```bash
docker-compose -f deployments/docker-compose.yml up -d postgres postgres_bids mongodb
```

### 3. Generate Protocol Buffers
```bash
# Install protobuf tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate proto files
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/*.proto

# Move generated files to correct directories
mkdir -p proto/user proto/auction proto/bid
mv proto/user*.go proto/user/
mv proto/auction*.go proto/auction/
mv proto/bid*.go proto/bid/
```

### 4. Run Services

#### Terminal 1 - User Service
```bash
cd user-service
go mod tidy
go run cmd/main.go
```

#### Terminal 2 - Bid Service  
```bash
cd bid-service
go mod tidy
go run cmd/main.go
```

### 5. Test the APIs
```bash
chmod +x test-api.sh
./test-api.sh
```

## 📡 API Endpoints

### User Service (Port 8081)
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - User authentication
- `GET /api/v1/users/profile` - Get current user profile (protected)
- `GET /api/v1/users/{id}` - Get user by ID (protected)
- `GET /health` - Health check

### Bid Service (Port 8083)
- `POST /api/v1/bids/{auction_id}` - Place bid (protected)
- `GET /api/v1/bids/{auction_id}` - Get bid history
- `GET /api/v1/bids/{auction_id}/highest` - Get highest bid
- `GET /api/v1/bids/user/me` - Get current user's bids (protected)
- `WS ws://localhost:8084/ws/bids/{auction_id}` - Real-time bid updates
- `GET /health` - Health check

## 📝 API Usage Examples

### 1. Register a User
```bash
curl -X POST http://localhost:8081/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

### 2. Login and Get Token
```bash
curl -X POST http://localhost:8081/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### 3. Place a Bid (with authentication)
```bash
TOKEN="your-jwt-token-here"
curl -X POST http://localhost:8083/api/v1/bids/auction-123 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "amount": 150.00
  }'
```

### 4. WebSocket Connection (JavaScript)
```javascript
const ws = new WebSocket('ws://localhost:8084/ws/bids/auction-123');

ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    console.log('Received:', data);
    
    if (data.type === 'new_bid') {
        console.log(`New bid: $${data.data.amount} by ${data.data.bidder_id}`);
    } else if (data.type === 'auction_end') {
        console.log(`Auction ended! Winner: ${data.data.winner_id}`);
    }
};
```

## 🏗️ Project Structure

```
QuikBid/
├── user-service/              # User management & authentication
│   ├── cmd/main.go           # Entry point
│   ├── internal/
│   │   ├── api/              # HTTP & gRPC handlers
│   │   ├── auth/             # JWT utilities
│   │   ├── repository/       # Database access
│   │   └── service/          # Business logic
│   └── pkg/                  # Shared packages
├── bid-service/              # Bidding & real-time updates
│   ├── cmd/main.go           # Entry point
│   ├── internal/
│   │   ├── api/              # HTTP handlers
│   │   ├── repository/       # Database access
│   │   ├── service/          # Business logic
│   │   └── ws/               # WebSocket hub
│   └── pkg/                  # Shared packages
├── auction-service/          # Auction management (skeleton)
├── proto/                    # gRPC protocol definitions
├── deployments/              # Docker & K8s configurations
└── test-api.sh              # API testing script
```

## 🔧 Development

### Running Tests
```bash
# Run tests for specific service
cd user-service && go test ./...
cd bid-service && go test ./...
```

### Building
```bash
# Build specific service
cd user-service && go build -o bin/user-service cmd/main.go
cd bid-service && go build -o bin/bid-service cmd/main.go
```

### Docker Build
```bash
# Build and run with Docker Compose
docker-compose -f deployments/docker-compose.yml up --build
```

## 🌟 Features Implemented

### ✅ Core Features
- [x] User registration and authentication with JWT
- [x] RESTful API with Fiber framework
- [x] Real-time bidding with WebSockets
- [x] PostgreSQL integration with GORM
- [x] gRPC inter-service communication
- [x] Microservices architecture
- [x] Docker containerization
- [x] API testing script

### ✅ Technical Features
- [x] JWT middleware for authentication
- [x] WebSocket hub for real-time updates
- [x] Database migrations
- [x] Error handling and logging
- [x] CORS support
- [x] Health check endpoints

## 🔮 Future Enhancements

### Auction Service
- Complete auction management implementation
- Auction scheduling and automatic ending
- Search and filtering capabilities

### Advanced Features
- Payment integration simulation
- Email notifications
- Auction categories and tags
- Advanced bidding strategies (proxy bidding)
- Rate limiting and security enhancements

### Infrastructure
- Service discovery with Consul
- Message queues (NATS/RabbitMQ)
- Monitoring and metrics
- Kubernetes deployment manifests
- CI/CD pipeline

## 🐛 Troubleshooting

### Common Issues

1. **Proto generation issues**
   ```bash
   # Make sure protoc is installed
   sudo apt install protobuf-compiler
   
   # Install Go plugins
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

2. **Database connection issues**
   ```bash
   # Check if databases are running
   docker-compose -f deployments/docker-compose.yml ps
   
   # Check logs
   docker-compose -f deployments/docker-compose.yml logs postgres
   ```

3. **Module dependency issues**
   ```bash
   # Clean and retry
   go clean -modcache
   go mod tidy
   ```

## 📄 License

MIT License

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Add tests for new functionality
5. Commit changes: `git commit -m 'Add amazing feature'`
6. Push to branch: `git push origin feature/amazing-feature`
7. Submit a pull request
