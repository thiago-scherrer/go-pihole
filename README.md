# go-pihole

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/go-pihole)](https://goreportcard.com/report/github.com/thiago-scherrer/go-pihole) ![go pihole](https://github.com/thiago-scherrer/go-pihole/actions/workflows/build.yml/badge.svg)

A simple golang project to update pi-hole-block-list project

### Requirements

- go >= 1.16.3


### Running

Clone:
```sh
git clone https://github.com/thiago-scherrer/go-pihole.git
```

Build:
```sh
cd go-pihole
go build cmd/go-pihole.go
```

Running
```sh
export LIST=<list to consult>
export OUTPUT=<where to save>
./go-pihole
```

