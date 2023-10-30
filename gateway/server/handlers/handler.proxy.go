package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxyHandler(targetURL string) http.Handler {
    target, err := url.Parse(targetURL)
    if err != nil {
        
    }
    proxy := httputil.NewSingleHostReverseProxy(target)

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            originalPath := r.URL.Path
            println("Proxying request:", originalPath)
            proxy.ServeHTTP(w, r)
    })
}