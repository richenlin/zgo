
# ent使用

https://entgo.io/docs/getting-started/  

```sh
PATH=$PATH:/root/go/bin
go get github.com/facebookincubator/ent/cmd/entc
go mod init <project>
entc init User
# generate the schema for User under <project>/ent/schema/
```
``` golang
package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
)

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.String("name").
            Default("unknown"),
    }
}
```
```sh
go generate ./ent
#entc generate
```
```
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```
```golang
package main

import (
    "log"

    "<project>/ent"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    // run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```
```golang
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %v", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}
```
```golang
package main

import (
    "log"

    "<project>/ent"
    "<project>/ent/user"
)

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Query().
        Where(user.NameEQ("a8m")).
        // `Only` fails if no user found,
        // or more than 1 user returned.
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %v", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}
```
增加关联关系
```sh
entc init Car Group
```
```golang
import (
    "regexp"

    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
)

// Fields of the Car.
func (Car) Fields() []ent.Field {
    return []ent.Field{
        field.String("model"),
        field.Time("registered_at"),
    }
}


// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            // regexp validation for group name.
            Match(regexp.MustCompile("[a-zA-Z_]+$")),
    }
}
```
```golang
import (
   "log"

   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
)

// Edges of the User.
func (User) Edges() []ent.Edge {
   return []ent.Edge{
       edge.To("cars", Car.Type),
   }
}
```
