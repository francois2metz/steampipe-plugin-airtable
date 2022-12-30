---
organization: francois2metz
category: ["saas"]
brand_color: "#18BFFF"
display_name: "Airtable"
short_name: "Airtable"
description: "Steampipe plugin for querying Airtable."
og_description: "Query Airtable with SQL! Open source CLI. No DB required."
icon_url: "/images/plugins/francois2metz/airtable.svg"
og_image: "/images/plugins/francois2metz/airtable-social-graphic.png"
---

# Airtable + Steampipe

[Airtable](https://airtable.com/) is an easy-to-use online platform for creating and sharing relational databases.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  "Name"
from
  appid_design_projects
```

```
+--------------------+--------------------------+
| id                 | Name                     |
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

    # Personal Access Token (get it on https://airtable.com/create/tokens)
    # Scopes:
    # - data.records:read
    # - schema.bases:read
    # token = "pat.xxx"
}

```

You can also set the token via the `AIRTABLE_TOKEN` environment variable.

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-airtable
