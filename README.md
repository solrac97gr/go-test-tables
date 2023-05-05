# How to work with Test Tables and Mocks in Go

For this example we will use 2 external packages:

-gomock
-testify/assert

## Mocks

For the mocks I use gomock. It is a mock framework for Go. It integrates well with Go's built-in testing package, but can be used in other contexts too.

You can see in the file generate-mocks.sh how to generate the mocks.

## Test Tables

For the test tables I use a `map[string]struct{}` with the key being the name of the test and the value being a `struct{}`.

Inside of the struct we add all the external information we will need like parameters.

We also add there 2 functions:

- testSetup: This function will be called before the test is run. It will setup the behavior of the mocks and the test data.

- assertSetup: This function will be called before the test is run. It will setup the mocks and the test data.
