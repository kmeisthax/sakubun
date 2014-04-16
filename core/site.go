package core

import "github.com/gorilla/sessions"

var session_store Store

func enter_site(r *http.Request, name string) {
    session_store = sessions.NewCookieStore(r, []byte(config.Sites[name].SessionEntropy))
}

func exit_site(w http.ResponseWriter, r *http.Request) {
    session_store.Save(r, w)
}