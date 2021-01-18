# Fizzbuzz

[![Build Status](https://travis-ci.com/aftouh/fizzbuzz.svg?branch=main)](https://travis-ci.com/github/aftouh/fizzbuzz)
[![Coverage Status](https://coveralls.io/repos/github/aftouh/fizzbuzz/badge.svg?branch=main)](https://coveralls.io/github/aftouh/fizzbuzz?branch=main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aftouh/fizzbuzz)

A REST API implementation of Fizzbuzz.  
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

In this version, the endpoint `/v1/fizzbuzz` accepts five parameters : three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2` and returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
