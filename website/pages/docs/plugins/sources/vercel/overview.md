# Vercel Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `vercel`)}/>

The CloudQuery Vercel plugin pulls configuration out of Vercel resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Vercel, `cloudquery` needs to be authenticated. An access token is required for authentication. You can create a new access token in [Vercel's Account Tokens Page](https://vercel.com/account/tokens).

## Query Examples

### Detect domain registrations that will expire soon

```sql
select name, expires_at, date_trunc('day', expires_at - current_timestamp) as days_left from vercel_domains where (expires_at - interval '90 day') < current_timestamp order by 1;
```

### Get all name servers in Vercel domains

```sql
select name, intended_nameservers, custom_nameservers, nameservers from vercel_domains order by 1;
```

### Get all Vercel team members

```sql
select t.name AS team, u.username, u.name, u.role from vercel_teams t join vercel_team_members u on u.team_id=t.id order by 1, 2;
```

### Get all Vercel projects

```sql
select name from vercel_projects order by 1;
```
### Get Vercel deployment statistics

```sql
select state, count(1) from vercel_deployments group by 1 order by 2 desc;
```
