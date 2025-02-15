# 🌸 **Perfume Management API**

This section covers all **perfume-related** endpoints, including **creating, retrieving, updating, searching, and deleting perfumes**.

---

## **Create a Perfume**
### **Endpoint:** `POST /fume/create`
Uploads a new **perfume product** with an image.

***Request Type:** `multipart/form-data`
| Key         | Type          | Value (Example) |
|------------|--------------|----------------|
| `name`      | Text         | `Ocean Breeze` |
| `brand`     | Text         | `Aqua Scents`  |
| `types`     | Text         | `Eau de Parfum` |
| `categories`| Text         | `Fresh`        |
| `sizes`     | Text         | `100ml`        |
| `price`     | Text         | `50`           |
| `description` | Text       | `A refreshing ocean breeze scent.` |
| `stock`     | Text         | `10`           |
| `image`     | **File**     | **Upload an image file** |

**✅ Success Response**
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

**Error Responses**
- **400 Bad Request** – Missing required fields or invalid image format.
- **500 Internal Server Error** – Failed to upload image or insert into database.

---

## **Get All Perfumes**
### **Endpoint:** `GET /fume/all`
Retrieves all perfumes from the database.

**✅ Success Response**
```json
{
    "message": "Perfumes retrieved successfully",
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

**Error Responses**
- **500 Internal Server Error** – Database error.

---

## **Get Perfume by ID**
### **Endpoint:** `GET /fume/id/:id`
Retrieves a **specific perfume** by its ID.

**Example Request**
```sh
GET http://localhost:3000/fume/id/609c5f9...
```

**✅ Success Response**
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

**Error Responses**
- **404 Not Found** – Perfume ID does not exist.
- **500 Internal Server Error** – Database error.

---

## **Search Perfumes**
### **Endpoint:** `GET /fume/search`
Allows searching for perfumes using **filters**.

**Example Queries**
```sh
GET http://localhost:3000/fume/search?brand=Dior
GET http://localhost:3000/fume/search?name=Sauvage
GET http://localhost:3000/fume/search?size=100&brand=Dior
```

**✅ Success Response**
```json
{
    "message": "Perfumes retrieved successfully",
    "perfumes": [
        {
            "perfume_id": "67b0255f0616428b90c65b24",
            "name": "Dior Sauvage",
            "brand": "Dior",
            "types": "Eau de Parfum",
            "categories": "Fresh",
            "sizes": "100ml",
            "image": "https://raw.githubusercontent.com/yourgithubowner/yourgithubrepo/main/dior_sauvage.webp",
            "price": "100000",
            "description": "A wild and fresh masculine fragrance.",
            "stock": "5",
            "created_at": "2025-02-15T05:14:54.626Z",
            "updated_at": "2025-02-15T05:14:54.626Z"
        }
    ]
}
```

**Error Responses**
- **500 Internal Server Error** – Database error.

---

## **Update Perfume**
### **Endpoint:** `PUT /fume/update/:id`
Updates an existing perfume’s information.

**Example Request**
```sh
PUT http://localhost:3000/fume/update/609c5f9...
```

**Request Body (JSON)**
```json
{
    "name": "Dior Sauvage Intense",
    "brand": "Dior",
    "types": "Eau de Parfum",
    "categories": "Woody",
    "sizes": "100ml",
    "price": "110000",
    "description": "An intense and elegant masculine fragrance.",
    "stock": "7"
}
```

**✅ Success Response**
```json
{
    "message": "Perfume updated successfully"
}
```

**Error Responses**
- **400 Bad Request** – Missing required fields.
- **404 Not Found** – Perfume not found.
- **500 Internal Server Error** – Database error.

---

## **Delete Perfume**
### **Endpoint:** `DELETE /fume/delete/:id`
Deletes a perfume.

**Example Request**
```sh
DELETE http://localhost:3000/fume/delete/609c5f9...
```

**✅ Success Response**
```json
{
    "message": "Perfume deleted successfully"
}
```

**Error Responses**
- **404 Not Found** – Perfume does not exist.
- **500 Internal Server Error** – Database error.

---

## 🔒 **Security Notes**
- **Only authorized users** can create, update, or delete perfumes.
- **Image uploads are securely stored on GitHub** and linked via URL.

---

## 🚀 **Next Steps**
- 🔐 **[Authentication API](auth.md)** - Register, login, and logout.
- 👥 **[User Management API](user.md)** - Manage user accounts.

---