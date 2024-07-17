module github.com/hopeio/utils

go 1.22

require (
	bou.ke/monkey v1.0.2
	github.com/99designs/gqlgen v0.17.47
	github.com/IBM/sarama v1.43.2
	github.com/alta/protopatch v0.5.3
	github.com/aws/aws-sdk-go v1.47.2
	github.com/boombuler/barcode v1.0.1
	github.com/cockroachdb/pebble v1.1.0
	github.com/dgraph-io/badger/v3 v3.2103.5
	github.com/dgraph-io/ristretto v0.1.2-0.20240116140435-c67e07994f91
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gin-gonic/gin v1.10.0
	github.com/go-oauth2/oauth2/v4 v4.5.2
	github.com/go-openapi/loads v0.22.0 // indirect
	github.com/go-openapi/runtime v0.28.0
	github.com/go-openapi/spec v0.21.0
	github.com/go-openapi/swag v0.23.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.20.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/uuid v1.6.0
	github.com/gorilla/sessions v1.2.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/json-iterator/go v1.1.12
	github.com/markbates/goth v1.79.0
	github.com/microcosm-cc/bluemonday v1.0.26
	github.com/modern-go/reflect2 v1.0.2
	github.com/mozillazg/go-pinyin v0.20.0
	github.com/nsqio/go-nsq v1.1.0
	github.com/olivere/elastic/v7 v7.0.32
	github.com/prometheus/client_golang v1.19.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/pretty v1.2.1
	github.com/ugorji/go/codec v1.2.12
	github.com/valyala/fasthttp v1.53.0
	github.com/xuri/excelize/v2 v2.8.1
	go.etcd.io/etcd/client/v3 v3.5.13
	go.uber.org/atomic v1.11.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.25.0
	golang.org/x/oauth2 v0.20.0
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/grpc v1.64.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.34.1
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.6
	gorm.io/driver/postgres v1.5.7
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.10
	gorm.io/plugin/prometheus v0.1.0
)

require (
	github.com/apache/arrow-adbc/go/adbc v1.0.0
	github.com/blastrain/vitess-sqlparser v0.0.0-20201030050434-a139afbb1aba
	github.com/bytedance/sonic v1.11.7
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/chromedp/cdproto v0.0.0-20240202021202-6d0b6a386732
	github.com/chromedp/chromedp v0.9.5
	github.com/coocood/freecache v1.2.4
	github.com/danielvladco/go-proto-gql v0.10.1-0.20221227181908-22fca0a9469c
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/elastic/go-elasticsearch/v7 v7.17.10
	github.com/elastic/go-elasticsearch/v8 v8.13.1
	github.com/gofiber/fiber/v3 v3.0.0-beta.2
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/gonutz/w32/v2 v2.11.1
	github.com/grafana/loki-client-go v0.0.0-20230116142646-e7494d0ef70c
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.1
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/hashicorp/raft v1.6.1
	github.com/influxdata/influxdb-client-go/v2 v2.13.0
	github.com/jhump/protoreflect v1.16.0
	github.com/linxGnu/grocksdb v1.9.1
	github.com/lni/dragonboat/v4 v4.0.0-20231222133740-1d6e2d76cd57
	github.com/lni/goutils v1.4.0
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e
	github.com/marcboeker/go-duckdb v1.7.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/mwitkow/grpc-proxy v0.0.0-20230212185441-f345521cb9c9
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.5
	github.com/pelletier/go-toml/v2 v2.2.2
	github.com/prometheus/common v0.53.0
	github.com/quic-go/quic-go v0.44.0
	github.com/rs/cors v1.11.0
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/tidwall/buntdb v1.3.1
	github.com/tidwall/gjson v1.17.1
	github.com/tidwall/sjson v1.2.5
	github.com/ugorji/go v1.2.12
	github.com/vektah/gqlparser/v2 v2.5.12
	go.etcd.io/etcd/raft/v3 v3.5.13
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/prometheus v0.49.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.27.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.27.0
	go.opentelemetry.io/otel/metric v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
	go.opentelemetry.io/otel/sdk/metric v1.27.0
	go.opentelemetry.io/otel/trace v1.27.0
	golang.org/x/arch v0.8.0
	golang.org/x/crypto v0.24.0
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842
	golang.org/x/image v0.16.0
	golang.org/x/sys v0.21.0
	golang.org/x/text v0.16.0
	golang.org/x/time v0.5.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240513163218-0867130af1f8
	gopkg.in/ini.v1 v1.67.0
	nhooyr.io/websocket v1.8.11
)

require (
	cloud.google.com/go v0.112.2 // indirect
	cloud.google.com/go/auth v0.4.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	cloud.google.com/go/firestore v1.15.0 // indirect
	cloud.google.com/go/longrunning v0.5.6 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/VictoriaMetrics/metrics v1.33.1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alibabacloud-go/darabonba-array v0.1.0 // indirect
	github.com/alibabacloud-go/darabonba-encode-util v0.0.2 // indirect
	github.com/alibabacloud-go/darabonba-map v0.0.2 // indirect
	github.com/alibabacloud-go/darabonba-string v1.0.2 // indirect
	github.com/alibabacloud-go/debug v1.0.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.0 // indirect
	github.com/alibabacloud-go/tea v1.2.2 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.5 // indirect
	github.com/alibabacloud-go/tea-utils/v2 v2.0.5 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.712 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.5.1 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.9 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/apache/arrow/go/v14 v14.0.2 // indirect
	github.com/apache/arrow/go/v17 v17.0.0-20240503231747-7cd9c6fbd313 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/bufbuild/protocompile v0.13.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cockroachdb/errors v1.11.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.5.0 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.23.0 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/strfmt v0.23.0 // indirect
	github.com/go-openapi/validate v0.24.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.2 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/pprof v0.0.0-20240521024322-9665fa269a30 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/grafana/regexp v0.0.0-20220304095617-2e8d9baf4ac2 // indirect
	github.com/hashicorp/consul/api v1.28.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack/v2 v2.1.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.6 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.5.1 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hopeio/context v0.0.0-20240718075653-393bb05f2ba5 // indirect
	github.com/hopeio/protobuf v0.0.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lni/vfs v0.2.1-0.20220616104132-8852fd867376 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/markbates/going v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/miekg/dns v1.1.59 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mrjones/oauth v0.0.0-20190623134757-126b35219450 // indirect
	github.com/nats-io/nats.go v1.34.1 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/oapi-codegen/runtime v1.0.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/onsi/ginkgo/v2 v2.17.2 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/prometheus/prometheus v0.35.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/sagikazarmark/crypt v0.20.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/urfave/cli/v2 v2.27.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/valyala/histogram v1.2.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	github.com/xuri/efp v0.0.0-20240408161823-9ad904a10d6d // indirect
	github.com/xuri/nfp v0.0.0-20240318013403-ab9948c2c4a7 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.13 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.13 // indirect
	go.etcd.io/etcd/client/v2 v2.305.13 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/term v0.21.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.179.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
