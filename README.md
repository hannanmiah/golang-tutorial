# Go E-Commerce API Tutorial

A comprehensive e-commerce REST API built with Go (Golang) using the Gin web framework, GORM for database operations, and JWT authentication. This tutorial project demonstrates best practices for building scalable web applications in Go.

## 🚀 Features

- **User Authentication**: Registration, login with JWT tokens
- **Product Management**: CRUD operations for products
- **Shopping Cart**: Add, update, and remove items from cart
- **Order Management**: Create and manage orders with status tracking
- **Role-based Access**: User and admin role permissions
- **Database**: SQLite with GORM ORM
- **Middleware**: Authentication and authorization middleware

## 📋 Requirements

- Go 1.25.4 or higher
- SQLite3

## 🛠️ Installation & Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/hannanmiah/golang-tutorial.git
   cd golang-tutorial
   ```

2. **Install dependencies**
   ```bash
   make tidy
   ```

3. **Install development tools (optional)**
   ```bash
   make install-tools
   ```

## 🏃‍♂️ Running the Application

### Development Mode (with hot reload)
```bash
make dev
```

### Production Mode
```bash
make run
```

### Manual Build & Run
```bash
make build
./bin/server
```

## 📊 Database

The application uses SQLite as the database. The database file (`ecommerce.db`) and all tables are automatically created when you run the application.

### Database Models

- **User**: User accounts with authentication and roles
- **Product**: Product catalog with pricing and inventory
- **Cart**: Shopping cart functionality
- **Order**: Order management with items
- **OrderItem**: Individual items within orders

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

## 📚 API Endpoints

### Public Endpoints

#### Authentication
- `POST /register` - Register a new user
- `POST /login` - User login
- `GET /` - API welcome message

### Protected Endpoints (Require Authentication)

#### User Management
- `GET /profile` - Get user profile

#### Product Management
- `GET /products` - Get all products
- `GET /products/:id` - Get specific product
- `POST /products` - Create new product
- `PUT /products/:id` - Update product
- `DELETE /products/:id` - Delete product
- `GET /my-products` - Get current user's products

#### Cart Management
- `GET /cart` - Get user's cart
- `POST /cart` - Add item to cart
- `PUT /cart/:id` - Update cart item
- `DELETE /cart/:id` - Remove item from cart
- `DELETE /cart` - Clear entire cart

#### Order Management
- `GET /orders` - Get user's orders
- `GET /orders/:id` - Get specific order
- `POST /orders` - Create new order

### Admin Endpoints (Require Admin Role)

#### Order Administration
- `GET /admin/orders` - Get all orders (admin only)
- `PUT /admin/orders/:id/status` - Update order status (admin only)

## 🗂️ Project Structure

```
golang-tutorial/
├── cmd/
│   └── migrate/           # Database migration utilities
├── handlers/              # HTTP request handlers
│   ├── user.go           # User-related handlers
│   ├── product.go        # Product-related handlers
│   ├── cart.go           # Shopping cart handlers
│   └── order.go          # Order management handlers
├── middleware/            # Custom middleware
│   └── auth.go           # Authentication & authorization
├── models/               # Data models and database schemas
│   └── models.go         # All database models
├── functions/            # Utility functions and examples
├── main.go               # Application entry point
├── go.mod                # Go module dependencies
├── go.sum                # Dependency checksums
├── Makefile              # Build and development commands
├── .air.toml            # Air hot reload configuration
└── ecommerce.db         # SQLite database file (auto-generated)
```

## 🔧 Available Commands

```bash
make help        # Show all available commands
make tidy        # Download and organize dependencies
make migrate     # Run database migrations
make run         # Start the API server
make dev         # Run in development mode with auto-reload
make build       # Build the application
make clean       # Clean build artifacts
make install-tools  # Install development tools
```

## 📦 Dependencies

- **gin-gonic/gin** - HTTP web framework
- **gorm.io/gorm** - ORM for database operations
- **gorm.io/driver/sqlite** - SQLite driver for GORM
- **golang-jwt/jwt/v5** - JWT authentication
- **golang.org/x/crypto** - Cryptographic functions (bcrypt)

## 🎯 Usage Examples

### User Registration
```bash
curl -X POST http://localhost:8000/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### User Login
```bash
curl -X POST http://localhost:8000/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Product (Authenticated)
```bash
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "name": "Awesome Product",
    "description": "A great product description",
    "price": 29.99,
    "stock": 100
  }'
```

## 🤝 Contributing

This is a tutorial project. Feel free to fork it and use it for learning purposes.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🔗 Useful Resources

- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [JWT Authentication in Go](https://github.com/golang-jwt/jwt)
- [Go Official Documentation](https://golang.org/doc/)

---

**Built with ❤️ using Go**