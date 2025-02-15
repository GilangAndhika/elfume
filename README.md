# 🚀 **Elfume API**

Welcome to the **Elfume API**! This API serves as the backend for managing perfume products, users, roles, and authentication. Built using **Go Fiber**, **MongoDB**, and **JWT Authentication**, this API is designed for fast, secure, and scalable web applications.

---

## **Features**

- **User Authentication:** Register, login, and secure sessions using JWT.
- **Role-Based Access Control:** Manage user roles for different levels of access.
- **Perfume Management:** Create, update, search, and delete perfume products with image uploads.
- **Protected Routes:** Secure API endpoints with JWT-based authentication.
- **Seamless Image Uploads:** Upload and manage images directly to GitHub for easy access.

---

## **Installation & Setup**

### 1️ **Clone the Repository**
```sh
git clone https://github.com/yourusername/elfume.git
cd elfume
```

### 2️ **Install Dependencies**
```sh
go mod tidy
```

### 3️ **Setup Environment Variables**

Create a `.env` file in the root directory and add:
```ini
MONGO_URI=mongodb+srv://your_user:your_password@your_cluster.mongodb.net/
MONGO_DB=elfume
JWT_SECRET=your_secret_key

GITHUB_OWNER=your_github_username
GITHUB_REPO=your_github_repo
GITHUB_TOKEN=your_github_token
```

### 4️ **Run the Application**
```sh
go run main.go
```

---

## **API Documentation**

The detailed API documentation is available for each group of endpoints. Click on the links below to learn more about each section:

| Route Group     | Description                                  | Documentation                |
|-----------------|----------------------------------------------|------------------------------|
| 🔐 **Auth**     | Register, login, and logout users            | [View Auth Docs](docs/auth.md) |
| 👥 **Users**    | Manage users (CRUD operations)               | [View User Docs](docs/user.md) |
| 🎭 **Roles**    | Create and manage user roles                 | [View Role Docs](docs/role.md) |
| 🌸 **Perfumes** | Manage perfume products and images           | [View Perfume Docs](docs/perfume.md) |
| 🔒 **Protected**| Access protected routes with JWT             | [View Protected Docs](docs/protected.md) |

---

## **Running with Docker**

You can also run the API using **Docker**:

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

---

## **Why Use the Elfume API?**

- **📦 Scalable:** Built with **Go Fiber**, known for its speed and scalability.
- **🔐 Secure:** Uses **JWT Authentication** for secure API access.
- **🌸 Perfume-Focused:** Specifically designed for managing perfume products.
- **📚 Easy-to-Follow Docs:** Comprehensive documentation for each endpoint group.

---

Get started with the **Elfume API** and make managing perfume products effortless! 🌸🔥🚀