---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/hudsonrock.svg"
brand_color: "#ffc300"
display_name: Hudson Rock
name: hudsonrock
description: Steampipe plugin for querying domains, name servers and contact information from Hudson Rock.
og_description: Query Hudson Rock with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/hudsonrock-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# HudsonRock + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Hudson Rock](https://www.hudsonrock.com/) is an Israeli cybersecurity company specializing in cybercrime intelligence, with a focus on detecting and mitigating threats from Infostealer malware.

For example:

```sql
select
  domain,
  total_stealers
from
  hudsonrock_domain_search
where
  domain = 'steampipe.io';
```

```
+--------------+----------------+
| domain       | total_stealers |
+--------------+----------------+
| steampipe.io | 32,213,918     |
+--------------+----------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/hudsonrock/tables)**

## Get started

### Install

Download and install the latest Hudson Rock plugin:

```bash
steampipe plugin install hudsonrock
```

### Credentials

| Item | Description |
| - | - |
| Credentials | No creds required |
| Permissions | n/a |
| Radius | Steampipe connects to the correct Hudson Rock server based on the TLD |
| Resolution | n/a |

### Configuration

No configuration is needed. Installing the latest hudsonrock plugin will create a config file (`~/.steampipe/config/hudsonrock.spc`) with a single connection named `hudsonrock`:

```hcl
connection "hudsonrock" {
  plugin = "hudsonrock"
}
```


