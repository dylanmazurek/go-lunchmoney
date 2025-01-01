# go-lunchmoney

## Overview

`go-lunchmoney` is a Go client library for interacting with the Lunch Money API. It provides a convenient way to access and manage your financial data stored in Lunch Money.

## Installation

To install the library, use the following command:

```sh
go get github.com/dylanmazurek/go-lunchmoney
```

## Usage

### Importing the Library

```go
import (
    "context"
    "github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney"
)
```

### Creating a Client

```go
ctx := context.Background()
client, err := lunchmoney.New(ctx)
if err != nil {
    panic(err)
}
```

### Fetching an Asset

```go
assetID := int64(12345)
asset, err := client.FetchAsset(assetID)
if err != nil {
    panic(err)
}
fmt.Println("Asset:", asset)
```

### Listing Transactions

```go
startDate := "2023-01-01"
endDate := "2023-01-31"
filter := lunchmoney.ListTransactionFilter{
    StartDate: startDate,
    EndDate:   endDate,
}
transactions, err := client.ListTransaction(filter)
if err != nil {
    panic(err)
}
fmt.Println("Transactions:", transactions)
```

### Inserting Transactions

```go
transactions := []lunchmoney.Transaction{
    {
        Payee:        "Payee 1",
        Amount:       1000,
        Date:         "2023-01-01",
        CategoryName: "Category 1",
    },
    {
        Payee:        "Payee 2",
        Amount:       2000,
        Date:         "2023-01-02",
        CategoryName: "Category 2",
    },
}
ids, err := client.InsertTransactions(transactions, true)
if err != nil {
    panic(err)
}
fmt.Println("Inserted Transaction IDs:", ids)
```

### Updating a Transaction

```go
transaction := lunchmoney.Transaction{
    ID:           12345,
    Payee:        "Updated Payee",
    Amount:       1500,
    Date:         "2023-01-01",
    CategoryName: "Updated Category",
}
updated, err := client.UpdateTransaction(transaction, true)
if err != nil {
    panic(err)
}
fmt.Println("Transaction Updated:", updated)
```

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
