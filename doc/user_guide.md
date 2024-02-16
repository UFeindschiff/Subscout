# Subscout - Users' Guide

----

## Simple Examples For Getting Started

The subscout tool and all the subcommands show options using the **'-h'** and **'-help'** flags:

```bash
subscout -help
```

Check the version by performing the following:

```bash
subscout -version
```

The most basic use of the tool for subdomain enumeration:

```bash
subscout enum -d example.com
```

Typical parameters for DNS enumeration:

```bash
$ subscout enum -v -src -ip -brute -min-for-recursive 2 -d example.com
[Google] www.example.com
[VirusTotal] ns.example.com
...
```

The volume argument allows the subscout graph database to persist between executions and output files to be accessed on the host system. The first field (left of the colon) of the volume option is the subscout output directory that is external to Docker, while the second field is the path, internal to Docker, where subscout will write the output files.

## Command-line Usage Information

The subscout tool has several subcommands shown below for handling your Internet exposure investigation.

| Subcommand | Description |
|------------|-------------|
| intel | Collect open source intelligence for investigation of the target organization |
| enum | Perform DNS enumeration and network mapping of systems exposed to the Internet |
| viz | Generate visualizations of enumerations for exploratory analysis |
| track | Compare results of enumerations against common target organizations |
| db | Manage the graph databases storing the enumeration results |

All subcommands have some default global arguments that can be seen below.

| Flag | Description | Example |
|------|-------------|---------|
| -h/-help | Show the program usage message | subscout subcommand -h |
| -config | Path to the INI configuration file | subscout subcommand -config config.ini |
| -dir | Path to the directory containing the graph database | subscout subcommand -dir PATH -d example.com |
| -nocolor | Disable colorized output | subscout subcommand -nocolor -d example.com |
| -silent | Disable all output during execution | subscout subcommand -silent -json out.json -d example.com |

Please note that it is NOT possible to run certain subcommands or multiple instances thereof in parallel on the same (default) graph database. If you wish to run multiple subcommands or multiple instances of the same subcommand in parallel, make sure to specify a seperate graph database folder for each using the -dir flag.

Each subcommand's own arguments are shown in the following sections.

### The 'intel' Subcommand

The intel subcommand can help you discover additional root domain names associated with the organization you are investigating. The data source sections of the configuration file are utilized by this subcommand in order to obtain passive intelligence, such as reverse whois information.

| Flag | Description | Example |
|------|-------------|---------|
| -active | Enable active recon methods | subscout intel -active -addr 192.168.2.1-64 -p 80,443,8080 |
| -addr | IPs and ranges (192.168.1.1-254) separated by commas | subscout intel -addr 192.168.2.1-64 |
| -asn | ASNs separated by commas (can be used multiple times) | subscout intel -asn 13374,14618 |
| -cidr | CIDRs separated by commas (can be used multiple times) | subscout intel -cidr 104.154.0.0/15 |
| -d | Domain names separated by commas (can be used multiple times) | subscout intel -whois -d example.com |
| -demo | Censor output to make it suitable for demonstrations | subscout intel -demo -whois -d example.com |
| -df | Path to a file providing root domain names | subscout intel -whois -df domains.txt |
| -ef | Path to a file providing data sources to exclude | subscout intel -whois -ef exclude.txt -d example.com |
| -exclude | Data source names separated by commas to be excluded | subscout intel -whois -exclude crtsh -d example.com |
| -if | Path to a file providing data sources to include | subscout intel -whois -if include.txt -d example.com |
| -include | Data source names separated by commas to be included | subscout intel -whois -include crtsh -d example.com |
| -ip | Show the IP addresses for discovered names | subscout intel -ip -whois -d example.com |
| -ipv4 | Show the IPv4 addresses for discovered names | subscout intel -ipv4 -whois -d example.com |
| -ipv6 | Show the IPv6 addresses for discovered names | subscout intel -ipv6 -whois -d example.com |
| -list | Print the names of all available data sources | subscout intel -list |
| -log | Path to the log file where errors will be written | subscout intel -log subscout.log -whois -d example.com |
| -max-dns-queries | Maximum number of concurrent DNS queries | subscout intel -max-dns-queries 200 -whois -d example.com |
| -o | Path to the text output file | subscout intel -o out.txt -whois -d example.com |
| -org | Search string provided against AS description information | subscout intel -org Facebook |
| -p | Ports separated by commas (default: 80, 443) | subscout intel -cidr 104.154.0.0/15 -p 443,8080 |
| -r | IP addresses of preferred DNS resolvers (can be used multiple times) | subscout intel -r 8.8.8.8,1.1.1.1 -whois -d example.com |
| -rf | Path to a file providing preferred DNS resolvers | subscout intel -rf data/resolvers.txt -whois -d example.com |
| -src | Print data sources for the discovered names | subscout intel -src -whois -d example.com |
| -timeout | Number of minutes to execute the enumeration | subscout intel -timeout 30 -d example.com |
| -v | Output status / debug / troubleshooting info | subscout intel -v -whois -d example.com |
| -whois | All discovered domains are run through reverse whois | subscout intel -whois -d example.com |

