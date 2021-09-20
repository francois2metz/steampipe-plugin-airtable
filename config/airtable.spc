connection "airtable" {
    plugin = "francois2metz/airtable"

    # API token (get it on https://airtable.com/account)
    token = "YOUR_AIRTABLE_TOKEN"

    # Database ID (get it by clicking "help -> api documentation". The Base ID is in the URL and in all of the examples).
    databaseid = "YOUR_DATABASE_ID"

    # Tables to expose
    tables = []
}
