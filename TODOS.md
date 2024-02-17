# Current TODOs for subscout
* Migrate away from caffix' libraries which seem to be written only with Amass in mind and have features stripped away from them the moment upstream Amass no longer needs them, which makes them unsuitable long-term use in subscout. These are:
 - github.com/caffix/netmap (this will take some time)
 - github.com/caffix/pipeline
 - github.com/caffix/queue
 - github.com/caffix/resolve (should be relatively easy)
 - github.com/caffix/service
 - github.com/caffix/stringset
* more datasources
* Better DNS bruteforcing, better default wordlist
* Timeout functionality
* better exports
* better documentation
* fixes (will have to go through amass' issue list to see which ones affect subscout as well)
