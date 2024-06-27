# Parser

A blockchain parser that allows querying transactions for subscribed addresses. 
This implementation follows a clean architecture approach to facilitate future extension and maintenance.

## Usage example

#### Subscribe to an Address
To subscribe to an address, use the following command:
```shell
curl -X POST http://localhost:8080/apis/v1/subscribe --data '{"address": "0x5ea0b5899e912e0b0ba0190dd097251dbdaa70e0"}'
```

#### Get Transactions for a Specified Address

To retrieve all transactions for a specified address, use this command:
```shell
curl -X POST http://localhost:8080/apis/v1/transactions --data '{"address": "0x5ea0b5899e912e0b0ba0190dd097251dbdaa70e0"}'
```

#### Get Current Block

To obtain the current block, use the following command:
```shell
curl -X POST http://localhost:8080/apis/v1/block/current"
```

### Planned Improvements
- Add graceful shutdown for the application.
- Improve error handling.
- Enhance the transaction fetcher.
- Add websockets to send updates.