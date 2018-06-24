# DDNS
A dynamic DNS client made for [NS1](https://ns1.com)

This repository contains a DDNS client that uses the [NS1 API](https://ns1.com/api) through their [golang package](https://github.com/ns1/ns1-go) to set records to the public IPs of the machine running it.

It can be configured via command-line options (use `-h` to list all) or via a TOML config file (by default located in `/etc/ns1-ddns/config.toml`). A sample config is included in this repository as [config.toml](https://github.com/FossoresLP/ddns/blob/master/config.toml).

You need an external server that returns the IPv4 and IPv6 addresses of the client on seperate subdomains. There are three options:
1. You can use the one provided in request-server (it uses port 2460 to work independently from any webserver) and set it up on a device outside your network.
2. You can use an API provided by a service like https://whatismyipaddress.com which is used in the sample config file. Be aware of the [strict rate limiting](https://whatismyipaddress.com/api) though.
3. You can develop your own solution and deploy that.

Make sure that your choosen solution provides two subdomains, one with only an A record and a second one with only an AAAA record.

**This project is in no way affiliated with or endorsed by NS1.**
