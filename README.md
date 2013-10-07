# dns-canary

Run repeated DNS queries against a specific cache or the first one named in `/etc/resolv.conf` and emit success/error metrics.

## Configuration

dns-canary is configured via the environment. The variables it looks for:

* `NAMES`: a `,`-separated list of names to query for
* `INTERVAL`: how often to query each name in `NAMES`, needs to be acceptable for [`time.ParseDuration`](http://golang.org/pkg/time/#ParseDuration)
* `SERVER`: if specified, use the named server at `<ip>:<port>`. If unspecified, the first server in `/etc/resolv.conf` is used
