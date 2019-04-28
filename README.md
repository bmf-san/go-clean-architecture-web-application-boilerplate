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
app/
├── database
│   ├── migrations
│   │   └── schema.sql
│   └── seeds
│       └── faker.sql
├── domain
│   ├── post.go
│   └── user.go
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── post_controller.go
│   ├── post_repository.go
│   ├── sqlhandler.go
│   ├── user_controller.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
├── main.go
├── test.log
└── usecases
    ├── logger.go
    ├── post_interactor.go
    ├── post_repository.go
    ├── user_interactor.go
    └── user_repository.go
```

| Layer                | Directory      | Sub directory                        |
|----------------------|----------------|--------------------------------------|
| Frameworks & Drivers | infrastructure | database, router, env                |
| Interface            | interfaces     | controllers, database, repositories  |
| Usecases             | usecases       | repositories                         |
| Entities             | domain         | ex. user.go                          |

# API Document
// TODO: This section will be updated the date when ver.3.0.0 will be released.


# Controller method naming rule
Use these word as prefix.

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

# Repository method naming rule
Use these word as prefix.

- get
- add
- update
- delete

# Contribution

# References
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)
