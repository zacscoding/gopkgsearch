# Search for best practices

## Motivation

Search for best practices in your golang library.  
Although you can see the imported projects in the [go packages](https://pkg.go.dev/), It is difficult to find good projects.  
Because it is just sorted by project name. `gopkgsearch` helps you find projects sorted by stars.

## Getting started

**1. install**

```shell
$ go install github.com/zacscoding/gopkgsearch@latest

or build from source code

$ git clone https://github.com/zacscoding/gopkgsearch.git
$ cd gopkgsearch
$ make build
$ ls -la ./build/bin
```

**2. search repositories**

```shell
$ ./build/bin/gopkgsearch imported -n 10 -p github.com/aws/aws-sdk-go/service/cloudhsmv2 -stars '>=1' -output table
2022/11/09 00:04:14 Search imported github repositories. limit: 10, package: github.com/aws/aws-sdk-go/service/cloudhsmv2, enable github api: false
2022/11/09 00:04:15 > imported repositories: 162
....
+----+---------------------+-----------------------------------+------------+---------------------------------------------------------------+
| NO | USER                | REPOSITORY                        | STARGAZERS | URL                                                           |
+----+---------------------+-----------------------------------+------------+---------------------------------------------------------------+
|  1 | terraform-providers | terraform-provider-aws            |       7884 | https://github.com/terraform-providers/terraform-provider-aws |
|  2 | rebuy-de            | aws-nuke                          |       3938 | https://github.com/rebuy-de/aws-nuke                          |
|  3 | awsiv               | terraform-provider-rdsdataservice |         10 | https://github.com/awsiv/terraform-provider-rdsdataservice    |
|  4 | Optum               | aws-nuke                          |          5 | https://github.com/Optum/aws-nuke                             |
|  5 | glassechidna        | awsctx                            |          4 | https://github.com/glassechidna/awsctx                        |
|  6 | phzietsman          | terraform-provider-awsx           |          1 | https://github.com/phzietsman/terraform-provider-awsx         |
|  7 | lonnblad            | terraform-provider-aws            |          1 | https://github.com/lonnblad/terraform-provider-aws            |
|  8 | shatil              | aws-sdk-go                        |          1 | https://github.com/shatil/aws-sdk-go                          |
|  9 | roberth-k           | aws-sdk-go                        |          1 | https://github.com/roberth-k/aws-sdk-go                       |
| 10 | DrFaust92           | terraform-provider-aws            |          1 | https://github.com/DrFaust92/terraform-provider-aws           |
+----+---------------------+-----------------------------------+------------+---------------------------------------------------------------+

$ ./build/bin/gopkgsearch imported -h                                                                              
Usage of imported:
  -n uint
        Number of retrieved repositories. (default 10)
  -o string
        Output path. default: std
  -output string
        Output format. available: ['table', 'json', 'yaml', 'markdown' (default "table")
  -p string
        Golang Package. e.g) net/http, github.com/gin-gonic/gin
  -stars string
        Stars filter. e.g) '=10', '>10', '>=10'
```

---