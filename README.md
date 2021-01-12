> This is a RESTFUL API for tasks where this have a CRUD in the tasks logic. This API use Golang, Mux, Gorm(ORM for Golang), Godotenv and PostgreSQL(For this project i use the cloud service database ElephantSQL).

> Created by Luca Becci

## 1. Started ⌨️

for get the project you will use:

```tsx
git clone "https://github.com/lucabecci/gorm-golang"

```

## 2. Pre-requeriments 🛠

You will need this requeriments for this repository:

- Golang

## 3. Installation 🔩

You will need to run this comands for the installation:

```
 Run go install for to install the dependencies.

```

## 4. Use the API

### How to run a pattern

1. In your terminal, navigate to the main directory.
2. Run `go run src/main.go` for to run the API.
3. Run `go build src/maing.go` for to create the build.


## 5. API structure 📁

### ROOT 📂

```tsx
|-- src
    |-- controllers	
    	|-- taskController.go
    |-- database
    	|-- connection.go
    |-- helpers
    	|-- helpers.go
    |-- models
        |-- Task.go
    |-- main.go 
|-- .env
|-- .gitignore
|-- go.mod
|-- go.sum
|-- README.md
```

## 6. Build with 🛠

DEPENDENCIES:

- [Mux](https://github.com/gorilla/mux)- Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.
- [Gorm](https://gorm.io/)- The fantastic ORM library for Golang.
- [GodotEnv](https://github.com/joho/godotenv)- Storing configuration in the environment is one of the tenets of a twelve-factor app. Anything that is likely to change between deployment environments–such as resource handles for databases or credentials for external services–should be extracted from the code into environment variables.

## 7. Database 📫

For the database i use MongoDB

- PostgreSQL
- Cloud Services: https://elephantsql.com

## 8. Versioned 1️⃣

For the versioning i use [ConventionalCommits] ([https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/))

## 9. Author 🙎🏻‍♂️

***Luca Becci -** code, documentation and implementation.*

- [github](https://github.com/lucabecci)
- [twitter](https://twitter.com/lucabecci)
- [linkedin](https://www.linkedin.com/in/luca-becci-b8044b198/)