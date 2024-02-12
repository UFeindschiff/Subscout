# Subscout - A tool for active subdomain enumeration, network mapping and asset discovery

## Subscout is a fork of [OWASP Amass](https://github.com/OWASP/Amass) primarily focussed on flexible, performant active subdomain enumeration
### Why fork Amass?
Mostly out of frustration. Unfortunately, the Amass maintainers seem less and less invested in the project. Long-standing performance- and stability-issues remain unaddressed and many merge requests to fix bugs or improve funtionality don't even have any comment by the maintainer despite being open for years.
### What version of Amass did you fork from?
Version 3.21.2
### What advantages does Subscout have over Amass?
* It doesn't crash on marshalling output like Amass sometimes does (race condition which remains unaddressed in Amass)
* It supports enumeration over Tor
* It is possible to set trusted resolvers in the config file
* Queries per second per resolver are configurable (Amass hardcodes these at 25 queries per second per resolver)
* SRV records are not queried by default as they may drastically increase the enumeration runtime for some scopes while most often not yielding any extra results
### I am not interested in subdomain enumeration, but rather want to discover additional root domains.
You're better off using Amass instead. No work has been done on everything but subdomain enumeration whatsoever
# Building
Clone this repository and run `go build ./cmd/subscout`
# Usage
See the [User's Guide](./doc/user_guide.md) for additional information.
