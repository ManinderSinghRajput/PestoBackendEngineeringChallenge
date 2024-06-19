# Backend Engineering Challenge: E-Commerce Microservices System

## Table of Contents
1. [Introduction](#introduction)
2. [System Architecture](#system-architecture)
3. [Technology Stack](#technology-stack)
4. [Service Breakdown](#service-breakdown)
5. [Concurrency Control](#concurrency-control)
6. [Clustering and High Availability](#clustering-and-high-availability)
7. [Database Schema](#database-schema)
8. [APIs and Communication](#apis-and-communication)
9. [Authentication and Authorization](#authentication-and-authorization)
10. [General Requirements](#general-requirements)
11. [Bonus Features](#bonus-features)

## Introduction

In this assignment, we will build a microservices-based e-commerce system using Go (Golang) for the backend. The system will handle user authentication, product management, and order processing, with an emphasis on concurrency control and high availability.

## System Architecture

The system will be structured into three main microservices:
1. **User Service**: Manages user authentication and authorization.
2. **Product Service**: Handles product CRUD operations.
3. **Order Service**: Manages order processing and history.

These services will communicate with each other via RESTful APIs and will be designed to run on a clustered environment to ensure high availability.

## Technology Stack

- **Programming Language**: Go (Golang)
- **Framework**: Gin (HTTP web framework for Go)
- **Database**: PostgreSQL
- **Concurrency Control**: Optimistic locking with `ETag` headers
- **Clustering**: Docker Swarm / Kubernetes
- **Authentication**: JSON Web Tokens (JWT)
- **Version Control**: Git

## Service Breakdown

1. **User Service**
    - **Endpoints**:
        - POST `/register`: Register a new user.
        - POST `/login`: Authenticate a user and issue a JWT.
        - GET `/profile`: Retrieve user profile (protected endpoint).
    - **Database Schema**:
        - `users` table: `id`, `username`, `password_hash`, `email`, `created_at`, `updated_at`

2. **Product Service**
    - **Endpoints**:
        - GET `/products`: List all products.
        - POST `/products`: Create a new product (protected endpoint).
        - PUT `/products/{id}`: Update a product (protected endpoint).
        - DELETE `/products/{id}`: Delete a product (protected endpoint).
    - **Concurrency Control**: Implement optimistic locking using `ETag` headers.
    - **Database Schema**:
        - `products` table: `id`, `name`, `description`, `price`, `stock`, `version`, `created_at`, `updated_at`

3. **Order Service**
    - **Endpoints**:
        - GET `/orders`: List all orders for the authenticated user (protected endpoint).
        - POST `/orders`: Create a new order (protected endpoint).
        - GET `/orders/{id}`: Retrieve a specific order (protected endpoint).
    - **Database Schema**:
        - `orders` table: `id`, `user_id`, `product_id`, `quantity`, `total_price`, `status`, `created_at`, `updated_at`

## Concurrency Control

In the Product Service, concurrency control will be implemented using optimistic locking. Each product record will have a `version` field that increments with each update. When updating a product, the service will check that the `version` in the request matches the current `version` in the database. If they match, the update proceeds and the `version` is incremented. If not, a conflict error is returned.

## Clustering and High Availability

To ensure high availability, the system will be deployed on a clustered environment using Docker Swarm or Kubernetes. Each microservice will run in its own container, and multiple instances of each service will be deployed. Service discovery and load balancing will be handled by the orchestration tool.

## Database Schema

**User Service**:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
**Product Service**:
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
**Order Service**
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
```

## APIs and Communication
Each service will expose a set of RESTful endpoints. Services will communicate via HTTP requests, and authentication will be handled using JWT tokens passed in the Authorization header.

## Authentication and Authorization
Authentication will be implemented using JWT. Users will authenticate via the User Service, which will issue a JWT upon successful login. This token will be used to access protected endpoints in the Product and Order Services.

## General Requirements
- **Code Quality**: The code will be clean, well-documented, and maintainable. Proper naming conventions, modularity, and best practices will be followed.
- **Version Control**: All code will be tracked using Git, with a repository hosted on GitHub.
- **Error Handling**: Comprehensive error handling and logging will be implemented across all services.
- **Testing**: Unit tests will be written for critical components, including the concurrency control mechanisms.
- **Deployment**: Detailed instructions will be provided for setting up and running the services on a local development environment or a cloud platform.

## Bonus Features
- **API Rate Limiting**: Implemented using a middleware to limit the number of requests per user.
- **Message Queues**: Used for asynchronous communication between services, such as order processing.
- **Caching**: Implemented using Redis to cache frequently accessed data and improve performance.
- **Monitoring and Alerting**: Integrated with Prometheus and Grafana for real-time monitoring and alerting.

