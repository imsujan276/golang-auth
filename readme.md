# Golang REST API Project

This is the simple Golang project that covers basic API functionalities like authentication, send verification code, middleware etc.

### Topics Covered

- Golang, Gin & GORM JWT Authentication
- JWT Authentication Example with Golang and GORM
- Generate the Private and Public Keys
- Load and Validate the Environment Variables
- Create the Database Models with GORM
- Run the Database Migration with GORM
- Hash and Verify the Passwords with Bcrypt
- Sign and Verify the RS256 JSON Web Tokens
    - Function to Generate the Tokens
    - Function to Verify the Tokens
- Create the Authentication Route Controllers
    - Register User Controller
    - Login User Controller
    - Refresh Access Token Controller
    - Logout Controller
- Create a Middleware Guard
- Create a User Controller
    - Authentication Routes
    - User Routes
- Add the Routes to the Gin Middleware Stack
- Forgot and reset password
- Role based route using middlewares

---

#### Generate env file
``` bash
cp app.env.example app.env
```

---
#### Generate Private and Public Keys

- Generate private and public keys from [HERE](http://travistidwell.com/jsencrypt/demo/)
- Copy the key andnavigate [HERE](https://www.base64encode.org/) to convert it into base64 string
- Assign the values to respective keys in **app.env** file 
---

#### Migrate the database
``` bash
go run migrate/migrate.go
```
---

#### Run the project
``` bash
go run cmd/*
```
---
