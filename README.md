# Better enums

Adds helper methods to go protobuf enums that allow them to be used directly in 99designs/gqlgen and database/sql.

By default, the enum value will be used as the db & gql values, but both can be overridden via proto options e.g:

```proto
enum MyEnum {
    VALUE_1 = 0 [(benum.db)="db_value_1", (benum.gql)="gqlValue1"];
    VALUE_2 = 1 [(benum.db)="db_value_2", (benum.gql)="gqlValue2"];
    VALUE_3 = 2 [(benum.db)="db_value_3", (benum.gql)="gqlValue3"];
}
```
