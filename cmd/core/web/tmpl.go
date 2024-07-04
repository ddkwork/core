// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"text/template"

	"cogentcore.org/core/cmd/core/config"
)

// appJSTmpl is the template used in [makeAppJS] to build the app.js file
var appJSTmpl = template.Must(template.New("app.js").Parse(appJS))

// appJSData is the data passed to [appJSTmpl]
type appJSData struct {
	Env                     string
	WasmContentLengthHeader string
	AutoUpdateInterval      int64
}

// makeAppJS exectues [appJSTmpl] based on the given configuration information.
func makeAppJS(c *config.Config) ([]byte, error) {
	if c.Web.Env == nil {
		c.Web.Env = make(map[string]string)
	}
	c.Web.Env["GOAPP_STATIC_RESOURCES_URL"] = "/"
	c.Web.Env["GOAPP_ROOT_PREFIX"] = "."

	for k, v := range c.Web.Env {
		if err := os.Setenv(k, v); err != nil {
			slog.Error("setting app env variable failed", "name", k, "value", "err", err)
		}
	}

	wenv, err := json.Marshal(c.Web.Env)
	if err != nil {
		return nil, err
	}

	d := appJSData{
		Env:                     string(wenv),
		WasmContentLengthHeader: c.Web.WasmContentLengthHeader,
		AutoUpdateInterval:      c.Web.AutoUpdateInterval.Milliseconds(),
	}
	b := &bytes.Buffer{}
	err = appJSTmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// appWorkerJSData is the data passed to [config.Config.Web.ServiceWorkerTemplate]
type appWorkerJSData struct {
	Version          string
	ResourcesToCache string
}

// makeAppWorkerJS executes [config.Config.Web.ServiceWorkerTemplate]. If it empty, it
// sets it to [appWorkerJS].
func makeAppWorkerJS(c *config.Config) ([]byte, error) {
	resources := []string{
		"app.css",
		"app.js",
		"app.wasm",
		"manifest.webmanifest",
		"wasm_exec.js",
		"index.html",
	}

	tmpl, err := template.New("app-worker.js").Parse(appWorkerJS)
	if err != nil {
		return nil, err
	}

	rstr, err := json.Marshal(resources)
	if err != nil {
		return nil, err
	}

	d := appWorkerJSData{
		Version:          c.Version,
		ResourcesToCache: string(rstr),
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// manifestJSONTmpl is the template used in [makeManifestJSON] to build the mainfest.webmanifest file
var manifestJSONTmpl = template.Must(template.New("manifest.webmanifest").Parse(manifestJSON))

// manifestJSONData is the data passed to [manifestJSONTmpl]
type manifestJSONData struct {
	ShortName   string
	Name        string
	Description string
}

// makeManifestJSON exectues [manifestJSONTmpl] based on the given configuration information.
func makeManifestJSON(c *config.Config) ([]byte, error) {
	d := manifestJSONData{
		ShortName:   c.Name,
		Name:        c.Name,
		Description: c.About,
	}

	b := &bytes.Buffer{}
	err := manifestJSONTmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// indexHTMLTmpl is the template used in [makeIndexHTML] to build the index.html file
var indexHTMLTmpl = template.Must(template.New("index.html").Parse(indexHTML))

// indexHTMLData is the data passed to [indexHTMLTmpl]
type indexHTMLData struct {
	BasePath               string
	Author                 string
	Desc                   string
	Keywords               []string
	Title                  string
	SiteName               string
	Image                  string
	VanityURL              string
	GithubVanityRepository string
}

// makeIndexHTML exectues [indexHTMLTmpl] based on the given configuration information,
// base path for app resources (used in [MakePages]), and optional title (used in [MakePages],
// defaults to [config.Config.Name] otherwise).
func makeIndexHTML(c *config.Config, basePath string, title string) ([]byte, error) {
	if title == "" {
		title = c.Name
	}
	d := indexHTMLData{
		BasePath:               basePath,
		Author:                 c.Web.Author,
		Desc:                   c.About,
		Keywords:               c.Web.Keywords,
		Title:                  title,
		SiteName:               c.Name,
		Image:                  c.Web.Image,
		VanityURL:              c.Web.VanityURL,
		GithubVanityRepository: c.Web.GithubVanityRepository,
	}

	b := &bytes.Buffer{}
	err := indexHTMLTmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
