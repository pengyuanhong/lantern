package main

import "flag"

var (
	configdir     = flag.String("configdir", "", "directory in which to store configuration, including flashlight.yaml (defaults to current directory)")
	cloudconfig   = flag.String("cloudconfig", "", "optional http(s) URL to a cloud-based source for configuration updates")
	cloudconfigca = flag.String("cloudconfigca", "", "optional PEM encoded certificate used to verify TLS connections to fetch cloudconfig")
	addr          = flag.String("addr", "", "ip:port on which to listen for requests. When running as a client proxy, we'll listen with http, when running as a server proxy we'll listen with https (required)")
	unencrypted   = flag.Bool("unencrypted", false, "set to true to run server in unencrypted mode (no TLS)")
	role          = flag.String("role", "", "either 'client' or 'server' (required)")
	frontFQDNs    = flag.String("frontfqdns", "", "YAML string representing a map from the name of each front provider to a FQDN that will reach this particular server via that provider (e.g. '{cloudflare: fl-001.getiantem.org, cloudfront: blablabla.cloudfront.net}')")
	statsPeriod   = flag.Int("statsperiod", 0, "time in seconds to wait between reporting stats. If not specified, stats are not reported. If specified, statshub, instanceid and statshubAddr must also be specified.")
	statshubAddr  = flag.String("statshub", "pure-journey-3547.herokuapp.com", "address of statshub server")
	instanceid    = flag.String("instanceid", "", "instanceId under which to report stats to statshub. If not specified, no stats are reported.")
	registerat    = flag.String("registerat", "", "base URL for peer DNS registry at which to register (e.g. https://peerscanner.getiantem.org)")
	country       = flag.String("country", "xx", "2 digit country code under which to report stats. Defaults to xx.")
	cpuprofile    = flag.String("cpuprofile", "", "write cpu profile to given file")
	memprofile    = flag.String("memprofile", "", "write heap profile to given file")
	portmap       = flag.Int("portmap", 0, "try to map this port on the firewall to the port on which flashlight is listening, using UPnP or NAT-PMP. If mapping this port fails, flashlight will exit with status code 50")
	uiaddr        = flag.String("uiaddr", "", "if specified, indicates host:port the UI HTTP server should be started on")
	proxyAll      = flag.Bool("proxyall", false, "set to true to proxy all traffic through Lantern network")
	stickyConfig  = flag.Bool("stickyconfig", false, "set to true to only use the local config file")
)

// flagsAsMap returns a map of all flags that were provided at runtime
func flagsAsMap() map[string]interface{} {
	flags := make(map[string]interface{})
	flag.Visit(func(f *flag.Flag) {
		flags[f.Name] = f.Value.(flag.Getter).Get()
	})
	return flags
}