# Docker-Registry-Cleaner
> *n. person responsible for scrubbing a ship's deck.*

[![Go Report Card](https://goreportcard.com/badge/github.com/yxxhero/docker-registry-cleane)](https://goreportcard.com/report/github.com/yxxhero/docker-registry-cleaner)

Docker-Registry-Cleaner inspects images of a [Docker Registry](https://docs.docker.com/registry/) and removes those older than a given age.

## Thanks
https://github.com/fraunhoferfokus/deckschrubber

## Quick Start

```bash
go get github.com/yxxhero/docker-registry-cleane
$GOPATH/bin/docker-registry-cleaner
```

## Why this?
We run our own private registry on a server with limited storage and it was only a question of time, until it was full! Although there are similar approaches to do Docker registry house keeping (in Python, Ruby, etc), a native module (using Docker's own packages) was missing. This module has the following advantages:

* **Is binary!**: No runtime, Python, Ruby, etc. is required
* **Uses Docker API**: same packages and modules used to relaze Docker registry are used here

## Arguments
```
-day int
      max age in days
-debug
      run in debug mode      
-dry
      does not actually deletes
-latest int
      number of the latest matching images of an repository that won't be deleted (default 1)      
-month int
      max age in months
-registry string
      URL of registry (default "http://localhost:5000")
-repo string
      matching repositories (allows regexp) (default ".*")      
-repos int
      number of repositories to garbage collect (default 5)
-tag string
      matching tags (allows regexp) (default ".*")      
-ntag string
      match everything but this tag (allows regexp) (default empty)
-v    shows version and quits
-year int
      max age in days
      
      
```

## Registry preparation
Docker-Registry-Cleaner uses the Docker Registry API. 
Its delete endpoint is disabled by default, you have to enable it with the following entry in the registry configuration file: 

```
delete:
  enabled: true
```

See [the documentation](https://github.com/docker/distribution/blob/master/docs/configuration.md#delete) for details. 

## Examples

* **Remove all images older than 2 months and 2 days**

```
$GOPATH/bin/docker-registry-cleaner -month 2 -day 2
```

* **Remove all images older than 1 year from `http://myrepo:5000`**

```
$GOPATH/bin/docker-registry-cleaner -year 1 -registry http://myrepo:5000
```

* **Analyize (but do not remove) images of 30 repositories**

```
$GOPATH/bin/docker-registry-cleaner -repos 30 -dry
```

* **Remove all images of each repository except the 3 latest ones**

```
$GOPATH/bin/docker-registry-cleaner -latest 3 
```

* **Remove all images with tags that ends with '-SNAPSHOT'**

```
$GOPATH/bin/docker-registry-cleaner -tag ^.*-SNAPSHOT$ 
```
