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

# Popular Projects
- [Web Frameworks](#Web-Frameworks)
  - [gin](#gin)
  - [beego](#beego)
  - [echo](#echo)
  - [fiber](#fiber)
  - [fiber](#fiber)
- [ORM](#ORM)
  - [gorm](#gorm)
  - [ent](#ent)
- [Cache](#Cache)
  - [go-redis(v8)](#go-redis(v8))

---  

## Web Frameworks

### gin
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | hasura | graphql-engine | 28581 | https://github.com/hasura/graphql-engine |
| 2 | asim | go-micro | 19615 | https://github.com/asim/go-micro |
| 3 | go-kratos | kratos | 19197 | https://github.com/go-kratos/kratos |
| 4 | cloudreve | Cloudreve | 16292 | https://github.com/cloudreve/Cloudreve |
| 5 | HFO4 | cloudreve | 16290 | https://github.com/HFO4/cloudreve |
| 6 | flipped-aurora | gin-vue-admin | 14942 | https://github.com/flipped-aurora/gin-vue-admin |
| 7 | OpenAPITools | openapi-generator | 14276 | https://github.com/OpenAPITools/openapi-generator |
| 8 | alist-org | alist | 11831 | https://github.com/alist-org/alist |
| 9 | Xhofe | alist | 11830 | https://github.com/Xhofe/alist |
| 10 | gotify | server | 7965 | https://github.com/gotify/server |
| 11 | dtm-labs | dtm | 7773 | https://github.com/dtm-labs/dtm |
| 12 | IceWhaleTech | CasaOS | 7190 | https://github.com/IceWhaleTech/CasaOS |
| 13 | appleboy | gorush | 6661 | https://github.com/appleboy/gorush |
| 14 | Terry-Mao | goim | 6604 | https://github.com/Terry-Mao/goim |
| 15 | GoAdminGroup | go-admin | 6432 | https://github.com/GoAdminGroup/go-admin |
| 16 | Mrs4s | go-cqhttp | 6220 | https://github.com/Mrs4s/go-cqhttp |
| 17 | eddycjy | go-gin-example | 5789 | https://github.com/eddycjy/go-gin-example |
| 18 | EDDYCJY | go-gin-example | 5787 | https://github.com/EDDYCJY/go-gin-example |
| 19 | crowdsecurity | crowdsec | 5734 | https://github.com/crowdsecurity/crowdsec |
| 20 | EasyDarwin | EasyDarwin | 5702 | https://github.com/EasyDarwin/EasyDarwin |
| 21 | easydarwin | EasyDarwin | 5702 | https://github.com/easydarwin/EasyDarwin |
| 22 | didi | nightingale | 5552 | https://github.com/didi/nightingale |
| 23 | chaos-mesh | chaos-mesh | 5285 | https://github.com/chaos-mesh/chaos-mesh |
| 24 | fnproject | fn | 5270 | https://github.com/fnproject/fn |
| 25 | Tencent | bk-cmdb | 4771 | https://github.com/Tencent/bk-cmdb |
| 26 | aler9 | rtsp-simple-server | 4157 | https://github.com/aler9/rtsp-simple-server |
| 27 | answerdev | answer | 3963 | https://github.com/answerdev/answer |
| 28 | douyu | jupiter | 3900 | https://github.com/douyu/jupiter |
| 29 | evrone | go-clean-template | 3794 | https://github.com/evrone/go-clean-template |
| 30 | darjun | go-daily-lib | 3525 | https://github.com/darjun/go-daily-lib |

**[⬆ top](#Popular-Projects)**


### beego
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | micro | go-micro | 19615 | https://github.com/micro/go-micro |
| 2 | vmware | harbor | 18605 | https://github.com/vmware/harbor |
| 3 | goharbor | harbor | 18605 | https://github.com/goharbor/harbor |
| 4 | GoAdminGroup | go-admin | 6432 | https://github.com/GoAdminGroup/go-admin |
| 5 | mindoc-org | mindoc | 6038 | https://github.com/mindoc-org/mindoc |
| 6 | lifei6671 | mindoc | 6038 | https://github.com/lifei6671/mindoc |
| 7 | silenceper | wechat | 3732 | https://github.com/silenceper/wechat |
| 8 | Qihoo360 | wayne | 3664 | https://github.com/Qihoo360/wayne |
| 9 | phachon | mm-wiki | 2956 | https://github.com/phachon/mm-wiki |
| 10 | truthhun | BookStack | 2942 | https://github.com/truthhun/BookStack |
| 11 | TruthHun | BookStack | 2942 | https://github.com/TruthHun/BookStack |
| 12 | truthhun | bookstack | 2942 | https://github.com/truthhun/bookstack |
| 13 | TruthHun | bookstack | 2942 | https://github.com/TruthHun/bookstack |
| 14 | TruthHun | Dochub | 2778 | https://github.com/TruthHun/Dochub |
| 15 | truthhun | DocHub | 2778 | https://github.com/truthhun/DocHub |
| 16 | truthhun | docHub | 2778 | https://github.com/truthhun/docHub |
| 17 | truthhun | dochub | 2778 | https://github.com/truthhun/dochub |
| 18 | TruthHun | DocHub | 2778 | https://github.com/TruthHun/DocHub |
| 19 | TruthHun | dochub | 2778 | https://github.com/TruthHun/dochub |
| 20 | aimerforreimu | AUXPI | 2651 | https://github.com/aimerforreimu/AUXPI |
| 21 | tigerb | easy-tips | 2647 | https://github.com/tigerb/easy-tips |
| 22 | TIGERB | easy-tips | 2647 | https://github.com/TIGERB/easy-tips |
| 23 | ysrc | yulong-hids-archived | 2041 | https://github.com/ysrc/yulong-hids-archived |
| 24 | ysrc | yulong-hids | 2041 | https://github.com/ysrc/yulong-hids |
| 25 | tobegit3hub | seagull | 1920 | https://github.com/tobegit3hub/seagull |
| 26 | feiyu563 | PrometheusAlert | 1904 | https://github.com/feiyu563/PrometheusAlert |
| 27 | feiyu563 | prometheusalert | 1904 | https://github.com/feiyu563/prometheusalert |
| 28 | smallnest | go-web-framework-benchmark | 1772 | https://github.com/smallnest/go-web-framework-benchmark |
| 29 | ulule | Limiter | 1655 | https://github.com/ulule/Limiter |
| 30 | ulule | limiter | 1655 | https://github.com/ulule/limiter |

**[⬆ top](#Popular-Projects)**


### echo
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | go-kratos | kratos | 19197 | https://github.com/go-kratos/kratos |
| 2 | earthly | earthly | 8187 | https://github.com/earthly/earthly |
| 3 | GoAdminGroup | go-admin | 6432 | https://github.com/GoAdminGroup/go-admin |
| 4 | 42wim | matterbridge | 5279 | https://github.com/42wim/matterbridge |
| 5 | gaia-pipeline | gaia | 4849 | https://github.com/gaia-pipeline/gaia |
| 6 | bytebase | bytebase | 4200 | https://github.com/bytebase/bytebase |
| 7 | gotenberg | gotenberg | 4177 | https://github.com/gotenberg/gotenberg |
| 8 | thecodingmachine | gotenberg | 4177 | https://github.com/thecodingmachine/gotenberg |
| 9 | ovh | cds | 4017 | https://github.com/ovh/cds |
| 10 | monitoror | monitoror | 3939 | https://github.com/monitoror/monitoror |
| 11 | douyu | jupiter | 3900 | https://github.com/douyu/jupiter |
| 12 | deepmap | oapi-codegen | 2627 | https://github.com/deepmap/oapi-codegen |
| 13 | studygolang | studygolang | 2429 | https://github.com/studygolang/studygolang |
| 14 | alibaba | sentinel-golang | 2161 | https://github.com/alibaba/sentinel-golang |
| 15 | hahwul | dalfox | 2115 | https://github.com/hahwul/dalfox |
| 16 | determined-ai | determined | 1909 | https://github.com/determined-ai/determined |
| 17 | smallnest | go-web-framework-benchmark | 1772 | https://github.com/smallnest/go-web-framework-benchmark |
| 18 | labstack | armor | 1647 | https://github.com/labstack/armor |
| 19 | julienschmidt | go-http-routing-benchmark | 1598 | https://github.com/julienschmidt/go-http-routing-benchmark |
| 20 | usememos | memos | 1357 | https://github.com/usememos/memos |
| 21 | ngoduykhanh | wireguard-ui | 1318 | https://github.com/ngoduykhanh/wireguard-ui |
| 22 | apache | servicecomb-service-center | 1316 | https://github.com/apache/servicecomb-service-center |
| 23 | EndlessCheng | mahjong-helper | 1297 | https://github.com/EndlessCheng/mahjong-helper |
| 24 | teamhanko | hanko | 1265 | https://github.com/teamhanko/hanko |
| 25 | mailslurper | mailslurper | 1184 | https://github.com/mailslurper/mailslurper |
| 26 | algorand | go-algorand | 1168 | https://github.com/algorand/go-algorand |
| 27 | shellhub-io | shellhub | 995 | https://github.com/shellhub-io/shellhub |
| 28 | cozy | cozy-stack | 963 | https://github.com/cozy/cozy-stack |
| 29 | Thiht | smocker | 868 | https://github.com/Thiht/smocker |
| 30 | ixre | go2o | 863 | https://github.com/ixre/go2o |

**[⬆ top](#Popular-Projects)**


### fiber
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | GoAdminGroup | go-admin | 6432 | https://github.com/GoAdminGroup/go-admin |
| 2 | j3ssie | osmedeus | 4087 | https://github.com/j3ssie/osmedeus |
| 3 | gofiber | recipes | 1840 | https://github.com/gofiber/recipes |
| 4 | smallnest | go-web-framework-benchmark | 1772 | https://github.com/smallnest/go-web-framework-benchmark |
| 5 | finb | bark-server | 1542 | https://github.com/finb/bark-server |
| 6 | awslabs | aws-lambda-go-api-proxy | 729 | https://github.com/awslabs/aws-lambda-go-api-proxy |
| 7 | 1340691923 | ElasticView | 623 | https://github.com/1340691923/ElasticView |
| 8 | kubeshop | testkube | 521 | https://github.com/kubeshop/testkube |
| 9 | phuslu | log | 457 | https://github.com/phuslu/log |
| 10 | create-go-app | fiber-go-template | 451 | https://github.com/create-go-app/fiber-go-template |
| 11 | douyu | juno | 421 | https://github.com/douyu/juno |
| 12 | bangumi | server | 415 | https://github.com/bangumi/server |
| 13 | HotPotatoC | twitter-clone | 366 | https://github.com/HotPotatoC/twitter-clone |
| 14 | darkweak | souin | 354 | https://github.com/darkweak/souin |
| 15 | gofiber | jwt | 337 | https://github.com/gofiber/jwt |
| 16 | sujit-baniya | fiber-boilerplate | 251 | https://github.com/sujit-baniya/fiber-boilerplate |
| 17 | brendonmatos | golive | 248 | https://github.com/brendonmatos/golive |
| 18 | arsmn | fiber-swagger | 241 | https://github.com/arsmn/fiber-swagger |
| 19 | gofiber | websocket | 228 | https://github.com/gofiber/websocket |
| 20 | koddr | tutorial-go-fiber-rest-api | 212 | https://github.com/koddr/tutorial-go-fiber-rest-api |
| 21 | mrusme | journalist | 195 | https://github.com/mrusme/journalist |
| 22 | dataplane-app | dataplane | 183 | https://github.com/dataplane-app/dataplane |
| 23 | Harry-027 | go-notify | 167 | https://github.com/Harry-027/go-notify |
| 24 | gofiber | template | 165 | https://github.com/gofiber/template |
| 25 | shipyard-run | shipyard | 164 | https://github.com/shipyard-run/shipyard |
| 26 | zekroTJA | shinpuru | 141 | https://github.com/zekroTJA/shinpuru |
| 27 | gofiber | adaptor | 133 | https://github.com/gofiber/adaptor |
| 28 | go-awesome | shortlink | 129 | https://github.com/go-awesome/shortlink |
| 29 | audioo | bitcrook | 123 | https://github.com/audioo/bitcrook |
| 30 | whitaker-io | machine | 117 | https://github.com/whitaker-io/machine |

**[⬆ top](#Popular-Projects)**


### fiber
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | GoAdminGroup | go-admin | 6432 | https://github.com/GoAdminGroup/go-admin |
| 2 | j3ssie | osmedeus | 4087 | https://github.com/j3ssie/osmedeus |
| 3 | gofiber | recipes | 1840 | https://github.com/gofiber/recipes |
| 4 | smallnest | go-web-framework-benchmark | 1772 | https://github.com/smallnest/go-web-framework-benchmark |
| 5 | finb | bark-server | 1542 | https://github.com/finb/bark-server |
| 6 | awslabs | aws-lambda-go-api-proxy | 729 | https://github.com/awslabs/aws-lambda-go-api-proxy |
| 7 | 1340691923 | ElasticView | 623 | https://github.com/1340691923/ElasticView |
| 8 | kubeshop | testkube | 521 | https://github.com/kubeshop/testkube |
| 9 | phuslu | log | 457 | https://github.com/phuslu/log |
| 10 | create-go-app | fiber-go-template | 451 | https://github.com/create-go-app/fiber-go-template |
| 11 | douyu | juno | 421 | https://github.com/douyu/juno |
| 12 | bangumi | server | 415 | https://github.com/bangumi/server |
| 13 | HotPotatoC | twitter-clone | 366 | https://github.com/HotPotatoC/twitter-clone |
| 14 | darkweak | souin | 354 | https://github.com/darkweak/souin |
| 15 | gofiber | jwt | 337 | https://github.com/gofiber/jwt |
| 16 | sujit-baniya | fiber-boilerplate | 251 | https://github.com/sujit-baniya/fiber-boilerplate |
| 17 | brendonmatos | golive | 248 | https://github.com/brendonmatos/golive |
| 18 | arsmn | fiber-swagger | 241 | https://github.com/arsmn/fiber-swagger |
| 19 | gofiber | websocket | 228 | https://github.com/gofiber/websocket |
| 20 | koddr | tutorial-go-fiber-rest-api | 212 | https://github.com/koddr/tutorial-go-fiber-rest-api |
| 21 | mrusme | journalist | 195 | https://github.com/mrusme/journalist |
| 22 | dataplane-app | dataplane | 183 | https://github.com/dataplane-app/dataplane |
| 23 | Harry-027 | go-notify | 167 | https://github.com/Harry-027/go-notify |
| 24 | gofiber | template | 165 | https://github.com/gofiber/template |
| 25 | shipyard-run | shipyard | 164 | https://github.com/shipyard-run/shipyard |
| 26 | zekroTJA | shinpuru | 141 | https://github.com/zekroTJA/shinpuru |
| 27 | gofiber | adaptor | 133 | https://github.com/gofiber/adaptor |
| 28 | go-awesome | shortlink | 129 | https://github.com/go-awesome/shortlink |
| 29 | audioo | bitcrook | 123 | https://github.com/audioo/bitcrook |
| 30 | whitaker-io | machine | 117 | https://github.com/whitaker-io/machine |

**[⬆ top](#Popular-Projects)**




## ORM

### gorm
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | go-kratos | kratos | 19197 | https://github.com/go-kratos/kratos |
| 2 | flipped-aurora | gin-vue-admin | 14942 | https://github.com/flipped-aurora/gin-vue-admin |
| 3 | milvus-io | milvus | 14147 | https://github.com/milvus-io/milvus |
| 4 | alist-org | alist | 11831 | https://github.com/alist-org/alist |
| 5 | Xhofe | alist | 11830 | https://github.com/Xhofe/alist |
| 6 | juanfont | headscale | 8815 | https://github.com/juanfont/headscale |
| 7 | dtm-labs | dtm | 7773 | https://github.com/dtm-labs/dtm |
| 8 | IBAX-io | go-ibax | 7445 | https://github.com/IBAX-io/go-ibax |
| 9 | IceWhaleTech | CasaOS | 7190 | https://github.com/IceWhaleTech/CasaOS |
| 10 | pyroscope-io | pyroscope | 6616 | https://github.com/pyroscope-io/pyroscope |
| 11 | zhenghaoz | gorse | 6313 | https://github.com/zhenghaoz/gorse |
| 12 | didi | nightingale | 5552 | https://github.com/didi/nightingale |
| 13 | aquasecurity | kube-bench | 5290 | https://github.com/aquasecurity/kube-bench |
| 14 | anchore | grype | 4750 | https://github.com/anchore/grype |
| 15 | bishopfox | sliver | 4567 | https://github.com/bishopfox/sliver |
| 16 | j3ssie | osmedeus | 4087 | https://github.com/j3ssie/osmedeus |
| 17 | xinliangnote | go-gin-api | 3971 | https://github.com/xinliangnote/go-gin-api |
| 18 | douyu | jupiter | 3900 | https://github.com/douyu/jupiter |
| 19 | porter-dev | porter | 3515 | https://github.com/porter-dev/porter |
| 20 | naiba | nezha | 3405 | https://github.com/naiba/nezha |
| 21 | photoview | photoview | 3187 | https://github.com/photoview/photoview |
| 22 | zu1k | proxypool | 3092 | https://github.com/zu1k/proxypool |
| 23 | rocboss | paopao-ce | 2530 | https://github.com/rocboss/paopao-ce |
| 24 | openservicemesh | osm | 2508 | https://github.com/openservicemesh/osm |
| 25 | erda-project | erda | 2436 | https://github.com/erda-project/erda |
| 26 | marmotedu | iam | 2363 | https://github.com/marmotedu/iam |
| 27 | hwholiday | learning_tools | 2347 | https://github.com/hwholiday/learning_tools |
| 28 | LyricTian | gin-admin | 2102 | https://github.com/LyricTian/gin-admin |
| 29 | 8treenet | freedom | 2092 | https://github.com/8treenet/freedom |
| 30 | openflagr | flagr | 2036 | https://github.com/openflagr/flagr |

**[⬆ top](#Popular-Projects)**


### ent
| No | User | Repository | Stargazers | URL |
| ---:| --- | --- | ---:| --- |
| 1 | go-kratos | kratos | 19197 | https://github.com/go-kratos/kratos |
| 2 | crowdsecurity | crowdsec | 5734 | https://github.com/crowdsecurity/crowdsec |
| 3 | sagikazarmark | modern-go-application | 1346 | https://github.com/sagikazarmark/modern-go-application |
| 4 | mikestefanello | pagoda | 528 | https://github.com/mikestefanello/pagoda |
| 5 | go-kratos | beer-shop | 516 | https://github.com/go-kratos/beer-shop |
| 6 | direktiv | direktiv | 311 | https://github.com/direktiv/direktiv |
| 7 | gitploy-io | gitploy | 228 | https://github.com/gitploy-io/gitploy |
| 8 | kcarretto | paragon | 223 | https://github.com/kcarretto/paragon |
| 9 | mrusme | journalist | 195 | https://github.com/mrusme/journalist |
| 10 | efectn | go-orm-benchmarks | 192 | https://github.com/efectn/go-orm-benchmarks |
| 11 | hay-kot | homebox | 190 | https://github.com/hay-kot/homebox |
| 12 | go-kratos | examples | 137 | https://github.com/go-kratos/examples |
| 13 | dopedao | RYO | 105 | https://github.com/dopedao/RYO |
| 14 | go-saas | saas | 101 | https://github.com/go-saas/saas |
| 15 | masseelch | elk | 85 | https://github.com/masseelch/elk |
| 16 | hedwigz | entviz | 81 | https://github.com/hedwigz/entviz |
| 17 | adnaan | gomodest-starter | 75 | https://github.com/adnaan/gomodest-starter |
| 18 | adnaan | gomodest | 75 | https://github.com/adnaan/gomodest |
| 19 | vicanso | tiny-site | 72 | https://github.com/vicanso/tiny-site |
| 20 | YadaYuki | omochi | 70 | https://github.com/YadaYuki/omochi |
| 21 | zibbp | ganymede | 66 | https://github.com/zibbp/ganymede |
| 22 | ngocphuongnb | tetua | 61 | https://github.com/ngocphuongnb/tetua |
| 23 | appditto | pippin_nano_wallet | 59 | https://github.com/appditto/pippin_nano_wallet |
| 24 | yumenaka | comi | 57 | https://github.com/yumenaka/comi |
| 25 | orestonce | ChessGame | 57 | https://github.com/orestonce/ChessGame |
| 26 | m-mizutani | octovy | 48 | https://github.com/m-mizutani/octovy |
| 27 | long2ice | longurl | 44 | https://github.com/long2ice/longurl |
| 28 | efectn | fiber-boilerplate | 35 | https://github.com/efectn/fiber-boilerplate |
| 29 | valocode | bubbly | 34 | https://github.com/valocode/bubbly |
| 30 | degenerat3 | meteor | 32 | https://github.com/degenerat3/meteor |

**[⬆ top](#Popular-Projects)**