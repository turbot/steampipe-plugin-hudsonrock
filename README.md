# Steampipe Plugin for Hudson Rock

Query Hudson Rock's cybercrime and infostealer intelligence data with SQL using Steampipe.

## Setup

1. Install [Steampipe](https://steampipe.io/downloads).
2. Clone this repo and build the plugin:
   ```sh
   make
   ```
3. Configure your connection in `~/.steampipe/config/hudsonrock.spc`:
   ```hcl
   connection "hudsonrock" {
     plugin = "hudsonrock"
     api_key = "YOUR_API_KEY"
   }
   ```

## Example Query

Search for compromised AWS credentials:

```sql
select
  file_name,
  date_compromised,
  data
from
  hudsonrock_file_search
where
  file_name = 'aws-credentials';
```

## Documentation

See the [Hudson Rock API docs](https://docs.hudsonrock.com/docs/file-search) for more details on available endpoints and parameters.