### The 'enum' Subcommand

This subcommand will perform DNS enumeration and network mapping while populating the selected graph database. All the setting available in the configuration file are relevant to this subcommand. Please note that parrallel enumerations on the same default graph-DB are NOT possible. If you wish to do parallel enumerations, make sure to specify a seperate graph database folder using the -dir flag. The following flags are available for configuration:

| Flag | Description | Example |
|------|-------------|---------|
| -active | Enable active recon methods | subscout enum -active -d example.com -p 80,443,8080 |
| -alts | Enable generation of altered names | subscout enum -alts -d example.com |
| -aw | Path to a different wordlist file for alterations | subscout enum -aw PATH -d example.com |
| -awm | "hashcat-style" wordlist masks for name alterations | subscout enum -awm dev?d -d example.com |
| -bl | Blacklist of subdomain names that will not be investigated | subscout enum -bl blah.example.com -d example.com |
| -blf | Path to a file providing blacklisted subdomains | subscout enum -blf data/blacklist.txt -d example.com |
| -brute | Perform brute force subdomain enumeration | subscout enum -brute -d example.com |
| -d | Domain names separated by commas (can be used multiple times) | subscout enum -d example.com |
| -demo | Censor output to make it suitable for demonstrations | subscout enum -demo -d example.com |
| -df | Path to a file providing root domain names | subscout enum -df domains.txt |
| -dns-qps | Maximum number of DNS queries per second across all resolvers | subscout enum -dns-qps 200 -d example.com |
| -dosrvlookup | Do SRV (service name) lookup during subdomain enumeration. This may rarely yield more results, but may also drastically increase runtime | subscout enum -dosrvlookup -d example.com |
| -ef | Path to a file providing data sources to exclude | subscout enum -ef exclude.txt -d example.com |
| -exclude | Data source names separated by commas to be excluded | subscout enum -exclude crtsh -d example.com |
| -if | Path to a file providing data sources to include | subscout enum -if include.txt -d example.com |
| -iface | Provide the network interface to send traffic through | subscout enum -iface en0 -d example.com |
| -include | Data source names separated by commas to be included | subscout enum -include crtsh -d example.com |
| -ip | Show the IP addresses for discovered names | subscout enum -ip -d example.com |
| -ipv4 | Show the IPv4 addresses for discovered names | subscout enum -ipv4 -d example.com |
| -ipv6 | Show the IPv6 addresses for discovered names | subscout enum -ipv6 -d example.com |
| -json | Path to the JSON output file | subscout enum -json out.json -d example.com |
| -list | Print the names of all available data sources | subscout enum -list |
| -log | Path to the log file where errors will be written | subscout enum -log subscout.log -d example.com |
| -max-depth | Maximum number of subdomain labels for brute forcing | subscout enum -brute -max-depth 3 -d example.com |
| -max-dns-queries | Deprecated flag to be replaced by dns-qps in version 4.0 | subscout enum -max-dns-queries 200 -d example.com |
| -min-for-recursive | Subdomain labels seen before recursive brute forcing (Default: 1) | subscout enum -brute -min-for-recursive 3 -d example.com |
| -nf | Path to a file providing already known subdomain names (from other tools/sources) | subscout enum -nf names.txt -d example.com |
| -nordns | Disables reverse DNS lookups. This decreases the result quality, but may drastically lower the runtime for certain scopes such as universities | subscout enum -nordns -d example.com |
| -norecursive | Turn off recursive brute forcing | subscout enum -brute -norecursive -d example.com |
| -o | Path to the text output file | subscout enum -o out.txt -d example.com |
| -oA | Path prefix used for naming all output files | subscout enum -oA subscout_scan -d example.com |
| -p | Ports separated by commas (default: 443) | subscout enum -d example.com -p 443,8080 |
| -passive | A purely passive mode of execution | subscout enum --passive -d example.com |
| -r | IP addresses of untrusted DNS resolvers (can be used multiple times) | subscout enum -r 8.8.8.8,1.1.1.1 -d example.com |
| -rf | Path to a file providing untrusted DNS resolvers | subscout enum -rf data/resolvers.txt -d example.com |
| -rqps | Maximum number of DNS queries per second for each untrusted resolver | subscout enum -rqps 10 -d example.com |
| -scripts | Path to a directory containing ADS scripts | subscout enum -scripts PATH -d example.com |
| -src | Print data sources for the discovered names | subscout enum -src -d example.com |
| -timeout | Number of minutes to execute the enumeration | subscout enum -timeout 30 -d example.com |
| -tordns | Run in compatibility mode to allow subdomain enumeration over Tor (disables all lookups except for A and AAAA) | subscout enum -tordns -d example.com |
| -tr | IP addresses of trusted DNS resolvers (can be used multiple times) | subscout enum -tr 8.8.8.8,1.1.1.1 -d example.com |
| -trf | Path to a file providing trusted DNS resolvers | subscout enum -trf data/trusted.txt -d example.com |
| -trqps | Maximum number of DNS queries per second for each trusted resolver | subscout enum -trqps 20 -d example.com |
| -v | Output status / debug / troubleshooting info | subscout enum -v -d example.com |
| -w | Path to a different wordlist file for brute forcing | subscout enum -brute -w wordlist.txt -d example.com |
| -wm | "hashcat-style" wordlist masks for DNS brute forcing | subscout enum -brute -wm ?l?l -d example.com |


