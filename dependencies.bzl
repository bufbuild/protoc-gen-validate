load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_third_party():
    go_repository(
        name = "com_github_andybalholm_cascadia",
        importpath = "github.com/andybalholm/cascadia",
        sum = "h1:BuuO6sSfQNFRu1LppgbD25Hr2vLYW25JvxHs5zzsLTo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:4daAzAu0S6Vi7/lbWECcX0j45yZReDZ56BQsrVBOEEY=",
        version = "v0.0.0-20200428143746-21a406dcc535",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:D9otznteZZyN5pRyFETqveYia/85Xzk7+RaPGB1I9fE=",
        version = "v1.34.20",
    )
    go_repository(
        name = "com_github_aymerick_douceur",
        importpath = "github.com/aymerick/douceur",
        sum = "h1:Mv+mAeH1Q+n9Fr+oyamOlAkUNPWPlA8PPGR0QAaYuPk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_chris_ramon_douceur",
        importpath = "github.com/chris-ramon/douceur",
        sum = "h1:IDMEdxlEUUBYBKE4z/mJnFyVXox+MjuEVDJNN27glkU=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        importpath = "github.com/denisenkom/go-mssqldb",
        sum = "h1:83Wprp6ROGeiHFAP8WJdI2RoxALQYgdllERc3N5N2DM=",
        version = "v0.0.0-20191124224453-732737034ffd",
    )
    go_repository(
        name = "com_github_disintegration_imaging",
        importpath = "github.com/disintegration/imaging",
        sum = "h1:w1LecBlG2Lnp8B3jk5zSuNqd7b4DXhcjwek1ei82L+c=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_erikstmartin_go_testdb",
        importpath = "github.com/erikstmartin/go-testdb",
        sum = "h1:Yzb9+7DPaBjB8zlTR87/ElzFsnQfuHnVUVqpZZIcV5Y=",
        version = "v0.0.0-20160219214506-8d10e4a1bae5",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:8xPHl4/q1VyqGIPif1F+1V3Y3lSmrq01EabUW3CoW5s=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
        version = "v1.5.0",
    )

    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:lXe2qZdvpiX5WZkZR4hgp4KJVfY3nMkvmwbVkpv1rVY=",
        version = "v0.0.0-20190719163853-cb61b32ac6fe",
    )

    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_gorilla_context",
        importpath = "github.com/gorilla/context",
        sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_css",
        importpath = "github.com/gorilla/css",
        sum = "h1:BQqNyPTi50JCFMTw/b67hByjMVXZRwGha6wxVGkeihY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:S7P+1Hm5V/AT9cjEcUD5uDaQSX0OE577aCXgoaKpYbQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_gosimple_slug",
        importpath = "github.com/gosimple/slug",
        sum = "h1:r5vDcYrFz9BmfIAMC829un9hq7hKM4cHUrsv36LbEqs=",
        version = "v1.9.0",
    )

    go_repository(
        name = "com_github_iancoleman_strcase",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_jinzhu_configor",
        importpath = "github.com/jinzhu/configor",
        sum = "h1:u78Jsrxw2+3sGbGMgpY64ObKU4xWCNmNRJIjGVqxYQA=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_jinzhu_gorm",
        importpath = "github.com/jinzhu/gorm",
        sum = "h1:OdR1qFvtXktlxk73XFYMiYn9ywzTwytqe4QkuMRqc38=",
        version = "v1.9.15",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:g39TucaRWyV3dwDO++eEc6qf8TVIQ/Da48WmqjZ3i7E=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:OS12ieG61fsCg5+qLJ+SsW9NicxNkg3b25OyT2yCeUc=",
        version = "v0.3.0",
    )

    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:9xohqzkUwzR4Ga4ivdTcawVS89YSDVxXMa3xJX3cGzg=",
        version = "v1.8.0",
    )

    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:xOpFu4vwmIoUeUrRuAtdCrZZymT/6AkW/bsUWA506Fo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:snbPLB8fVfU9iwbbo30TPtbLRzwWu6aJS6Xh4eaaviA=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:FxPOTFNqGkuDUGi3H/qkUbQO4ZiBa2brKq5r0l8TGeM=",
        version = "v0.0.11",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:mLyGNKR8+Vv9CAU7PphKa2hkEqxxhn8i32J6FPj1/QA=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:EjVH7OqbU219kdm8acbveoclh2zZFqPJTJw6VUlTLAQ=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )

    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:VasscCm72135zRysgrJDKsntdmPN+OuU3+nnHYA9wyc=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_puerkitobio_goquery",
        importpath = "github.com/PuerkitoBio/goquery",
        sum = "h1:PSPBGne8NIUWw+/7vFBV+kG2J/5MOjbzc7154OaKCSE=",
        version = "v1.5.1",
    )

    go_repository(
        name = "com_github_qor_admin",
        importpath = "github.com/qor/admin",
        sum = "h1:tveOmg4NXBMyDyWqLx8xY/LdDMo0r1Tvh4P74BdjKbo=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_qor_assetfs",
        importpath = "github.com/qor/assetfs",
        sum = "h1:JRpyNNSRAkwNHd4WgyPcalTAhxOCh3eFNMoQkxWhjSw=",
        version = "v0.0.0-20170713023933-ff57fdc13a14",
    )
    go_repository(
        name = "com_github_qor_audited",
        importpath = "github.com/qor/audited",
        sum = "h1:87fy9oxrDUTN77K9VwHkmyNynRYl+Kn7HkX2HZbQYrE=",
        version = "v0.0.0-20171228121055-b52c9c2f0571",
    )

    go_repository(
        name = "com_github_qor_cache",
        importpath = "github.com/qor/cache",
        sum = "h1:XBEqqX1xirq3pYMGw8Jr0urCfhPXbJEh1KPwKLd24MM=",
        version = "v0.0.0-20171031031927-c9d48d1f13ba",
    )
    go_repository(
        name = "com_github_qor_i18n",
        importpath = "github.com/qor/i18n",
        sum = "h1:7tqz0lsahPvTvFuIeeW8VrwmElrOzrE89hOQc38VWCk=",
        version = "v0.0.0-20210601022951-0f75814734d3",
    )
    go_repository(
        name = "com_github_qor_l10n",
        importpath = "github.com/qor/l10n",
        sum = "h1:zMwbO23nfrcv6IqPbY4KknT7RqhDO6sg7VaV5WoyujA=",
        version = "v0.0.0-20181031091737-2ca95fb3b4dd",
    )
    go_repository(
        name = "com_github_qor_media",
        importpath = "github.com/qor/media",
        sum = "h1:LgWyG2gPmfiEukIo0yjlQ93oWWyTbmvX04rcOs4SDtk=",
        version = "v0.0.0-20200720100650-60c52edf57cb",
    )
    go_repository(
        name = "com_github_qor_middlewares",
        importpath = "github.com/qor/middlewares",
        sum = "h1:+WCc1IigwWpWBxMFsmLUsIF230TakGHstDajd8aKDAc=",
        version = "v0.0.0-20170822143614-781378b69454",
    )
    go_repository(
        name = "com_github_qor_oss",
        importpath = "github.com/qor/oss",
        sum = "h1:J2Xj92efYLxPl3BiibgEDEUiMsCBzwTurE/8JjD8CG4=",
        version = "v0.0.0-20191031055114-aef9ba66bf76",
    )
    go_repository(
        name = "com_github_qor_publish",
        importpath = "github.com/qor/publish",
        sum = "h1:U+g8WMSHCibub/ZnX5gIG8SKk8d1UA0Ppv863vGhfoE=",
        version = "v0.0.0-20181014061411-abfbacee9e0d",
    )
    go_repository(
        name = "com_github_qor_publish2",
        importpath = "github.com/qor/publish2",
        sum = "h1:yPZKRv1lxMYphCyWyzUYvqXH8NvSNr8O/LwNeSONp8E=",
        version = "v0.0.0-20200729081509-d97fdb5620a5",
    )

    go_repository(
        name = "com_github_qor_qor",
        importpath = "github.com/qor/qor",
        sum = "h1:13GClzwPmVNPHjUXAuZiVX6pOea8JXyr6mES0zuP+Uc=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_qor_responder",
        importpath = "github.com/qor/responder",
        sum = "h1:sKELSAyL+z5BRHFe97Bx71z197cBEobVJ6rASKTMSqU=",
        version = "v0.0.0-20171031032654-b6def473574f",
    )
    go_repository(
        name = "com_github_qor_roles",
        importpath = "github.com/qor/roles",
        sum = "h1:F0BNcPJKfubM/+IIILu/GbrH9v2vPZWQ5/StSRKUfK4=",
        version = "v0.0.0-20171127035124-d6375609fe3e",
    )
    go_repository(
        name = "com_github_qor_serializable_meta",
        importpath = "github.com/qor/serializable_meta",
        sum = "h1:dnqlvo4M/uR0KhB6Tyhsv6XbSokcpfwQ4ublD9D/PBQ=",
        version = "v0.0.0-20180510060738-5fd8542db417",
    )
    go_repository(
        name = "com_github_qor_session",
        importpath = "github.com/qor/session",
        sum = "h1:8l21EEdlZ9R0AA3FbeUAANc5NAx8Y3tn1VKbyAgjYlI=",
        version = "v0.0.0-20170907035918-8206b0adab70",
    )
    go_repository(
        name = "com_github_qor_sorting",
        importpath = "github.com/qor/sorting",
        sum = "h1:uic2CBCjBtGSwss9NiknoB52PLUZustFRjXL3tkOTCE=",
        version = "v0.0.0-20200724034229-cdba739ba535",
    )

    go_repository(
        name = "com_github_qor_validations",
        importpath = "github.com/qor/validations",
        sum = "h1:dRlsVUhwD1pwrasuVbNooGQITYjKzmXK5eYoEEvBGQI=",
        version = "v0.0.0-20171228122639-f364bca61b46",
    )
    go_repository(
        name = "com_github_qor_worker",
        importpath = "github.com/qor/worker",
        sum = "h1:EUo1U/EGuZ/CQeLvdQoqwRA/lxcf0rIiZ/Cc101+hV4=",
        version = "v0.0.0-20190805090529-35a245417f70",
    )
    go_repository(
        name = "com_github_rainycape_unidecode",
        importpath = "github.com/rainycape/unidecode",
        sum = "h1:ta7tUOvsPHVHGom5hKW5VXNc2xZIkfCKP8iaqOyYtUQ=",
        version = "v0.0.0-20150907023854-cb7f23ec59be",
    )

    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:xoax2sJ2DT8S8xA2paPFjDCScCNeWsg75VG0DLRreiY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:hDPOHmpOpP40lSULcqw7IrRb/u7w6RpDC9399XyoNd0=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_theplant_cldr",
        importpath = "github.com/theplant/cldr",
        sum = "h1:di0cR5qqo2DllBMwmP75kZpUX6dAXhsn1O2dshQfMaA=",
        version = "v0.0.0-20190423050709-9f76f7ce4ee8",
    )
    go_repository(
        name = "com_github_theplant_htmltestingutils",
        importpath = "github.com/theplant/htmltestingutils",
        sum = "h1:yPrgtU8bj7Q/XbXgjjmngZtOhsUufBAraruNwxv/eXM=",
        version = "v0.0.0-20190423050759-0e06de7b6967",
    )
    go_repository(
        name = "com_github_theplant_testingutils",
        importpath = "github.com/theplant/testingutils",
        sum = "h1:757/ruZNgTsOf5EkQBo0i3Bx/P2wgF5ljVkODeUX/uA=",
        version = "v0.0.0-20190603093022-26d8b4d95c61",
    )
    go_repository(
        name = "com_github_yosssi_gohtml",
        importpath = "github.com/yosssi/gohtml",
        sum = "h1:YWaOkupKL+BRRJSWRq/uhSkWXc1K0QVIYVG36XUBGOc=",
        version = "v0.0.0-20200519115854-476f5b4b8047",
    )

    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:dPmz1Snjq0kmkz159iL7S6WzdahUTHnHB5M56WFVifs=",
        version = "v1.3.5",
    )

    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:QRR6H1YWRnHb4Y/HeNFCTJLFVxaq6wH4YuVdsUOr75U=",
        version = "v1.0.0-20200902074654-038fdea0a05b",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
        version = "v2.2.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
        version = "v3.0.0-20200313102051-9f266ea9e77c",
    )

    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:bxAC2xTBsZGibn2RTntX0oH50xLsqy1OxA9tTL3p/lk=",
        version = "v1.27.1",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:GGJVjV8waZKRHrgwvtH66z9ZGVurTD1MT0n1Bb+q4aM=",
        version = "v0.0.0-20191205180655-e7c4368fe9dd",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:hVwzHzIUGRjiF7EcUjqNxk3NCfkPxbDKRdnNE1Rpg0U=",
        version = "v0.0.0-20191009234506-e7c1f5e7dbb8",
    )

    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:UG21uOlmZabA4fW5i7ZX6bjw1xELEGg/ZLgZq9auk/Q=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:LO7XpTYMwTqxjLcGWPijK3vRXg1aWdlNOVOHRq45d7c=",
        version = "v0.0.0-20210813160813-60bc85c4be6d",
    )

    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
        version = "v0.0.0-20210220032951-036812b2e83c",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:uCLL3g5wH2xjxVREVuAbP9JM5PPKjRbXKRa6IBjkzmU=",
        version = "v0.0.0-20210816183151-1e6c022a8912",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
        version = "v0.0.0-20201126162022-7de9c90e9dd1",
    )

    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
        version = "v0.3.7",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:ouewzE6p+/VEB31YYnTbEJdi8pFqKp4P4n85vwo3DHA=",
        version = "v0.1.5",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
