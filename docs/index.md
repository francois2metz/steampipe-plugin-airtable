---
organization: francois2metz
category: ["saas"]
brand_color: "#f82b60"
display_name: "Airtable"
short_name: "Airtable"
description: "Steampipe plugin for querying airtable."
og_description: "Query Airtable with SQL! Open source CLI. No DB required."
icon_url: "/images/plugins/francois2metz/airtable.svg"
---

# Airtable + Steampipe

[Airtable](https://airtable.com/) is an easy-to-use online platform for creating and sharing relational databases.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  fields->'Name' as name
from
  airtable_design_projects
```

```
+--------------------+--------------------------+
| id                 | name                     |
+--------------------+--------------------------+
| recHQSd02Tjhba3ue  | Coffee packaging         |
| rec1x6H2wuyJArcwM  | EngineerU brand identity |
+--------------------+--------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/airtable/tables)**

## Get started

### Install

Download and install the latest Airtable plugin:

```bash
steampipe plugin install francois2metz/airtable
```

### Configuration

Installing the latest airtable plugin will create a config file (`~/.steampipe/config/airtable.spc`) with a single connection named `airtable`:

```hcl
connection "airtable" {
  plugin = "francois2metz/airtable"

  # API token
  token = "YOUR_AIRTABLE_TOKEN"

  # Database ID
  databaseid = "YOUR_DATABASE_ID"

  # Tables to expose
  tables = []
}
```

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-airtable
