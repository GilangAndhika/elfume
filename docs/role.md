# 🎭 **Role Management API**

This section covers **role-related** endpoints, including **creating roles**.

---

## **Create a Role**
### **Endpoint:** `POST /role/create`
Creates a new user role.

**Request Body (JSON)**
```json
{
    "role_name": "Admin"
}
```

**✅ Success Response**
```json
{
    "message": "Role created successfully",
    "role": {
        "role_id": "67aff19a533432bc3af88fe2",
        "role_name": "Admin"
    }
}
```

**Error Responses**
- **400 Bad Request** – Missing required fields.
- **500 Internal Server Error** – Database error.

---

## 🔒 **Security Notes**
- **Roles are used for access control**, ensuring only authorized users can access certain endpoints.
- **Each user has a `role_id`** that determines their level of access.

---

## 🚀 **Next Steps**
- 🔐 **[Authentication API](auth.md)** - Register, login, and logout.
- 👥 **[User Management API](user.md)** - Manage user accounts.
- 🌸 **[Perfume Management API](perfume.md)** - Manage perfume products.

---