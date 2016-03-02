// Copyright 2016 Google Inc. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package onetwoseven is a simple web server.
package onetwoseven

import (
	"fmt"
	"net/http"

	"github.com/aliafshar/toylog"
)

type Config struct {
	Path     string
	HostPort string
}

func (c *Config) String() string {
	return fmt.Sprintf("%v on %v", c.Path, c.HostPort)
}

type server struct {
	config *Config
}

type wrappingHandler struct {
	pre     func(http.ResponseWriter, *http.Request)
	post    func(http.ResponseWriter, *http.Request)
	handler http.Handler
}

func (h *wrappingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.pre != nil {
		h.pre(w, r)
	}
	h.handler.ServeHTTP(w, r)
	if h.post != nil {
		h.post(w, r)
	}
}

func (s *server) FileServer() http.Handler {
	h := http.FileServer(http.Dir(s.config.Path))
	return &wrappingHandler{
		handler: h,
		post: func(w http.ResponseWriter, r *http.Request) {
			log(w, r)
		},
	}
}

func (s *server) start() {
	http.Handle("/", s.FileServer())
	http.ListenAndServe(s.config.HostPort, nil)
}

func log(w http.ResponseWriter, r *http.Request) {
	toylog.Infof("%s %s %s", r.Method, r.URL, w.Header())
}

func Serve(config *Config) {
	s := &server{config: config}
	s.start()
}
