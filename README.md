# Airtable plugin for Steampipe

Use SQL to query your [Airtable][] tables.

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/airtable

## Development

To buile the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/airtable.spc ~/.steampipe/config/airtable.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[airtable]: https://airtable.com
