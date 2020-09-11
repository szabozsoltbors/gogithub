# GoGin

A simple GO project with Gin web framework

## Overview

```
.
|──build                // Build folder for generated binaries
|──src
|   |──api              // Schema files, protocol definition files
|   |──internal         // Private application and library code
|   |──web              // Web application specific components
```

## Build he project

To generate the binaries issue the following command:

```sh
make build
```

## Run the web application

To run the web server issue the following command:

```sh
make run
```

## Clean the build folder

To clean up the build folder issue the following command:

```sh
make clean
```

## Clean build and run

Issue the following command to clean, build and run the project

```sh
make all
```
