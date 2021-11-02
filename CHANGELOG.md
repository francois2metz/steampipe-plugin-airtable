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
- 404 errors are not handled
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

