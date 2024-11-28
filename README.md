# w13_go_1
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
http://localhost:8080/current-time: 
```
Retrieve all logged times from the database.
```bash
http://localhost:8080/all-times: 
```

