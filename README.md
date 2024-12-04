# Stock-Management-API
Go API used in a stock management web application.

This is a Go project that uses a MySQL database configured via XAMPP. Below are the instructions to run the project on your local machine with MySQL and XAMPP.

## Prerequisites

Before getting started, ensure you have the following installed:

- [XAMPP](https://www.apachefriends.org/) to run MySQL.
- [Go](https://golang.org/dl/) installed on your machine.

## Database Setup

1. **Install XAMPP**:
   - Download and install XAMPP from the [official website](https://www.apachefriends.org/).
   - After installation, open the XAMPP control panel and start the **MySQL** service.

2. **Create the Database and Tables**:
   - Open the XAMPP control panel and click **Shell** to open the MySQL terminal.
   - Connect to MySQL using the following command:
     ```bash
     mysql -u root -p
     ```
   - If prompted for a password, use `root` (default for XAMPP).
   - Then, create the database and tables required for the project. You can execute the provided `.sql` script included in this repository or copy the commands below:

     ```sql
     CREATE DATABASE IF NOT EXISTS stock_db;
     USE stock_db;

     -- Create 'categories' table
      CREATE TABLE IF NOT EXISTS categories (
          id INT(11) NOT NULL AUTO_INCREMENT,
          category_name VARCHAR(255) UNIQUE,
          PRIMARY KEY (id)
      );
      
      -- Create 'products' table
      CREATE TABLE IF NOT EXISTS products (
          id INT(11) NOT NULL AUTO_INCREMENT,
          name VARCHAR(255) NOT NULL,
          description TEXT,
          category VARCHAR(255),
          PRIMARY KEY (id)
      );
      
      -- Create 'registers' table
      CREATE TABLE IF NOT EXISTS registers (
          id INT(11) NOT NULL AUTO_INCREMENT,
          name VARCHAR(255) NOT NULL,
          quantity INT(11) NOT NULL,
          price DECIMAL(10,2) NOT NULL,
          date DATE NOT NULL,
          total_spent DOUBLE NOT NULL DEFAULT 0,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
          PRIMARY KEY (id)
      );
      
      -- Create 'users' table
      CREATE TABLE IF NOT EXISTS users (
          id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
          name VARCHAR(255) NOT NULL,
          mobile VARCHAR(20) NOT NULL,
          password VARCHAR(255) NOT NULL,
          role VARCHAR(50) NOT NULL,
          PRIMARY KEY (id)
      );
     ```

   - This will create a database called `stock_db` and the tables.

## Go Project Setup

1. **Clone the Repository**:
   - Clone this repository to your local machine using the following command:
     ```bash
     https://github.com/CaioEd/Stock-Management.git
     cd your-repository
     ```

2. **Configure Database Connection**:
   - Open the file where the database connection is defined in your Go project. This file is typically where the MySQL connection is initialized. Example configuration:
     ```go
     db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/stock_db")
     if err != nil {
         log.Fatal(err)
     }
     ```
   - Make sure the connection string contains the correct database name (`stock_db`) and the appropriate password (usually `root` for XAMPP unless modified).

3. **Install Dependencies**:
   - To ensure all required dependencies for the project are installed, run the following command:
     ```bash
     go mod tidy
     or
     make install (if you have Make installed)
     ```

## Running the Project

1. **Start the Go Application**:
   - After setting up the database and installing dependencies, run the project using:
     ```bash
     go run main.go
     or
     command: air
     ```

2. **Access the Application**:
   - Follow any additional instructions provided in the project to interact with the application, such as accessing an API endpoint or a frontend interface.

## Contribution

Feel free to open issues or create pull requests if you encounter any problems or have suggestions for improvement.

## License

This project is licensed under the [MIT License](LICENSE).

---

Follow these steps to set up and run the project successfully. 
