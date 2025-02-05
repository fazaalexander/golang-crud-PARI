# Super PARI Backend Case Study

## Created using Clean Architecture -> ![alt text](image.png) <-(VS CODE: hover to view the image)
- Basically, the principle of clean coding aims to separate responsibilities, making it:
    - independent of frameworks, allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
    - independent of database, you can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, Postgres, or something else. Your business rules are not bound to the database.

## How to setup and run the project:
- make sure your go version is more than 1.20 (required)
- Create MySQL database with the name "superparidb"
- Customize the database data in the .env file according to your local database configuration
- Run the project by simply sending the command “go run main.go” in the terminal

## Functional Requirements:
- CRUD operations for items (Case Study)

## Endpoint Spesifications:
- GET /items: Return a list of all items.
- GET /items/{id}: Return a single item by ID.
- POST /items: Create a new item. The request body should contain the item details.
- PUT /items/{id}: Update an existing item by ID. The request body should contain the updated item details.
- DELETE /items/{id}: Delete an item by ID.

## Tech Stack (Core):
- Golang (Programming Language)
- Echo (Go Framework)
- MySQL (database) + GORM (for ORM purposes)
- Go Libraries, such as:
    - Go playground validator (to help validate user request data)

## Tech Stack (Additional):
- Postman (API Documentation)