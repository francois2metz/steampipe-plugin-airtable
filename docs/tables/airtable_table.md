# Table: airtable_table

List Airtable table from a specific database. You must specify which database in the where or join clause using the `base_id` column.

## Examples

### List all tables from a specific database

```sql
select
  id,
  name
from
  airtable_table
where
  base_id='appSAeVjtGBji7vNg';
```
