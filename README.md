# E-Commerce Project
This Repository focuses on the backend, built with Go (Golang).
## Frontend Part
<img width="1914" height="807" alt="image" src="https://github.com/user-attachments/assets/2db25357-a093-429d-817a-2a63a043fedc" />

 ##  Backend Project Structure
 
```plaintext
e-commerce-project/
├── cmd/
│   └── myApp/
│       └── serve.go                # server file
├── internal/
│   ├── api/
|   |   └── handlers
|   |       └── create_products.go  # create product
|   |       └──get_products.go      # get all product
|   |       └──global_router.go     # create global router
|   |   └── routes
|   |       └── routes.go           # api
│   ├── models/
│   │   └── product.go              # Data models
│   └── services/
│       └── service.go              # Business logic 
├── go.mod                          # Go module definition
└── main.go                         # main file
```
## 🚀 Usage
### 1. Start the Go Backend Server

```bash
cd Backend
go run main.go
```
The backend server will run on http://localhost:8080.

### 2. Start the React App

```bash
cd Frontend
npm start
```
The React App will run on http://localhost:3000.

### 3. Access the App:
Open  web browser and navigate to http://localhost:3000 to interact with the React app.
