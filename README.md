# Fiber User Management API

A simple and modular **User Management API** built using [Golang](https://golang.org/) and the [Fiber](https://github.com/gofiber/fiber) web framework.  
This project demonstrates a clean architecture approach with JWT authentication, user role management, and SQLite persistence.

---

## 🛠 Features

- ✅ Register / Login users
- 🔐 JWT Authentication (Access & Refresh Tokens)
- 👤 Role-based authorization (e.g. Admin, User)
- 🗃️ SQLite database with GORM ORM
- 🌿 Modular folder structure
- 📦 Fiber v3 support
- 🧪 Swagger & Postman support

---

## 📁 Project Structure

```
Fiber_UserManagement_API/
│
├── internal/database/       # DB initialization & connection
├── internal/user/middleware/     # Authentication & authorization logic
├── internal/model/          # GORM models (User struct)
├── internal/user/router/         # All routes and endpoints
├── docs/            #swager docs
├── main.go         # Entry point
└── go.mod
```

---

## 🚀 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/dayiamin/Fiber_UserManagement_API.git
cd Fiber_UserManagement_API
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Run the project
```bash
go run main.go
```

### 4. API will be available at:
```
http://localhost:3000
```

---

## 🔐 Authentication Flow

- **Login/Register** → Generates `access_token` and `refresh_token`
- **Protected Routes** → Require Bearer token in `Authorization` header
---

## 🧪 API Documentation

You can use the included Postman collection or set up Swagger support for testing and documentation:

- Swagger: *(coming soon or add instructions if supported)*
- Postman: Import the included collection file to test endpoints

---

## 🧰 Tech Stack

- [Go](https://golang.org/)
- [Fiber v3](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [JWT](https://jwt.io/)
- [SQLite](https://www.sqlite.org/index.html)

---


## 🙋‍♂️ Author

Developed by [@dayiamin](https://github.com/dayiamin)  
If you like this project, feel free to ⭐ it!

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).
