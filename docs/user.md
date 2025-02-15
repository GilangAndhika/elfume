# 👥 **User Management API**

This section covers all **user-related** endpoints, including **retrieving, updating, and deleting users**.

---

## **Get All Users**
### **Endpoint:** `GET /user/all`
Retrieves all users from the database.

**✅ Success Response**
```json
{
    "message": "Users retrieved successfully",
    "users": [
        {
            "user_id": "609c5f9...",
            "username": "testuser",
            "email": "test@example.com",
            "phone": "08123456789",
            "role_id": "67aff183533432bc3af88fe1",
            "created_at": "2024-02-15T12:00:00Z",
            "updated_at": "2024-02-15T12:00:00Z"
        }
    ]
}
```

**Error Responses**
- **500 Internal Server Error** – Database error

---

## **Get User by ID**
### **Endpoint:** `GET /user/id/:id`
Retrieves a **specific user** by their ID.

**Example Request**
```sh
GET http://localhost:3000/user/id/609c5f9...
```

**✅ Success Response**
```json
{
    "message": "User retrieved successfully",
    "user": {
        "user_id": "609c5f9...",
        "username": "testuser",
        "email": "test@example.com",
        "phone": "08123456789",
        "role_id": "67aff183533432bc3af88fe1",
        "created_at": "2024-02-15T12:00:00Z",
        "updated_at": "2024-02-15T12:00:00Z"
    }
}
```

**Error Responses**
- **404 Not Found** – User ID does not exist
- **500 Internal Server Error** – Database error

---

## **Update User**
### **Endpoint:** `PUT /user/update/:id`
Updates an existing user's information.

**Example Request**
```sh
PUT http://localhost:3000/user/update/609c5f9...
```

**Request Body (JSON)**
```json
{
    "username": "newusername",
    "email": "newemail@example.com",
    "phone": "08123456789",
    "role_id": "67aff183533432bc3af88fe1"
}
```

**✅ Success Response**
```json
{
    "message": "User updated successfully"
}
```

**Error Responses**
- **400 Bad Request** – Missing required fields
- **404 Not Found** – User not found
- **500 Internal Server Error** – Database error

---

## **Delete User**
### **Endpoint:** `DELETE /user/delete/:id`
Deletes a user from the system.

**Example Request**
```sh
DELETE http://localhost:3000/user/delete/609c5f9...
```

**✅ Success Response**
```json
{
    "message": "User deleted successfully"
}
```

**Error Responses**
- **404 Not Found** – User does not exist
- **500 Internal Server Error** – Database error

---

## 🔒 **Security Notes**
- **User data is protected**; only authorized users should access these endpoints.
- **Passwords are encrypted** and cannot be retrieved in plaintext.
- **Deleting a user removes all associated data permanently.**

---

## 🚀 **Next Steps**
- 🔐 **[Authentication API](auth.md)** - Register, login, and logout.
- 🌸 **[Perfume Management API](perfume.md)** - Manage perfume products.

---