package core

import (
    "github.com/gorilla/sessions"
    "net/http"
)

var session_store sessions.Store
var session *sessions.Session

func enter_site(r *http.Request, name string) {
    session_store = sessions.NewCookieStore([]byte(config.Sites[name].SessionEntropy))
    session, _ = session_store.Get(r, "sakubun")
}

func exit_site(w http.ResponseWriter, r *http.Request) {
    session.Save(r, w)
}

func SiteHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        enter_site(r, config_site)
		defer exit_site(w, r)
		h.ServeHTTP(w, r)
	})
}