# Fizzbuzz

[![Build Status](https://travis-ci.com/aftouh/fizzbuzz.svg?branch=main)](https://travis-ci.com/github/aftouh/fizzbuzz)
[![Coverage Status](https://coveralls.io/repos/github/aftouh/fizzbuzz/badge.svg?branch=main)](https://coveralls.io/github/aftouh/fizzbuzz?branch=main)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aftouh_fizzbuzz&metric=alert_status)](https://sonarcloud.io/dashboard?id=aftouh_fizzbuzz)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aftouh/fizzbuzz)

A REST API implementation of Fizzbuzz.  
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

In this version, the endpoint `/v1/fizzbuzz` accepts five parameters : three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2` and returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

example:

```bash
# Run the server locally before using: make build && ./out/bin/fizzbuzz
curl 'http://localhost:8080/v1/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz'
```

## Repository organization

This section describes the repository folders:

- `config`: application config is handled using a yaml file that is pared with the [viper library](github.com/spf13/viper)
  The default file is the `config.yaml` one that could be overridden by setting the `CONFIG_PATH` environment variable
- `router`: creates the server routers and middlewares based on [go-chi library](github.com/go-chi/chi)
- `handlers`: implements the function handler of each server endpoint
- `telemetry`: configure application logger, tracer and meter
- `telemetry/logger`: Logger is based on the [zap library](go.uber.org/zap)
- `telemetry/tracer`: Tracer exporter and provider are created using the [opentelemetry sdk](https://github.com/open-telemetry/opentelemetry-go). Server request spans are generated using the [httptracer middleware](github.com/go-chi/httptracer)
- `telemetry/meter`: Exposes server runtime metrics on `:8082/metrics` endpoint. These metrics could be scrapped by prometheus
- `chart`: contains a helm chart used to deploy the application in a kubernetes cluster
