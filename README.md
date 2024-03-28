# letCodeSpeakForItself

## Overview
This is a simple Golang web application backed by a Postgres database used to showcase my DevOps skills. To build this app, please visit this [repo](https://github.com/yuanrenc/IaCForLetCodeSpeakForItself).

## Tech used(showed) in this project
- Golang
- Github Actions
- Docker
- web application design

## Tech used to build this project
- AWS
- Terraform
- MakeFile

## IF you are interested 

You could run the following command to download this application and run it locally. 
```shell
docker pull colinwang847/letcodespeakforitself:latest
```

## How it works 

### configuration
Environment variables has precedence over configuration from the `config.yaml` file.And the application will look for environment variables that are able to override the configuration defined in the `config.yaml` file. These environment variables are prefixed with `SD` and follow this pattern `SD_<config value>`. e.g. `SD_LISTENPORT`. Note, they are UPPERCASE.

### Commands
`database` to create a database, tables, and seed it with test data. 

`server` will start serving requests
