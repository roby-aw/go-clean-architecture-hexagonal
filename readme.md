# Go Clean Architecture Hexagonal Project

This project implementation was inspired by port and adapter pattern or known as hexagonal:
In general is divided into 3 major parts, namely primary (driving adapter), business, and secondary (driven adapter).
- **Primary / driving adapter**<br/>driving adapter is a technology that we use to interact with users such as REST API, Graphql, gRPC, and so on. (also called user-side adapters in hexagonal's term)
- **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
- **Secondary / driven adapter**<br/>Contains implementations of interfaces defined in the business such as databases, external APIs, clouds, and so on. (also called as server-side adapters in hexagonal's term)

## What is Clean Architecture?
Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. The main rule of clean architecture is that code dependencies can only come from the outer levels inward. Code on the inner layers can have no knowledge of functions on the outer layers. The more external a component is, the higher level of abstraction it must have.

## Benefits of Clean Architecture Hexagonal
- **Independent of Frameworks**<br/>The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- **Testable**<br/>The business rules can be tested without the UI, Database, Web Server, or any other external element.
- **Maintainable**<br/>The Dependency Rule ensures that the coupling between components is always pointing inwards. This greatly reduces the risk of changing external code.

## Table of Contents
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [File Structure](#file-structure)
- [TL;DR](#tldr)
- [Start From Scratch](#start-from-scratch)
- [Coding Style](#coding-style)
- [Useful References](#useful-references)

## Project Structure

This project follows a clean architecture with the following main components:
- `api/`: Contains the REST API implementation like middleware and routes.
- `app/`: Contains the main application modules.
- `business/`: It is the business/service layer of the application and all logic will be put here.
- `config/`: Contains the configuration of the application for environment.
- `docs/`: Contains the swagger documentation.
- `repository/`: Contains the implementation of the interfaces defined in the business layer.
- `utils/`: Contains the utility functions that can be used in the application.
- `utils/driver/`: Contains the database connection.

## Tech Stack
- Framework : Fiber
- jose2go 
- go-json
- validator-v10
- godotenv
- swaggo
- google/uuid
- mongo-driver
- primitive
- Docker
- Jenkins

---
## File Structure
```
.               
├── api/
│   ├── middlewares           
│   │   └── middlewares.go   
│   ├── user/                
│   │   └── controller.go     
│   └── router.go             
├── app/                     
│   ├── modules/             
│   │   └── modules.go        
│   └── router.go 
├── business/
│   └── user/                
│       ├── user.go          
│       └── service.go        
├── config/                  
│       └── config.go         
├── repository/
│   ├── user/                
│   │    ├── factory.go      
│   │    └── mongo.go         
│   └── model.go 
├── utils/                   
│   ├── driver.go            
│   ├── jwt.go               
│   ├── handleError.go     
│   └── hashencrypt.go       
├── .env                     
├── .gitignore               
├── Dockerfile               
├── go.mod                   
├── go.sum                   
├── Jenkinsfile              
└── main.go                  
```
---
## TL;DR
This document will guide you some notices or warning usages, this will not teach you how to use Golang, if so, please refer [Golang Documentation](https://go.dev/doc/tutorial/getting-started) to build the app.

---
## Start From Scratch
Please clone this repository to your computer and run the shell command (recommended):
```
git clone https://github.com/roby-aw/go-clean-architecture-hexagonal.git
```
Downloading all of the modules in the dependency graph, which it can determine from reading only the go.mod files.:
```
go mod download
```
If all set up, then run the command:
```
go run main.go
```
Open the browser (recommended using Chrome) surf the URL:
```
http://localhost:8080/
```
You will see response api

For build the app :
```
go build .
```

---
## Coding Style
### For Golang
- Variabel
    - Use Camel case for naming your variables : var myVariable = ""
    - Use short but descriptive names for variables : var res Result
    - Use single letters for indexes : for i:=0;i<100:i++{} 
- Folder & file name : lowercase

In Response API must be :
- For error must have code and message
```
return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
```
- For result success must have code,message,and result is optional
```
return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success get data",
		"result":  res,
	})
```
- for handle error in service/bussines layer must use this:
```
return utils.HandleError(404, "id not found") // if we use this handle error, we can custom response code from service/business layer
```
- In Golang import:
```golang
import userBusiness "business/user" // import here
```
- In func Controller:
```golang
func (Controller *Controller) Login (c *fiber.ctx)
```
- In func business:
```golang
func (s *service) Login (auth *authLogin)(responseLogin, error)
```
- In func repository
```golang
func (repo *repository) FindUserByEmail (email string)(*userBusiness.User, error)
```
### Repository
Using design pattern for query data if not aggregate file on [here](./repository/model.go)
# Useful References
- [Golang Official Documents](https://go.dev/doc/)
- [Fiber Framework Documents](https://docs.gofiber.io/)
- [Echo Framework Documents](https://echo.labstack.com/guide/)