# Clean Architecture Challenge

Este projeto implementa um sistema de orders usando Clean Architecture, com suporte a múltiplos protocolos de comunicação (REST, gRPC e GraphQL).

## Portas dos Serviços

- REST API: http://localhost:8000
- gRPC: localhost:50051
- GraphQL: http://localhost:8080
- MySQL: localhost:3306
- RabbitMQ: localhost:5672 (Management: http://localhost:15672)

## Como Executar

1. Inicie os serviços do MySQL e RabbitMQ:
```bash
docker-compose up -d
```

2. Execute a aplicação:
```bash
go run cmd/ordersystem/main.go
```

## Endpoints Disponíveis

### REST API

- Criar Order: `POST http://localhost:8000/order`
- Listar Orders: `GET http://localhost:8000/order/list`

### GraphQL

- Endpoint: `http://localhost:8080/query`
- Playground: `http://localhost:8080`

Queries disponíveis:
```graphql
# Criar Order
mutation {
  createOrder(input: { id: "123", Price: 100.0, Tax: 10.0 }) {
    id
    Price
    Tax
    FinalPrice
  }
}

# Listar Orders
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

Use o `grpcurl` para testar os endpoints gRPC:

```bash
# Listar Orders
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

# Criar Order
grpcurl -plaintext -d '{"id": "123", "price": 100.0, "tax": 10.0}' localhost:50051 pb.OrderService/CreateOrder
```

## Estrutura do Projeto

O projeto segue os princípios da Clean Architecture:

- `cmd/ordersystem`: Ponto de entrada da aplicação
- `internal/entity`: Entidades do domínio
- `internal/usecase`: Casos de uso da aplicação
- `internal/infra`: Implementações de infraestrutura
  - `database`: Repositórios
  - `web`: Handlers REST
  - `grpc`: Serviços gRPC
  - `graph`: Resolvers GraphQL

## Tecnologias Utilizadas

- Go
- MySQL
- RabbitMQ
- gRPC
- GraphQL (gqlgen)
- Docker
- Wire (Dependency Injection) 