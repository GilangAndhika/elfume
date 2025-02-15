# 🚀 Elfume API

Welcome to the **Elfume API**, a RESTful API built using **Go Fiber**, **MongoDB**, and **JWT Authentication**.

## Features
- User authentication (Register, Login)
- Role-based access control
- Secure JWT authentication with HTTP-only cookies
- MongoDB as the database
- **Perfume product management with image uploads**

## 📌 Installation & Setup

### 1️ **Clone the repository**
```sh
git clone https://github.com/GilangAndhika/elfume.git
cd elfume
```

### 2️ **Install dependencies**
```sh
go mod tidy
```

### 3️ **Setup environment variables**
Create a `.env` file in the root directory and add:
```ini
MONGO_URI=mongodb+srv://your_user:your_password@your_cluster.mongodb.net/
MONGO_DB=elfume
JWT_SECRET=your_secret_key

GITHUB_OWNER=your_github_username
GITHUB_REPO=your_github_repo
GITHUB_TOKEN=your_github_token
```

### 4️ **Run the application**
```sh
go run main.go
```

---

## 📌 API Endpoints

### **Base URL**: `http://localhost:3000`

---

## **🔑 Authentication**
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

## **👥 User Routes**
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

## **🌸 Perfume Routes**
| Method  | Endpoint           | Description                      | Authentication |
|---------|--------------------|----------------------------------|----------------|
| `POST`  | `/perfume/create`  | Create a new perfume with image | ✅ Requires JWT |
| `GET`   | `/perfume/all`     | Get all perfumes                | ❌ No auth required |
| `GET`   | `/perfume/:id`     | Get a perfume by ID             | ❌ No auth required |

---

### **🔹 Create Perfume (`/perfume/create`)**
Uploads a **new perfume product** along with an **image file** to GitHub.

#### **📌 Request Type:** `multipart/form-data`
| Key         | Type          | Value (Example)                   |
|------------|--------------|----------------------------------|
| `name`      | Text         | `Ocean Breeze`                  |
| `brand`     | Text         | `Aqua Scents`                   |
| `types`     | Text         | `Eau de Parfum`                 |
| `categories`| Text         | `Fresh`                          |
| `sizes`     | Text         | `100ml`                          |
| `price`     | Text         | `50`                             |
| `description` | Text       | `A refreshing ocean breeze scent.` |
| `stock`     | Text         | `10`                             |
| `image`     | **File**     | **Upload an image file**         |

#### **✅ Example Response**
```json
{
  "message": "Perfume created successfully",
  "perfume": {
    "perfume_id": "609c5f9...",
    "name": "Ocean Breeze",
    "brand": "Aqua Scents",
    "types": "Eau de Parfum",
    "categories": "Fresh",
    "sizes": "100ml",
    "image": "https://raw.githubusercontent.com/yourgithubowner/yourgithubrepo/main/ocean_breeze.jpg",
    "price": "50",
    "description": "A refreshing ocean breeze scent.",
    "stock": "10",
    "created_at": "2024-02-15T12:00:00Z",
    "updated_at": "2024-02-15T12:00:00Z"
  }
}
```

---

### **🔹 Get All Perfumes (`/perfume/all`)**
#### **📌 Request Type:** `GET`
Retrieves **all perfumes** stored in the database.

#### **✅ Example Response**
```json
{
  "message": "Perfume list retrieved successfully",
  "perfumes": [
    {
      "perfume_id": "609c5f9...",
      "name": "Ocean Breeze",
      "brand": "Aqua Scents",
      "types": "Eau de Parfum",
      "categories": "Fresh",
      "sizes": "100ml",
      "image": "https://raw.githubusercontent.com/yourgithubowner/yourgithubrepo/main/ocean_breeze.jpg",
      "price": "50",
      "description": "A refreshing ocean breeze scent.",
      "stock": "10",
      "created_at": "2024-02-15T12:00:00Z",
      "updated_at": "2024-02-15T12:00:00Z"
    }
  ]
}
```

---

### **🔹 Get Perfume by ID (`/perfume/:id`)**
#### **📌 Request Type:** `GET`
Retrieves **a single perfume** based on its `id`.

#### **✅ Example Response**
```json
{
  "message": "Perfume retrieved successfully",
  "perfume": {
    "perfume_id": "609c5f9...",
    "name": "Ocean Breeze",
    "brand": "Aqua Scents",
    "types": "Eau de Parfum",
    "categories": "Fresh",
    "sizes": "100ml",
    "image": "https://raw.githubusercontent.com/yourgithubowner/yourgithubrepo/main/ocean_breeze.jpg",
    "price": "50",
    "description": "A refreshing ocean breeze scent.",
    "stock": "10",
    "created_at": "2024-02-15T12:00:00Z",
    "updated_at": "2024-02-15T12:00:00Z"
  }
}
```

---

## **🎭 Role Routes**
| Method | Endpoint       | Description | Request Body |
|--------|--------------|-------------|--------------|
| `POST` | `/role/create` | Create a new role | ```json { "role_name": "Admin" }``` |

---

## **Run the API with Docker**
```sh
docker build -t elfume-api .
docker run -p 3000:3000 elfume-api
```

---

## **Contributors**
- **Gilang Andhika** - [GitHub](https://github.com/GilangAndhika)

---

## **License**
This project is licensed under the **MIT License**.