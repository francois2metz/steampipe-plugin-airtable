# Table: {airtable_table}

The `{airtable_table}` table can be used to query your Airtable table. A table is automatically created to represent each configured `tables`.

For instance, if `tables` is set to `["Design Projects", "Tasks", "Clients"]`, then this plugin will create 3 tables:

- design_projects
- tasks
- clients

## Examples

### Get all ids

```sql
select
  id
from
  design_projects;
```

### Get a record by id

```sql
select
  created_time,
  fields
from
  design_projects
where
  id='recdTpx4c0kPPDTtf';
```

### Get the 5 last created rows

```sql
select
  id
from
  design_projects
order by
  created_time desc
limit
  5;
```

### Join 2 tables

```sql
select
  d.fields->'Name' as name,
  c.fields->'Name' as client_name
from
  design_projects d
cross join lateral
  jsonb_array_elements(d.fields->'Client') j(client)
join
  clients c on c.id = j.client#>>'{}';
```
