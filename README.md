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

# API Document
// TODO
// list	一覧を取得する
// get	IDを指定して、一件取得する
// insert	新規作成する
// update	上書きする
// patch	部分的に上書きする
// delete	指定したIDに対応するものを削除する


// TODO: method 振り分け
// switch r.Method {
// 	case http.MethodPost:
//     	Register(w, r)
// 	case http.MethodGet:
//     	Reader(w, r)
// 	case http.MethodPut:
//     	Updater(w, r)
// 	case http.MethodDelete:
//     	Deleter(w, r)
// 	default:
//     	NotFoundResources(w, r)

# References
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)