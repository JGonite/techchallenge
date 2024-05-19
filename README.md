# Skina Lanches Management API

## Overview

The Skina Lanches Management API is designed to facilitate the management of products, categories, clients, and orders for a sandwich shop. The API provides endpoints to create, read, update, and delete entities within the system.

## Table of Contents

- [Skina Lanches Management API](#skina-lanches-management-api)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Setup](#setup)
    - [Docker Setup](#docker-setup)
  - [API Endpoints](#api-endpoints)
    - [Categories](#categories)
    - [Clients](#clients)
    - [Orders](#orders)
    - [Products](#products)

## Setup

### Docker Setup

To set up the application using Docker, follow these steps:

1. Ensure you have Docker installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Build the Docker image using the provided Dockerfile:
   ```sh
   docker build -t skina-lanches-api .
5. Run the Docker container:
    ```sh
    docker run -p 9090:9090 skina-lanches-api


## API Endpoints

### Categories

- **GET /categories**
  - Retrieves a paginated list of categories.
  - Parameters:
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of categories per page.
  - Responses:
    - `200`: Successfully retrieved list of categories.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /categories**
  - Adds a new category to the database.
  - Body: `dto.CreateCategoryRequest`
  - Responses:
    - `201`: Successfully created category.
    - `400`: Bad request if the category data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /categories/{id}**
  - Retrieves details of a category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Responses:
    - `200`: Successfully retrieved the category details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Category not found if the ID does not match any category.
    - `500`: Internal server error if there is a problem on the server side.

- **PUT /categories/{id}**
  - Updates category details by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Body: `dto.CreateCategoryRequest`
  - Responses:
    - `200`: Successfully updated category.
    - `400`: Invalid input, object is invalid.
    - `404`: Category not found.
    - `500`: Internal server error.

- **DELETE /categories/{id}**
  - Deletes a category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Responses:
    - `200`: Message indicating successful deletion.
    - `400`: Bad request if the ID is not provided or is invalid.
    - `404`: Category not found if the ID does not match any category.
    - `500`: Internal server error if there is a problem deleting the category.

### Clients

- **POST /clients**
  - Adds a new client to the database.
  - Body: `dto.CreateClientRequest`
  - Responses:
    - `201`: Client successfully created.
    - `400`: Bad request if the client data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /clients/{cpf}**
  - Retrieves details of a client by its CPF.
  - Parameters:
    - `cpf` (string): Client CPF.
  - Responses:
    - `200`: Successfully retrieved the client details.
    - `400`: Bad request if the CPF is not provided or invalid.
    - `404`: Client not found if the CPF does not match any client.
    - `500`: Internal server error if there is a problem on the server side.

### Orders

- **GET /orders**
  - Retrieves a paginated list of orders.
  - Parameters:
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of orders per page.
  - Responses:
    - `200`: Successfully retrieved list of orders.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /orders**
  - Adds a new order to the database.
  - Body: `dto.CreateOrderRequest`
  - Responses:
    - `201`: Successfully created order.
    - `400`: Bad request if the order data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /orders/{id}**
  - Retrieves details of an order by its ID.
  - Parameters:
    - `id` (string): Order ID.
  - Responses:
    - `200`: Successfully retrieved the order details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Order not found if the ID does not match any order.
    - `500`: Internal server error if there is a problem on the server side.

### Products

- **GET /products**
  - Retrieves a paginated list of products, optionally filtered by category.
  - Parameters:
    - `category` (string, optional): Filter by category ID.
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of products per page.
  - Responses:
    - `200`: Successfully retrieved list of products.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /products**
  - Adds a new product to the database.
  - Body: `dto.CreateProductRequest`
  - Responses:
    - `201`: Product successfully created.
    - `400`: Bad request if the product data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /products/{id}**
  - Retrieves details of a product by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Responses:
    - `200`: Successfully retrieved the product details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Product not found if the ID does not match any product.
    - `500`: Internal server error if there is a problem on the server side.

- **PUT /products/{id}**
  - Updates product details by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Body: `dto.CreateProductRequest`
  - Responses:
    - `200`: Product successfully updated.
    - `400`: Invalid input, object is invalid.
    - `404`: Product not found.
    - `500`: Internal server error.

- **DELETE /products/{id}**
  - Deletes a product by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Responses:
    - `200`: Message indicating successful deletion.
    - `400`: Bad request if the ID is not provided or is invalid.
    - `404`: Product not found if the ID does not match any product.
    - `500`: Internal server error if there is a problem deleting the product.
