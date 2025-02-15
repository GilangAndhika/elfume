Here's your **README.md** fully written in Markdown format. 🚀  

---

### **📌 README.md**
```markdown
# 🚀 Elfume API

Welcome to the **Elfume API**, a RESTful API built using **Go Fiber**, **MongoDB**, and **JWT Authentication**.

## 📌 Features
- User authentication (Register, Login)
- Role-based access control
- Secure JWT authentication with HTTP-only cookies
- MongoDB as the database

---

## 📌 Installation & Setup

### 1️⃣ **Clone the repository**
```sh
git clone https://github.com/your-username/elfume.git
cd elfume
```

### 2️⃣ **Install dependencies**
```sh
go mod tidy
```

### 3️⃣ **Setup environment variables**
Create a `.env` file in the root directory and add:
```ini
MONGO_URI=mongodb+srv://your_user:your_password@your_cluster.mongodb.net/
MONGO_DB=elfume
JWT_SECRET=your_secret_key
```

### 4️⃣ **Run the application**
```sh
go run main.go
```

---

## 📌 API Endpoints

### 🏠 **Base URL**: `http://localhost:3000`

### **🔑 Authentication**
| Method | Endpoint         | Description          | Request Body |
|--------|----------------|----------------------|--------------|
| `POST` | `/auth/register` | Register a new user | ```json { "username": "test", "email": "test@example.com", "password": "123456", "phone": "08123456789" }``` |
| `POST` | `/auth/login` | Login user & get JWT | ```json { "email": "test@example.com", "password": "123456" }``` |

**🔹 Response (on success)**:
```json
{
    "message": "Login successful",
    "token": "eyJhbGciOi..."
}
```
**🔹 The JWT token is set in an HTTP-only cookie for security.**

---

### **👥 User Routes**
| Method | Endpoint       | Description          | Authentication |
|--------|--------------|----------------------|----------------|
| `GET`  | `/protected` | Access protected content | ✅ Requires JWT |

**🔹 Example Response:**
```json
{
    "message": "Access granted",
    "user": {
        "user_id": "609c5f9...",
        "username": "testuser",
        "role_id": "67aff183533432bc3af88fe1",
        "role_name": "Customer"
    }
}
```

---

### **🎭 Role Routes**
| Method | Endpoint       | Description | Request Body |
|--------|--------------|-------------|--------------|
| `POST` | `/role/create` | Create a new role | ```json { "role_name": "Admin" }``` |

**🔹 Example Response:**
```json
{
    "message": "Role created successfully",
    "role": {
        "role_id": "67aff19a533432bc3af88fe2",
        "role_name": "Admin"
    }
}
```

---

## 📌 Authentication & Security

- **JWT Authentication**: Tokens are stored in **HTTP-only cookies** to prevent **XSS attacks**.
- **Protected Routes**: Routes like `/protected` require **valid JWT tokens**.

To access protected routes, **include the JWT token in cookies**.

---

## 📌 Run the API with Docker
You can run the API in a **Docker container**:
```sh
docker build -t elfume-api .
docker run -p 3000:3000 elfume-api
```

---

## 📌 Contributors
- **Your Name** - [GitHub](https://github.com/your-username)

---

## 📌 License
This project is licensed under the **MIT License**.
```

---

### **✅ Why This is Useful**
- 📌 **Clear API Documentation** with examples
- 🔑 **Authentication Details**
- 🔐 **JWT Usage & Security**
- 🚀 **Docker Instructions**

Now your **README.md** is **professional and informative**! 🚀🔥 Let me know if you want to modify anything!