### The 'viz' Subcommand

Create enlightening network graph visualizations that add structure to the information gathered. This subcommand only leverages the 'output_directory' and remote graph database settings from the configuration file.

The files generated for visualization are created in the current working directory and named subscout_TYPE

Switches for outputting the DNS and infrastructure findings as a network graph:

| Flag | Description | Example |
|------|-------------|---------|
| -d | Domain names separated by commas (can be used multiple times) | subscout viz -d3 -d example.com |
| -d3 | Output a D3.js v4 force simulation HTML file | subscout viz -d3 -d example.com |
| -df | Path to a file providing root domain names | subscout viz -d3 -df domains.txt |
| -dot | Generate the DOT output file | subscout viz -dot -d example.com |
| -enum | Identify an enumeration via an index from the db listing | subscout viz -enum 1 -d3 -d example.com |
| -gexf | Output to Graph Exchange XML Format (GEXF) | subscout viz -gexf -d example.com |
| -graphistry | Output Graphistry JSON | subscout viz -graphistry -d example.com |
| -i | Path to the subscout data operations JSON input file | subscout viz -d3 -d example.com |
| -maltego | Output a Maltego Graph Table CSV file | subscout viz -maltego -d example.com |
| -o | Path to a pre-existing directory that will hold output files | subscout viz -d3 -o OUTPATH -d example.com |
| -oA | Prefix used for naming all output files | subscout viz -d3 -oA example -d example.com |

### The 'track' Subcommand

Shows differences between enumerations that included the same target(s) for monitoring a target's attack surface. This subcommand only leverages the 'output_directory' and remote graph database settings from the configuration file. Flags for performing Internet exposure monitoring across the enumerations in the graph database:

| Flag | Description | Example |
|------|-------------|---------|
| -d | Domain names separated by commas (can be used multiple times) | subscout track -d example.com |
| -df | Path to a file providing root domain names | subscout track -df domains.txt |
| -history | Show the difference between all enumeration pairs | subscout track -history |
| -last | The number of recent enumerations to include in the tracking | subscout track -last NUM |
| -since | Exclude all enumerations before a specified date (format: 01/02 15:04:05 2006 MST) | subscout track -since DATE |

