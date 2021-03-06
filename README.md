# WebMap

A Go tool to create a map of the web. The main goal is to collect all registered 
domains. The tool hast multiple approaches to collect these domains. The first one is
to generate domains. It then checks if these domains are registered.

### Start database

```
# Required for database
$ docker-compose up -d
```

### Build & Run

```
$ go build -o wemp

# Generate domains
$ ./wemp generate domain

# Fetch domains
$ ./wemp fetch domain
```

### Generated Domains

Within a generator the program will create random strings for domain names (e. g. a.de,
aa.de, aaa.de, b.de, ...).

### Crawl Search Engines

Todo: Query search engines with words from dictionaries, e.g. House, Car, Door, ... to 
receive result pages with domains.  

### Crawl websites for Domains

Todo: If a domain exists, we can crawl this website for more domains.

### Domain Checks

To check if a domain is registered there are multiple ways.
- Check if DNS settings like A, AAAA or SOA record are set
- Crawl a domain seller