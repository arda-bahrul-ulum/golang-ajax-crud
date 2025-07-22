# Golang AJAX CRUD

This is a simple CRUD (Create, Read, Update, Delete) application built with Golang, MySQL, jQuery AJAX, and the Bootstrap framework.

## Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/golang-ajax-crud.git
    cd golang-ajax-crud
    ```

2.  **Install dependencies:**
    This project uses Go Modules to manage dependencies. They will be downloaded automatically when you run the application.

## Libraries Used

- **Backend:**
  - [Go](https://golang.org/) - The core programming language.
  - [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go.
- **Frontend:**
  - [jQuery](https://jquery.com/) - For making AJAX requests.
  - [Bootstrap](https://getbootstrap.com/) - For styling the user interface.

## How to Run

1.  **Make sure you have Go installed** on your system.

2.  **Make sure you have a MySQL database** running and that you have updated the connection string in `config/database.go` with your database credentials.

3.  **Run the application:**

    ```bash
    go run main.go
    ```

4.  **Open your browser** and navigate to `http://localhost:4000` or `http://127.0.0.1:4000`.
