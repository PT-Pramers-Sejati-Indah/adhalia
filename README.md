# Adhalia

Adhalia is a lightweight `sqlx` wrapper designed to simplify the implementation of Command Query Responsibility Segregation (CQRS) in Go applications. This library helps you manage separate read and write databases, making it easier to maintain a scalable and maintainable architecture.

## Requirements
- Go 1.10

## Features

- **Separation of Read and Write Operations**: Adhalia provides a clear separation of read and write operations, allowing you to optimize your database interactions.
- **Transactional Support**: Ensure consistency and integrity with built-in support for transactions on write operations.
- **Ease of Use**: Simplifies the usage of `sqlx` with CQRS pattern in mind, providing a clean and straightforward API.
- **Security**: Utilizes prepared statements to help prevent SQL injection attacks.

## Installation

To install Adhalia, use the following command:

```bash
go get github.com/PT-Pramers-Sejati-Indah/adhalia
