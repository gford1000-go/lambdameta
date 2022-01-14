[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)


LambdaMeta | Geoff Ford
=======================

Collates metadata associated with an AWS Lambda function Invoke, within a JSON encodeable struct `Meta`.

Create the `Meta` instance by calling `GetMeta()`.


Installing and building the library
===================================

This project requires Go 1.17.2

To use this package in your own code, use `go get`:

    go get github.com/gford1000-go/lambdameta

Then, you can include it in your project:

	import "github.com/gford1000-go/lambdameta"

