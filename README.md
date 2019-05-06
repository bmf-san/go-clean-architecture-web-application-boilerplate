# go-api-boilerplate
This is an api boilerplate built with golang.

- golang
- mysql
- docker

# Get Started
`cp app/.env_example app/.env`
`docker-compose build`
`docker-compose up`

After running docker, you need to execute sql files in `app/database/sql`.

# Architecture
```
src
├── database                      ･･･ Out of layer packages
│   ├── migrations
│   │   └── schema.sql
│   └── seeds
│       └── faker.sql
├── domain                        ･･･ Entities Layer
│   └── user.go
├── infrastructure                ･･･ Frameworks & Drivers Layer
│   ├── database                  
│   │   └── sqlhandler.go
│   ├── router
│   │   └── router.go
│   └── env 
│       └── env.go
├── interfaces                    ･･･ Interface Layer
│   ├── controllers
│   │   └── user_controller.go
│   ├── database
│   │   └── sqlhandler.go
│   └── repositories
│       └── user_repository.go
├── main.go
└── usecases                      ･･･ Usecases Layer
    └── repositories
        └── user_repository.go
```

| Layer                | Directory      | Sub directory                        |
|----------------------|----------------|--------------------------------------|
| Frameworks & Drivers | infrastructure | database, router, env                |
| Interface            | interfaces     | controllers, database, repositories  |
| Usecases             | usecases       | repositories                         |
| Entities             | domain         | ex. user.go                          |

# API Document
// TODO


# Controller method naming rule
- index
  - Display a listing of the resource.
- show
  - Display the specified resource.
- store
  - Store a newly created resource in storage.
- update
  - Update the specified resource in storage.
- destory
  - Remove the specified resource from storage.

This naming rule is referring to laravel's rest api methods.
cf. [laravel](https://laravel.com/)

# References
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)