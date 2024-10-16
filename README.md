# gRPC Microservices Project with GraphQL API


## Overview

This project demonstrates a microservices architecture using gRPC for inter-service communication and GraphQL as the API gateway. It includes services for account management, product catalog, and order processing.

## Project Structure

The project consists of the following main components:

- Account Service
- Catalog Service
- Order Service
- GraphQL API Gateway

Each service has its own database:
- Account and Order services use PostgreSQL
- Catalog service uses Elasticsearch

## Key Features

   - gRPC Communication.
   - GraphQL API Gateway.
   - gRPC Streaming.
   - Protobuf Serialization.
   - Dockerized Services.

## Technology Stack

- **Golang:** A high-performance language used for building efficient, scalable microservices.
- **gRPC:** A remote procedure call framework for fast, efficient communication between services.
- **GraphQL:** A flexible API query language that allows clients to request only the data they need.
- **Protocol Buffers (Protobuf):** A compact and fast serialization format used for data exchange in gRPC.
- **Docker:** A containerization platform that simplifies deployment and scaling of microservices.
- **PostgreSQL:** A powerful relational database known for its reliability, performance, and SQL compliance.
- **Elasticsearch:** A distributed search and analytics engine used for full-text search and log analysis.


## Getting Started

- Clone the repository:
   ```
   git clone <repository-url>
   cd <project-directory>
   ```

- Start the services using Docker Compose:
   ```
   docker-compose up -d --build
   ```

- Access the GraphQL playground at `http://localhost:8000/playground`


## GraphQL API Usage

The GraphQL API provides a unified interface to interact with all the microservices.

### Query Accounts

```graphql
query {
  accounts {
    id
    name
  }
}
```

### Create an Account

```graphql
mutation {
  createAccount(account: {name: "New Account"}) {
    id
    name
  }
}
```

### Query Products

```graphql
query {
  products {
    id
    name
    price
  }
}
```

### Create a Product

```graphql
mutation {
  createProduct(product: {name: "New Product", description: "A new product", price: 19.99}) {
    id
    name
    price
  }
}
```

### Create an Order

```graphql
mutation {
  createOrder(order: {accountId: "account_id", products: [{id: "product_id", quantity: 2}]}) {
    id
    totalPrice
    products {
      name
      quantity
    }
  }
}
```

### Query Account with Orders

```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        name
        quantity
        price
      }
    }
  }
}
```

## Advanced Queries

### Pagination and Filtering

```graphql
query {
  products(pagination: {skip: 0, take: 5}, query: "search_term") {
    id
    name
    description
    price
  }
}
```

### Calculate Total Spent by an Account

```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      totalPrice
    }
  }
}
```

## Contributing

Contributions are always welcome!Please fork the repository and submit a pull request.
