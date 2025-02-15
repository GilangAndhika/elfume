# 🔒 **Protected Routes API**

This section covers **JWT-protected** endpoints that require authentication.

---

## **Access Protected Content**
### **Endpoint:** `GET /protected`
Retrieves protected content that requires a **valid JWT token**.

**Example Request**
```sh
GET http://localhost:3000/protected
```

✅ **This request must include a valid JWT token** stored in an **HTTP-only cookie**.

**✅ Success Response**
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

**Error Responses**
- **401 Unauthorized** – Missing or invalid token.
- **403 Forbidden** – User does not have the required permissions.

---

## 🔐 **How to Use JWT for Authentication**
- After **login**, the **JWT token** is stored in an **HTTP-only cookie**.
- The **token is automatically sent** with each request, allowing access to **protected routes**.
- To **log out**, call `POST /auth/logout`, which clears the cookie.

---

## 🚀 **Next Steps**
- 🔐 **[Authentication API](auth.md)** - Register, login, and logout.
- 👥 **[User Management API](user.md)** - Manage user accounts.
- 🎭 **[Role Management API](role.md)** - Assign roles to users.

---