package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/miekg/dns"
)

func check(name string, duration time.Duration) {
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(name), dns.TypeA)
	m.RecursionDesired = true

	for _ = range time.Tick(duration) {
		r, _, err := c.Exchange(m, config.Servers[0]+":"+config.Port)
		if err != nil {
			fmt.Printf("count#dns-canary.error=1 name=%q error=%q r=%q\n", name, err, r)
			continue
		}
		if r.Rcode != dns.RcodeSuccess {
			fmt.Printf("count#dns-canary.error=1 name=%q r=%q\n", name, err, r)
			continue
		}

		fmt.Printf("count#dns-canary.success=1 name=%q r=%q\n", name, r)
		fmt.Printf("count#dns-canary.error=0\n")
	}
}

func main() {
	namesEnv := os.Getenv("NAMES")
	if namesEnv == "" {
		panic("missing $NAMES")
	}
	names := strings.Split(namesEnv, ",")

	intervalEnv := os.Getenv("INTERVAL")
	if intervalEnv == "" {
		panic("missing $INTERVAL")
	}
	interval, err := time.ParseDuration(intervalEnv)
	if err != nil {
		panic("can't parse $DURATION")
	}

	for _, name := range names {
		go check(name, interval)
	}

	select {}
}
