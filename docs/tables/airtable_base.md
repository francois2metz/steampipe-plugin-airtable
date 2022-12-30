# Table: airtable_base

List all your Airtable bases.

## Examples

### List all bases

```sql
select
  id,
  name
from
  airtable_base;
```

### List base with only read access

```sql
select
  id,
  name
from
  airtable_base
where
  permission_level='read';
```
