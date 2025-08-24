# Identity Reconciliation
A service for identity data reconciliation and record matching.  

This project allows managing user contact information (emails and phone numbers), ensuring **duplicate detection**, **primary/secondary contact handling**, and **linked record management**.

---

## 🚀 Tech Stack
- **Language:** Go (Golang)  
- **Framework:** [Gin](https://github.com/gin-gonic/gin) HTTP Web Framework  
- **Database:** PostgreSQL  
- **ORM:** [GORM](https://gorm.io/)  

---

## 📌 Features
- Store contact details with **email** and **phone number**.  
- Detect if an incoming record matches existing contacts.  
- Maintain **Primary** and **Secondary** contact precedence.  
- Automatically link secondary contacts to the correct primary contact.  
- Provide all linked records for a given identity.  

---

## 📂 Project Structure
```
.
├── database/        # DB connection and initialization
├── models/          # GORM models (Contact, etc.)
├── helpers/         # Business logic for reconciliation
├── routes/          # API route definitions
├── main.go          # Application entry point
└── go.mod           # Dependencies
```

---

## ⚙️ Setup & Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/<your-username>/identity-reconciliation.git
cd identity-reconciliation
```

### 2️⃣ Install dependencies
```bash
go mod tidy
```

### 3️⃣ Configure Environment Variables
Create a `.env` file in the root folder:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=identity_db
PORT=8080
```

### 4️⃣ Run the application
```bash
go run main.go
```

The server will start on:
```
http://localhost:8080
```

---

## 📡 API Endpoint

### **POST /api/identify**
Identifies and reconciles a user’s contact details.  

#### Request Body:
```json
{
  "email": "test@example.com",
  "phoneNumber": "1234567890"
}
```

#### Response Example:
```json
{
  "contact": {
    "primaryContactId": 1,
    "emails": ["test@example.com", "alt@example.com"],
    "phoneNumbers": ["1234567890", "9876543210"],
    "secondaryContactIds": [2, 3]
  }
}
```
