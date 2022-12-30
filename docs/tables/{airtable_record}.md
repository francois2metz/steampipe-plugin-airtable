# Table: {airtable_record}

Query data from Airtable tables. A table is automatically created for each tables on each bases available.

For instance, if you have 2 bases with 2 tables on each, then this plugin will create 4 tables:

- baseid1_table_name1
- baseid1_table_name2
- baseid2_table_name1
- baseid2_table_name2

## Examples

### Get all IDs

```sql
select
  id
from
  baseid_design_projects;
```

### Get a record by ID

```sql
select
  *
from
  baseid_design_projects
where
  id = 'recdTpx4c0kPPDTtf';
```

### Get the 5 last created rows

```sql
select
  id
from
  baseid_design_projects
order by
  created_time desc
limit
  5;
```

### Expand `Client` field to list clients associated with projects

```sql
select
  d.id as project_id,
  cid as client_id
from
  baseid_design_projects as d,
  jsonb_array_elements_text(d.fields -> 'Client') as cid;
```

### Get full information for each client associated with a project

```sql
select
  d.id as project_id,
  c.id as client_id,
  d."Name" as project_name,
  c."Name" as client_name,
  c."About" as client_description
from
  baseid_design_projects as d,
  jsonb_array_elements_text(d."Client") as cid,
  baseid_clients as c
where
  c.id = cid;
```

### List projects using the [formula filter](https://support.airtable.com/hc/en-us/articles/203255215)

```sql
select
  "Name",
  "Kickoff date"
from
  baseid_design_projects
where
  filter_formula = 'IS_AFTER({Kickoff date}, "2020-10-01")';
```
