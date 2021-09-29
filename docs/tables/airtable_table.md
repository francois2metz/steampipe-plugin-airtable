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

### Get a record by id

```sql
select
  created_time,
  fields
from
  airtable_design_projects
where
  id='recdTpx4c0kPPDTtf';
```

### Get the 5 last created rows

```sql
select
  id
from
  airtable_design_projects
order by
  created_time desc
limit
  5
```

### Join 2 tables

```sql
select
  d.fields->'Name' as name,
  c.fields->'Name' as client_name
from
  airtable_design_projects d
cross join lateral
  jsonb_array_elements(d.fields->'Client') j(client)
join
  airtable_clients c on c.id = j.client#>>'{}'
```
