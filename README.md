# CRM SAAS - Multi-Tenant Customer Relationship Management System

Backend API untuk sistem CRM SAAS berbasis Golang dengan arsitektur multi-tenant.

## 🚀 Fitur Utama

### User & Tenant Management

- **Multi-Tenant Architecture**: Setiap tenant memiliki data yang terisolasi
- **User Registration dengan Auto Tenant Creation**: Ketika tenant pertama dibuat, otomatis mendaftarkan user sebagai owner
- **Role-Based Access Control (RBAC)**: Owner, Admin, Sales
- **JWT Authentication**: Secure token-based authentication

### CRM Features

- Customer Management
- Lead Management
- Deal Management
- Pipeline & Stage Management
- Activity Tracking
- Notes Management
- Reminder System

## 📁 Struktur Project

```
server/
├── cmd/
│   └── main.go                 # Entry point aplikasi
├── internal/
│   ├── tenant/                 # Tenant module
│   ├── user/                   # User module
│   ├── customer/               # Customer module
│   ├── lead/                   # Lead module
│   ├── deal/                   # Deal module
│   ├── pipeline_stage/         # Pipeline Stage module
│   ├── activity/               # Activity module
│   ├── note/                   # Note module
│   └── reminder/               # Reminder module
├── pkg/
│   ├── config/                 # Configuration management
│   ├── middlewares/            # HTTP middlewares
│   ├── response/               # Response helpers
│   ├── validator/              # Input validation
│   └── utils/                  # Utility functions
├── .env.example                # Environment variables template
├── go.mod                      # Go module dependencies
└── README.md                   # Documentation
```

## 🛠️ Tech Stack

- **Language**: Go 1.25+
- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: MySQL
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Validation**: go-playground/validator

## 📦 Installation

### Prerequisites

- Go 1.25 atau lebih tinggi
- MySQL 8.0+
- Git

### Setup Steps

1. **Clone repository**

   ```bash
   git clone <repository-url>
   cd CRM-SAAS/server
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Setup environment variables**

   ```bash
   cp .env.example .env
   ```

   Edit file `.env` dan sesuaikan dengan konfigurasi Anda

4. **Create database**

   ```bash
   mysql -u root -p
   CREATE DATABASE crm_saas_db;
   ```

5. **Run application**
   ```bash
   go run cmd/main.go
   ```

## 🔐 Authentication Flow

### 1. Register Tenant & Owner (First Time)

```bash
POST /api/tenants/register

{
  "tenant_name": "PT. ABC Company",
  "tenant_domain": "abc-company",
  "owner_name": "John Doe",
  "owner_email": "john@abc.com",
  "owner_password": "password123"
}
```

### 2. Login

```bash
POST /api/users/login

{
  "email": "john@abc.com",
  "password": "password123"
}
```

## 🔑 User Roles

- **Owner**: Full access, dibuat otomatis saat registrasi tenant
- **Admin**: Mengelola user dan CRM features
- **Sales**: Akses CRM features

## 📄 License

This project is licensed under the MIT License.
