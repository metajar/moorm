# MoORM

----
Simply build filters.

### Example

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
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.88.134:27999"))
	collection := client.Database("testing").Collection("test")
	doc := moorm.M{
		"hostname": "br01.dfw02",
		"tags":     []string{"NA", "Cisco", "IosXR"},
		"ip":       "192.168.88.4",
		"updays":   10000,
		"active":   true,
		"issues":   19,
	}
	collection.InsertOne(ctx, doc)

	// filter selector.
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
	fmt.Println("Filter:", filter)
	r := collection.FindOne(ctx, filter)

	record := make(map[string]interface{})
	err := r.Decode(&record)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Record:", record)
}

```

### Output

```bash
moorm go run main.go
Filter: map[active:map[$eq:true] hostname:map[$options:i $regex:.*dfw02.*] ip:map[$regex:.*168.*] issues:map[$lte:20] tags:map[$in:[Cisco NA] $nin:[Juniper]] updays:map[$gte:999]]
Record: map[_id:ObjectID("620470b873c862741ec17e38") active:true hostname:br01.dfw02 ip:192.168.88.4 issues:19 tags:[NA Cisco IosXR] updays:10000]


```