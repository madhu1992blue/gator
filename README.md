# gator

## Dependencies

- go 1.23.4+
- postgres 16.6+
- goose v3.24.0+
- sqlc 1.27.0+

## Installation:

1. Clone this source
2. Run `go install`


## Quick Start

Create .gatorconfig.json in home directory:  
Note: Modify connection string with DB connection string. Leave username blank
```
{"db_url":"","current_user_name":""}
```
Note: Modify connection string with DB connection string leave username blank

```bash
# Replace <yourname> with your name
gator register "<yourname>"
```

```bash
# Add a few feeds (This will automatically make you follow them)
gator addfeed "Hacker News RSS" "https://hnrss.org/newest"
gator addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"
```

```bash
# See what you follow
gator following
```

```bash
# Let's scrape every 10s. Let this run for about a minute or 2 and then press Ctrl+C to interrupt.
gator agg 10s
```

```bash
# View upto 5 latest posts. Change 5 to whatever max you like.
gator browse 5
```

## How-To
### Onboarding

1. Create .gatorconfig.json in home directory:
Note: Modify connection string with DB connection string leave username blank
```
{"db_url":"","current_user_name":""}
```
Note: Modify connection string with DB connection string leave username blank

### Registration and Login

1. Register (Replace <username> with your name)
```bash
gator register <username>
```
The above command will switch you to that user.

1. If you have multiple accounts, you can switch to a different user with:
```bash
gator login <username>
```

### Add and Follow a feed

#### Syntax
```bash
# Replace name and URL. See example.
gator addfeed <name> <url>
```
Note: This may fail if a feed with same URL already exists.

#### Example to add and follow a feed
```bash
# Example:
gator addfeed "Hacker News RSS" "https://hnrss.org/newest"
```

### View Existing Feeds

```bash
gator feeds
```

### Follow existing Feeds

```bash
gator follow <feedurl>
```

### View Following Feeds

```bash
gator following
```

### Unfollow existing Feeds

```
gator unfollow <feedurl>
```


### Scrape Feeds and Posts

```bash
# Replace duration actual duration like 1s, 20m, 1h, etc.
gator agg <duration>
```

### Browse saved Posts

```bash
# Specify numposts (else default is 2)
gator browse [numposts]
```

### Reset
```bash
gator reset
```