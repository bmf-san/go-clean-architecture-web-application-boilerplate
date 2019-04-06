package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"../repository"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	// TODO: 以下のメソッド方針で実装
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

	// TODO: ルーティング処理は切り出したい controllerとhandlerは別・・？ handlerにできるならhandlerにリネームした
	// TODO: ここが処理のメイン userrepositoryはどっかの引数に取るとかしても良いのでは
	users, err := repository.UserRepositoryIndex(uh.DB)
	if err != nil {
		log.Fatalf("show users list error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
