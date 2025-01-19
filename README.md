# MoORM

A lightweight, fluent MongoDB query builder for Go that simplifies the creation of MongoDB queries.

## Features

- Fluent interface for building MongoDB queries
- Type-safe query construction
- Support for common MongoDB operators
- Easy to read and maintain query syntax
- Works with the official MongoDB Go driver

## Installation

```bash
go get github.com/metajar/moorm
```

## Query Operations

MoORM supports the following MongoDB operators:

### Comparison
- `Eq()` - Equal to
- `Ne()` - Not equal to
- `Gt()` - Greater than
- `Gte()` - Greater than or equal
- `Lt()` - Less than
- `Lte()` - Less than or equal
- `In()` - Match any value in array
- `Nin()` - Not match any value in array

### String Matching
- `Like()` - Case-insensitive pattern matching
- `NotLike()` - Negated pattern matching
- `Regex()` - Regular expression matching

### Logical
- `And()` - Logical AND between conditions
- `Or()` - Logical OR between conditions

### Field
- `Exists()` - Check if field exists

## Basic Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/metajar/moorm/pkg/moorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.TODO()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("testing").Collection("test")

	// Insert a test document
	doc := moorm.M{
		"hostname": "br01.dfw02",
		"tags":     []string{"NA", "Cisco", "IosXR"},
		"ip":       "192.168.88.4",
		"updays":   10000,
		"active":   true,
		"issues":   19,
	}
	collection.InsertOne(ctx, doc)

	// Build a filter with multiple conditions
	filter := moorm.
		Filter().
		Like("hostname", "dfw02").
		Regex("ip", ".*168.*").
		Gte("updays", 999).
		Eq("active", true).
		Lte("issues", 20).
		In("tags", []interface{}{"Cisco", "NA"}).
		Nin("tags", []interface{}{"Juniper"}).
		Build()

	// Find and decode the result
	result := collection.FindOne(ctx, filter)
	record := make(map[string]interface{})
	if err := result.Decode(&record); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Record:", record)
}
```

## Logical Operations Example

```go
// Create multiple filters
filter1 := moorm.Filter().Eq("status", "active").Build()
filter2 := moorm.Filter().Gte("priority", 5).Build()
filter3 := moorm.Filter().Like("category", "urgent").Build()

// Combine filters with OR
orFilter := moorm.Or([]moorm.Document{filter1, filter2, filter3})

// Combine filters with AND
andFilter := moorm.And([]moorm.Document{filter1, filter2})
```

## Update Operations

MoORM also supports building update queries:

```go
update := moorm.
	Update().
	Unset("temporary_field").
	Build()
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT](LICENSE)