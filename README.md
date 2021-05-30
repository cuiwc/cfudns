# cfudns: Cloudflare client to update a single DNS record

This client updates a single DNS entry of a given DNS name in a given DNS zone on Cloudflare. It assumes that the DNS name has already been created with exactly one entry on Cloudflare. It is used for dyndns.

### Usage:

```
$ cfudns -token <Cloudflare API token> -zone <Zone Name> -dns-name <DNS Name> -dns-value <Updated DNS value>
```
or
```
$ export CF_TOKEN=<Cloudflare API token>
$ cfudns -zone <Zone Name> -dns-name <DNS Name> -dns-value <Updated DNS value>
```
