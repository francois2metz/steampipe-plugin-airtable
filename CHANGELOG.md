## v0.5.0 [2023-10-15]

_What's new?_

* Update SDK to v5.6.2
* Update go to 1.21

## v0.4.0 [2023-01-06]

New token, new config, new queries, this new release break the the compatibility with the previous versions as it's use the new meta API from Airtable.

_**Breaking changes**_
- The new SDK support only the new personal access token: https://airtable.com/create/tokens with the scopes *data.records:read* and *schema.bases:read*.
- The config fields *database_id* and *tables* have been removed
- The plugin expose now all tables from all bases by default. To access the records, the table is now named {{databaseid}}_{{tablename}}.
- The column fields on *airtable_record* table has been removed. You can now use the same columns name as in Airtable (with the correct type). If one field match an existing one, it is prefixed by an underscore.

_What's new?_

* Update SDK to v5
* Add table *airtable_base*
* Add table *airtable_tables*

## v0.3.0 [2022-09-01]

_What's new?_

* Update SDK to 4.1.5
* Update to go 1.19

## v0.2.1 [2022-05-25]

_What's new?_

* Add error logs

## v0.2.0 [2022-05-04]

_What's new?_

* Update SDK to 3.1.0
* Update to go 1.18
* Add arm64 support

## v0.1.4 [2022-01-24]

_What's new?_

* Update SDK to 1.8.3
* Add link on how to get the database id
* Allow to the set the airtable token via the `AIRTABLE_TOKEN` variable

## v0.1.3 [2021-11-24]

_What's new?_

* Update SDK to 1.8.2

## v0.1.2 [2021-11-05]

_What's new?_

* Update SDK to 1.7.2

## v0.1.1 [2021-11-02]

_What's new?_

* Fixed the number of results by default that was maximum 100

## v0.1.0 [2021-10-24]

_What's new?_

- **Breaking change**: rename the `query` column to `filter_formula`
* Added examples

## v0.0.5 [2021-10-18]

_What's new?_

- **Breaking change**: the `airtable_` prefix on tables has been dropped

## v0.0.4 [2021-10-18]

_What's new?_

- Update sdk to 1.7.0
- Set the schema mode to dynamic
- Update doc

## v0.0.3 [2021-10-17]

_What's new?_

- **Breaking change**: the *databaseid* config param has been renamed to *database_id*
- 404 errors are now handled
- Add the *query* column to use the *filterByFormula* param
- Limit results when limit parameter < 100

## v0.0.2 [2021-09-29]

_What's new?_

- The `created_time` column has been added
- Querying for one recordID is now much more faster
- API Queries are cancelled when the limit has been reached or the query has been cancelled

## v0.0.1 [2021-09-28]

_What's new?_

- Initial release with the generic table to query any of your Airtable tables

