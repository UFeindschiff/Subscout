// Copyright © by Jeff Foley 2017-2022. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/OWASP/Amass/v3/net/http"
	"github.com/caffix/stringset"
	"github.com/go-ini/ini"
)

// DefaultQueriesPerPublicResolver is the number of queries sent to each public DNS resolver per second.
var DefaultQueriesPerPublicResolver = 150

// DefaultQueriesPerBaselineResolver is the number of queries sent to each trusted DNS resolver per second.
var DefaultQueriesPerBaselineResolver = 150

const minResolverReliability = 0.85


// PublicResolvers includes the addresses of public resolvers obtained dynamically.
var PublicResolvers []string

// GetPublicDNSResolvers obtains the public DNS server addresses from public-dns.info and assigns them to PublicResolvers.
func GetPublicDNSResolvers() error {
	url := "https://public-dns.info/nameservers-all.csv"
	page, err := http.RequestWebPage(context.Background(), url, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to obtain the Public DNS csv file at %s: %v", url, err)
	}

	var resolvers []string
	var ipIdx, reliabilityIdx int
	r := csv.NewReader(strings.NewReader(page))
	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if i == 0 {
			for idx, val := range record {
				if val == "ip_address" {
					ipIdx = idx
				} else if val == "reliability" {
					reliabilityIdx = idx
				}
			}
			continue
		}
		if rel, err := strconv.ParseFloat(record[reliabilityIdx], 64); err == nil && rel >= minResolverReliability {
			resolvers = append(resolvers, record[ipIdx])
		}
	}
	
	for _, addr := range resolvers {
		PublicResolvers = append(PublicResolvers, addr)
	}
	return nil
}

// SetResolvers assigns the untrusted resolver names provided in the parameter to the list in the configuration.
func (c *Config) SetResolvers(resolvers ...string) {
	c.Resolvers = []string{}
	c.AddResolvers(resolvers...)
}

// AddResolvers appends the untrusted resolver names provided in the parameter to the list in the configuration.
func (c *Config) AddResolvers(resolvers ...string) {
	for _, r := range resolvers {
		c.AddResolver(r)
	}
	c.CalcMaxQPS()
}

// AddResolver appends the untrusted resolver name provided in the parameter to the list in the configuration.
func (c *Config) AddResolver(resolver string) {
	c.Lock()
	defer c.Unlock()

	// Check that the domain string is not empty
	r := strings.TrimSpace(resolver)
	if r == "" {
		return
	}

	c.Resolvers = stringset.Deduplicate(append(c.Resolvers, resolver))
}

// SetTrustedResolvers assigns the trusted resolver names provided in the parameter to the list in the configuration.
func (c *Config) SetTrustedResolvers(resolvers ...string) {
	c.Resolvers = []string{}
	c.AddResolvers(resolvers...)
}

// AddTrustedResolvers appends the trusted resolver names provided in the parameter to the list in the configuration.
func (c *Config) AddTrustedResolvers(resolvers ...string) {
	for _, r := range resolvers {
		c.AddTrustedResolver(r)
	}
	c.CalcMaxQPS()
}

// AddTrustedResolver appends the trusted resolver name provided in the parameter to the list in the configuration.
func (c *Config) AddTrustedResolver(resolver string) {
	c.Lock()
	defer c.Unlock()

	// Check that the domain string is not empty
	r := strings.TrimSpace(resolver)
	if r == "" {
		return
	}

	c.TrustedResolvers = stringset.Deduplicate(append(c.TrustedResolvers, resolver))
}

// CalcMaxQPS updates the MaxDNSQueries field of the configuration based on current settings.
func (c *Config) CalcMaxQPS() {
	c.MaxDNSQueries = (len(c.Resolvers) * c.ResolversQPS) + (len(c.TrustedResolvers) * c.TrustedQPS)
}

func (c *Config) loadResolverSettings(cfg *ini.File) error {
	sec, err := cfg.GetSection("resolvers")
	if err != nil {
		return nil
	}

	c.Resolvers = stringset.Deduplicate(sec.Key("resolver").ValueWithShadows())
	if len(c.Resolvers) == 0 {
		return errors.New("no resolver keys were found in the resolvers section")
	}

	return nil
}

func (c *Config) loadTrustedResolverSettings(cfg *ini.File) error {
	sec, err := cfg.GetSection("trusted_resolvers")
	if err != nil {
		return nil
	}

	c.TrustedResolvers = stringset.Deduplicate(sec.Key("resolver").ValueWithShadows())
	if len(c.TrustedResolvers) == 0 {
		return errors.New("no resolver keys were found in the trusted_resolvers section")
	}

	return nil
}
