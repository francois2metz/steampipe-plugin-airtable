# Table: `airtable_<table>`

The `airtable_<table>` table can be used to query your Airtable table.

the `<table>` is dynamically generated from your config `tables`.

## Examples

### Get all ids

```sql
select
  id
from
  airtable_design_projects;
```
