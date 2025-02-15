# 🔐 **Authentication API**

This section covers all authentication-related endpoints, including **user registration, login, and logout**.

---

## **Register a New User**
### **Endpoint:** `POST /auth/register`
Registers a new user in the system.

**Request Body (JSON)**
```json
{
    "username": "testuser",
    "email": "test@example.com",
    "password": "123456",
    "phone": "08123456789"
}
```

**✅ Success Response**
```json
{
    "message": "User registered successfully"
}
```

**Error Responses**
- **400 Bad Request** – Missing required fields
- **500 Internal Server Error** – Database error

---

## **Login**
### **Endpoint:** `POST /auth/login`
Authenticates a user and returns a **JWT token** in an **HTTP-only cookie**.

**Request Body (JSON)**
```json
{
    "email": "test@example.com",
    "password": "123456"
}
```

**✅ Success Response**
```json
{
    "message": "Login successful"
}
```
🔹 **Note:** The JWT token is set in an **HTTP-only cookie** for security.

**Error Responses**
- **401 Unauthorized** – Invalid credentials
- **500 Internal Server Error** – Database error

---

## **Logout**
### **Endpoint:** `POST /auth/logout`
Clears the JWT authentication cookie, effectively logging out the user.

**✅ Success Response**
```json
{
    "message": "Logged out successfully"
}
```

**Error Responses**
- **400 Bad Request** – If the user is not logged in

---

## 🔒 **Security Measures**
- **JWT Authentication:** Tokens are stored in **HTTP-only cookies** to prevent XSS attacks.
- **Secure Login:** Encrypted passwords using **bcrypt**.
- **Session Management:** Users stay logged in via cookies until they log out.

---

## 🚀 **Next Steps**
- ✅ **[User Management API](user.md)** - Get, update, and delete users.
- 🌸 **[Perfume Management API](perfume.md)** - Manage perfume products.

---