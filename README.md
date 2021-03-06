[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Documentation](https://godoc.org/github.com/jmbarzee/dominion?status.svg)](https://godoc.org/github.com/jmbarzee/dominion)
[![Code Quality](https://goreportcard.com/badge/github.com/jmbarzee/dominion)](https://goreportcard.com/report/github.com/jmbarzee/dominion)
[![Build Status](https://github.com/jmbarzee/dominion/workflows/build/badge.svg)](https://github.com/jmbarzee/dominion/actions)
[![Coverage](https://codecov.io/gh/jmbarzee/dominion/branch/master/graph/badge.svg)](https://codecov.io/gh/jmbarzee/dominion)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjmbarzee%2Fdominion.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjmbarzee%2Fdominion?ref=badge_shield)


# Dominion
Golang distributed system for service ecosystem management

This library serves on an IoT network were services (lights, speakers, thermostat, cameras, processing...) will be auto-started, auto-distributed, and (maybe) auto-scaled. The Dominion & Domains are sorts of managers of all devices which are participating. They manage service start, maintance, and discovery.


## Abstractions
### Dominion (leader)
Run the Domain with `go run cmd/dominion/main.go`

Don't forget to set `DOMINION_CONFIG_FILE` [example](../main/cmd/dominion/ex.config.toml)

Listen for new Domains by:
1. Wait for ZeroConf Broadcasts advertizing a Domain
2. Dial Domain & establish lasting connection

Review Domains by:
1. Repeadetly send heartbeat to Domains to keep connection alive
2. Update Domains service list from heartbeat reply

Review Domain's Services by:
1. Routinely reviewing Domains and checking that required services are available/started 
2. Routinely reviewing Services and checking that dependencies are available/started



### Domain (follower)
Run the Domain with `go run cmd/domain/main.go`

Don't forget to set `DOMAIN_CONFIG_FILE` [example](../main/cmd/domain/ex.config.toml)

Domains find a Dominion by:
1. Identifying that they are lonely (no history of a dominion or heartbeats stopped)
2. Broadcasting to the network using ZeroConf

Domains remain in a Dominion by:
1. Listening for incomming Heartbeat RPCs
2. Update stored Dominion's identity 

Domains start Services by:
1. Listening for incomming StartService RPCs
2. Calling make in the specified Service Directory



### Service (Ecosystem) 
Services do whatever you want them too. Services are language agnostic. They can locate other services through the Dominion's GRPC server. Service dependencies are defined in the `DOMINION_CONFIG_FILE`

Services I use -> [ExMachina](github.com/jmbarzee/exmachina)


## Utilized Libraries

`github.com/blang/semver`

`google.golang.org/grpc`

`github.com/grandcat/zeroconf`

`github.com/BurntSushi/toml`



## Planned Development

1. Expose Logfiles - Dominion & Domain offer routes with logfiles
2. Connection encryption - encrypt RPCs
3. Identity verification - sign communication with preestablished keypairs


## Development Tools
[gRPC & protoc](https://grpc.io/docs/languages/go/quickstart/) are used by `go generate` to update [bitbox/grpc](grpc/).

[gRPCox](https://github.com/gusaul/grpcox) is a lightweight docker container for easy manual testing.