connection "airtable" {
    plugin = "francois2metz/airtable"

    # API token (get it on https://airtable.com/account)
    token = "YOUR_AIRTABLE_TOKEN"

    # Database ID (get it by clicking "help -> api documentation". The Base ID is in the URL and in all of the examples).
    # See https://support.airtable.com/hc/en-us/articles/4405741487383-Understanding-Airtable-IDs#h_01FCC2DH96SWG2TQXF4YCXRYB7
    database_id = "YOUR_DATABASE_ID"

    # Tables to expose
    # For instance: ["Design Projects", "Tasks", "Client"]
    tables = []
}
