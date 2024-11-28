# w13_go_1

This Go application provides an API that returns the current time in Toronto, stores it in a MySQL database, and allows retrieval of all stored time records. It demonstrates how to work with MySQL in Go, handle time zone conversions, and expose RESTful API endpoints.
The application connects to a MySQL database and performs two main operations:
- Insert the current time (in Toronto's local time) into the database.
- Retrieve and display all stored time records from the database

Install MySQL on macos
``` bash 
brew install mysql
```
Start MySQL Service
``` bash 
brew services start mysql
```
Secure Installation:
``` bash 
mysql_secure_installation
```
Access MySQL
``` bash 
mysql -u root -p
```
Create Database and Table
``` bash
CREATE DATABASE time_api;
USE time_api;

CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```
Install Dependencies:
Ensure you have Go installed.
Install the MySQL driver:
``` bash
go get -u github.com/go-sql-driver/mysql
```
Access Endpoints:

Get the current Toronto time and log it in the database. 
```bash 
http://localhost:8080/current-time
```
![image](https://github.com/RishwanthReddy17/user-attachments/raw/721390bb3ef7576fc6cf9edd5d6eea5947ba20d2/1.png)

Retrieve all logged times from the database.
```bash
http://localhost:8080/all-times
```
![image](https://github.com/RishwanthReddy17/user-attachments/raw/721390bb3ef7576fc6cf9edd5d6eea5947ba20d2/2.png)

DSN for MYSQL:
```bash
dsn := "user:securepass@tcp(mysql-container:3306)/time_api"
```