### The 'db' Subcommand

Performs viewing and manipulation of the graph database. This subcommand only leverages the 'output_directory' and remote graph database settings from the configuration file. Flags for interacting with the enumeration findings in the graph database include:

| Flag | Description | Example |
|------|-------------|---------|
| -d | Domain names separated by commas (can be used multiple times) | subscout db -d example.com |
| -demo | Censor output to make it suitable for demonstrations | subscout db -demo -d example.com |
| -df | Path to a file providing root domain names | subscout db -df domains.txt |
| -enum | Identify an enumeration via an index from the listing | subscout db -enum 1 -show |
| -ip | Show the IP addresses for discovered names | subscout db -show -ip -d example.com |
| -ipv4 | Show the IPv4 addresses for discovered names | subscout db -show -ipv4 -d example.com |
| -ipv6 | Show the IPv6 addresses for discovered names | subscout db -show -ipv6 -d example.com |
| -json | Path to the JSON output file or '-' | subscout db -names -silent -json out.json -d example.com |
| -list | Print enumerations in the database and filter on domains specified | subscout db -list |
| -names | Print just discovered names | subscout db -names -d example.com |
| -o | Path to the text output file | subscout db -names -o out.txt -d example.com |
| -show | Print the results for the enumeration index + domains provided | subscout db -show |
| -src | Print data sources for the discovered names | subscout db -show -src -d example.com |
| -summary | Print just ASN table summary | subscout db -summary -d example.com |

## The Output Directory

subscout has several files that it outputs during an enumeration (e.g. the log file). If you are not using a database server to store the network graph information, then subscout creates a file based graph database in the output directory. These files are used again during future enumerations, and when leveraging features like tracking and visualization.

By default, the output directory is created in the operating system default root directory to use for user-specific configuration data and named *subscout*. If this is not suitable for your needs, then the subcommands can be instructed to create the output directory in an alternative location using the **'-dir'** flag.

If you decide to use an subscout configuration file, it will be automatically discovered when put in the output directory and named **config.ini**.

## The Configuration File

You will need a config file to use your API keys with subscout. See the [Example Configuration File](../examples/config.ini) for more details.

The location of the configuration file can be specified using the `-config` flag or the `subscout_CONFIG` environment variable.

subscout automatically tries to discover the configuration file (named `config.ini`) in the following locations:

| Operating System | Path |
| ---------------- | ---- |
| Linux / Unix | `$XDG_CONFIG_HOME/subscout/config.ini` or `$HOME/.config/subscout/config.ini` or `/etc/subscout/config.ini` |
| Windows | `%AppData%\subscout\config.ini` |
| OSX | `$HOME/Library/Application Support/subscout/config.ini` |

These are good places for you to put your configuration file.

