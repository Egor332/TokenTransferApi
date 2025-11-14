# Features

The service provides a single GraphQL mutation for handling token transfers.

# GraphQL Schema

The available mutation is defined as:
```sh
type Mutation {
  transfer(from_address: String!, to_address: String!, amount: Int!): Int!
}
```

The transfer mutation takes the sender's address, the receiver's address, and the amount to transfer. It returns the new balance of the sender's (from_address) wallet upon a successful operation.

# Setup and Running

1. **Prerequisites**

You must have Docker and Docker Compose installed to manage the services (Go server and PostgreSQL database).

2. **Environment Configuration**

Create a file named .env in the root directory. Use the structure below as a guide:

.env.example
```sh
POSTGRES_USER=my_db_user
POSTGRES_PASSWORD=my_strong_password
POSTGRES_DB=my_app_db
```

**IMPORTANT**: Ensure your actual .env file is completed with the credentials you want to use.

3. **Build and Run**

Execute the following command to build the Go server container and start both the application and the database services:
```sh
docker-compose -f docker-compose.yml up --build
```

# Database Details

The service utilizes a PostgreSQL database.

The database contains one table named wallets:
```sh
CREATE TABLE wallets (
    wallet_address VARCHAR(42) PRIMARY KEY,
    balance BIGINT NOT NULL
);
```

# Initial Data Seeding

For testing and demonstration, two initial wallets are pre-loaded with a balance of 1,000,000 tokens:
```sh
INSERT INTO wallets (wallet_address, balance)
VALUES ('0x0000000000000000000000000000000000000000', 1000000);

INSERT INTO wallets (wallet_address, balance)
VALUES ('0x0000000000000000000000000000000000000001', 1000000);
```

# Manual Database Insertion

If you need to manually insert a new wallet into the running database, use the docker compose exec command (replace {YOUR_USER} and {YOUR_DB} with your .env values):
```sh
docker compose exec db psql -U {YOUR_USER} -d {YOUR_DB} -c "INSERT INTO wallets (wallet_address, balance) VALUES ('0x0000000000000000000000000000000000000002', 500000);"
```

# GraphQL Usage

You can test the API using a GraphQL playgraund available at http://localhost:8080/.

Example Mutation: 

This mutation demonstrates a transfer of 1 unit from the first seeded wallet to the second:
```sh
mutation TransferTest {
  transfer(
    from_address: "0x0000000000000000000000000000000000000000",
    to_address: "0x0000000000000000000000000000000000000001",
    amount: 1
  )
}
```
# Tests
You can run test from separete docker compose file. To succesfully run tests .env file is required (you can use the same .env as for application). To run tests run in your terminal:
```sh
docker compose -f docker-compose.test.yml up --build --abort-on-container-exit
```
Test results would be printed in your terminal.

