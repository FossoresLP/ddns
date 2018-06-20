# DDNS
A dynamic DNS client made for [NS1](https://ns1.com)

This repository contains a DDNS client that uses the [NS1 API](https://ns1.com/api) to set records to the public IPs of the machine running it.

It can be configured via command-line options (use `-h` to list all) or via a TOML config file (by default located in `/etc/ns1-ddns/config.toml`)

An external server is needed to get the IP addresses. You can either use the one provided in request-server and set it up on a device outside your network or use an API provided by a service like https://whatismyipaddress.com which is used in the sample config file.

**This project is in no way affiliated with or endorsed by NS1.**