Note that these locations are based on the [output directory](#the-output-directory). If you use the `-dir` flag, the location where subscout will try to discover the configuration file will change. For example, if you pass in `-dir ./my-out-dir`, subscout will try to discover a configuration file in `./my-out-dir/config.ini`.

### Default Section

| Option | Description |
|--------|-------------|
| mode | Determines which mode the enumeration is performed in: default, passive or active |
| output_directory | The directory that stores the graph database and other output files |
| maximum_dns_queries | The maximum number of concurrent DNS queries that can be performed |
| torfriendly | Enable compatibility mode to allow subdomain enumeration over Tor |
| srv_lookup | Enable service name lookup |
| disable_rdns | Disable reverse DNS (PTR) lookups |
| queries_per_server | Maximum queries per resolver per second |

### The `resolvers` Section

| Option | Description |
|--------|-------------|
| resolver | The IP address of a DNS resolver and used globally by the subscout package |

### The `trusted_resolvers` Section

| Option | Description |
|--------|-------------|
| resolver | The IP address of a DNS resolver and used globally by the subscout package |

### The `scope` Section

| Option | Description |
|--------|-------------|
| address | IP address or range (e.g. a.b.c.10-245) that is in scope |
| asn | ASN that is in scope |
| cidr | CIDR (e.g. 192.168.1.0/24) that is in scope |
| port | Specifies a port to be used when actively pulling TLS certificates or crawling |

#### The `scope.domains` Section

| Option | Description |
|--------|-------------|
| domain | A root DNS domain name to be added to the enumeration scope |

#### The `scope.blacklisted` Section

| Option | Description |
|--------|-------------|
| subdomain | A DNS subdomain name to be considered out of scope during the enumeration |

### The `graphdbs` Section

#### The `graphdbs.postgres` Section

| Option | Description |
|--------|-------------|
| primary | When set to true, the graph database is specified as the primary db |
| url | URL in the form of "postgres://[username:password@]host[:port]/database-name?sslmode=disable" where subscout will connect to a PostgreSQL database |
| options | Additional PostgreSQL database options |

#### The `graphdbs.mysql` Section

| Option | Description |
|--------|-------------|
| url | URL in the form of "[username:password@]tcp(host[:3306])/database-name?timeout=10s" where subscout will connect to a MySQL database |

### The `bruteforce` Section

| Option | Description |
|--------|-------------|
| enabled | When set to true, brute forcing is performed during the enumeration |
| recursive | When set to true, brute forcing is performed on discovered subdomain names as well |
| minimum_for_recursive | Number of discoveries made in a subdomain before performing recursive brute forcing |
| wordlist_file | Path to a custom wordlist file to be used during the brute forcing |

### The `alterations` Section

| Option | Description |
|--------|-------------|
| enabled | When set to true, permuting resolved DNS names is performed during the enumeration |
| edit_distance | Number of times an edit operation will be performed on a name sample during fuzzy label searching |
| flip_words | When set to true, causes words in DNS names to be exchanged for others in the alteration word list |
| flip_numbers | When set to true, causes numbers in DNS names to be exchanged for other numbers |
| add_words | When set to true, causes other words in the alteration word list to be added to resolved DNS names |
| add_numbers | When set to true, causes numbers to be added and removed from resolved DNS names |
| wordlist_file | Path to a custom wordlist file that provides additional words to the alteration word list |

### The `data_sources` Section

| Option | Description |
|--------|-------------|
| ttl | The number of minutes that the responses of **all** data sources for the target are cached |

#### The `data_sources.SOURCENAME` Section

| Option | Description |
|--------|-------------|
| ttl | The number of minutes that the response of the data source for the target is cached |

##### The `data_sources.SOURCENAME.CREDENTIALSETID` Section

| Option | Description |
|--------|-------------|
| apikey | The API key to be used when accessing the data source |
| secret | An additional secret to be used with the API key |
| username | User for the data source account |
| password | Valid password for the user identified by the 'username' option |

#### The `data_sources.disabled` Section

| Option | Description |
|--------|-------------|
| data_source | One of the subscout data sources that is **not** to be used during the enumeration |

## The Graph Database

All subscout enumeration findings are stored in a graph database. This database is either located in a single file within the output directory or connected to remotely using settings provided by the configuration file.

When a new enumeration begins and a graph database already exists with previous findings for the same target(s), the subdomain names from those previous enumerations are utilized in the new enumeration. New DNS queries are performed against those subdomain names to ensure that they are still legitimate and to obtain current IP addresses.

The results from each enumeration is stored separately in the graph database, which allows the tracking subcommand to look for differences across the enumerations and provide the user with highlights about the target.

There is nothing preventing multiple users from sharing a single (remote) graph database and leveraging each others findings across enumerations.

### Cayley Graph Schema

The GraphDB is storing all the domains that were found for a given enumeration. It stores the associated information such as the ip, ns_record, a_record, cname, ip block and associated source for each one of them as well. Each enumeration is identified by a uuid.

Here is an example of graph for an enumeration run on example.com:

![GraphDB](../images/example_graphDB.png)

## Importing subscout Results into Maltego

1. Convert the subscout data into a Maltego graph table CSV file:

```bash
subscout viz -maltego
```

2. Import the CSV file with the correct Connectivity Table settings:

![Connectivity table](../images/maltego_graph_import_wizard.png "Connectivity Table Settings")

3. All the subscout findings will be brought into your Maltego Graph:

![Maltego results](../images/maltego_results.png "Maltego Results")
