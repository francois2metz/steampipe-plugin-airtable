# Table: {airtable_table}

Query data from Airtable tables. A table is automatically created to represent each Airtable table found in the configured `tables`.

For instance, if `tables` is set to `["Design Projects", "Tasks", "Clients"]`, then this plugin will create 3 tables:

- design_projects
- tasks
- clients

## Examples

### Get all IDs

```sql
select
  id
from
  design_projects;
```

### Get a record by ID

```sql
select
  created_time,
  fields
from
  design_projects
where
  id = 'recdTpx4c0kPPDTtf';
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

### Expand JSON field data and join with another table

```sql
select
  bi.id as bug_id,
  tm.id as team_member_id,
  tm.fields ->> 'Name' as assigned_to,
  bi.fields ->> 'Description' as description
from
  bugs_and_issues as bi,
  jsonb_array_elements_text(fields -> 'Assigned to') as a,
  team_members as tm
order by
  bug_id;
```
