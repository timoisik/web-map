# WebMap

A Go tool to create a map of the web. The main goal is to collect all registered 
domains. The tool hast multiple approaches to collect these domains. The first one is
to generate domains. It then checks if these domains are registered.

### Generated Domains

Within a generator the program will create random strings for domain names (e. g. a.de,
aa.de, aaa.de, b.de, ...).

### Crawl Domains

Todo

### Domain Checks

To check if a domain is registered there are multiple ways.
- Check if DNS settings like A, AAAA or SOA record are set
- Crawl a domain seller