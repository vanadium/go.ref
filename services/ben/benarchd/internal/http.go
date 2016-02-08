// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"v.io/x/ref/services/ben"
)

// NewHTTPHandler returns a handler that provides web interface for browsing
// benchmark results in store.
func NewHTTPHandler(store Store) http.Handler {
	return &handler{store}
}

type handler struct {
	store Store
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if id := strings.TrimSpace(r.FormValue("id")); len(id) > 0 {
		bm, itr := h.store.Runs(id)
		defer itr.Close()
		h.runs(w, bm, itr)
		return
	}
	if qstr := strings.TrimSpace(r.FormValue("q")); len(qstr) > 0 {
		query, err := ParseQuery(qstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			args := struct {
				Query string
				Error error
			}{qstr, err}
			executeTemplate(w, tmplBadQuery, args)
			return
		}
		h.handleQuery(w, query)
		return
	}
	if src := strings.TrimSpace(r.FormValue("s")); len(src) > 0 {
		h.describeSource(w, src)
		return
	}
	executeTemplate(w, tmplHome, nil)
}

func (h *handler) handleQuery(w http.ResponseWriter, query *Query) {
	bmarks := h.store.Benchmarks(query)
	defer bmarks.Close()
	if !bmarks.Advance() {
		executeTemplate(w, tmplNoBenchmarks, query)
		return
	}
	bm := bmarks.Value()
	// Advance once more, if there are more benchmarks (different names,
	// different scenarios, different uploaders) for the same query, then
	// present a list to choose from.
	if !bmarks.Advance() {
		itr := bmarks.Runs()
		defer itr.Close()
		h.runs(w, bm, itr)
		return
	}
	h.benchmarks(w, query.String(), bm, bmarks)
}

func (h *handler) benchmarks(w http.ResponseWriter, query string, first Benchmark, itr BenchmarkIterator) {
	var (
		cancel = make(chan struct{})
		items  = make(chan Benchmark, 2)
		errs   = make(chan error, 1)
	)
	defer close(cancel)
	items <- first
	items <- itr.Value()
	go func() {
		defer close(errs)
		defer close(items)
		for itr.Advance() {
			select {
			case items <- itr.Value():
			case <-cancel:
				return
			}
			if err := itr.Err(); err != nil {
				errs <- err
			}
		}
	}()
	args := struct {
		Query string
		Items <-chan Benchmark
		Err   <-chan error
	}{
		Query: query,
		Items: items,
		Err:   errs,
	}
	executeTemplate(w, tmplBenchmarks, args)
}

func (h *handler) runs(w http.ResponseWriter, bm Benchmark, itr RunIterator) {
	type item struct {
		Run          ben.Run
		SourceCodeID string
		UploadTime   time.Time
	}
	var (
		cancel = make(chan struct{})
		items  = make(chan item)
		errs   = make(chan error, 1)
	)
	defer close(cancel)
	go func() {
		defer close(errs)
		defer close(items)
		for itr.Advance() {
			r, s, t := itr.Value()
			select {
			case items <- item{r, s, t}:
			case <-cancel:
				return
			}
		}
		if err := itr.Err(); err != nil {
			errs <- err
		}
	}()
	args := struct {
		Benchmark Benchmark
		Items     <-chan item
		Err       <-chan error
	}{
		Benchmark: bm,
		Items:     items,
		Err:       errs,
	}
	executeTemplate(w, tmplRuns, args)
	return
}

func (h *handler) describeSource(w http.ResponseWriter, src string) {
	w.Header().Set("Content-Type", "text/plain")
	code, err := h.store.DescribeSource(src)
	if err != nil {
		fmt.Fprintf(w, "ERROR:%v", err)
		return
	}
	fmt.Fprintln(w, code)
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, args interface{}) {
	if err := tmpl.Execute(w, args); err != nil {
		fmt.Fprintf(w, "ERROR:%v", err)
	}

}

func newTemplate(name string, contents ...string) *template.Template {
	t := template.New(name)
	for _, c := range contents {
		t = template.Must(t.Parse(c))
	}
	return t
}

var (
	htmlStyling = `{{define "styling"}}
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://storage.googleapis.com/code.getmdl.io/1.0.6/material.indigo-pink.min.css">
<script src="https://storage.googleapis.com/code.getmdl.io/1.0.6/material.min.js"></script>
<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
{{end}}`
	htmlFooter = `{{define "footer"}}
<footer class="mdl-mini-footer">
  <div class="mdl-mini-footer__left-section">
    <ul class="mdl-mini-footer__link-list">
    <li><a href="/">Home</a></li>
    <li><a href="https://github.com/vanadium/go.ref/tree/master/services/ben">Github</a></li>
    </ul>
  </div>
</footer>
{{end}}`
	tmplHome = newTemplate(".home", `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Benchmark Results Archive</title>
    {{template "styling"}}
    <style>
    .fixed-width { font-family: monospace; }
    </style>
</head>
<body>
  <div class="mdl-layout mdl-js-layout">
    <main class="mdl-layout__content">
    <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--1-col"></div>
    <div class="mdl-cell mdl-cell--6-col">
    <form action="#">
        <div class="mdl-textfield mdl-js-textfield">
          <input class="mdl-textfield__input" type="text" id="q" name="q"/>
          <label class="mdl-textfield__label" for="q">Search query</label>
        </div>
        <input value="Search" type="submit" class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored"/>
    </form>
    Sample queries:
    <ul>
        <li><a href="/?q=VomEncode">VomEncode</a> - All benchmarks with names matching VomEncode</li>
        <li><a href="/?q=os:linux+cpu:amd64+v.io/v23/security">os:linux cpu:amd64 v.io/v23/security</a> - Benchmarks on desktop linux for the <span class="fixed-width">v.io/v23/security</span> package</li>
    </ul>
    Search operators:
    <ul>
       <li>os - Operating system, e.g., os:linux, os:&quot;Ubuntu 14.04&quot; etc.</li>
       <li>cpu - CPU, e.g., cpu:arm, cpu:xeon etc.</li>
       <li>uploader - Identity of uploader, e.g., uploader:janedoe</li>
       <li>label - Label assigned by the uploader, e.g., label:mylabel</li>
    </ul>
    </div>
    <div class="mdl-cell mdl-cell--1-col"></div>
    </div>
    </main>
    {{template "footer"}}
  </div>
</body>
</html>
`, htmlStyling, htmlFooter)
	tmplBadQuery = newTemplate(".badquery", `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>ERROR: Benchmark Results Archive</title>
    {{template "styling"}}
</head>
<body>
  <div class="mdl-layout mdl-js-layout">
    <main class="mdl-layout__content">
      <div class="mdl-grid">
      <div class="mdl-cell mdl-cell--1-col"></div>
      <div class="mdl-cell mdl-cell--6-col">
      <i class="material-icons">error</i>[{{.Query}}] is not valid: {{.Error}}
      </div>
      <div class="mdl-cell mdl-cell--1-col"></div>
      </div>
    </main>
    {{template "footer"}}
  </div>
</body>
</html>
`, htmlStyling, htmlFooter)
	tmplNoBenchmarks = newTemplate(".nobenchmarks", `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Benchmark Results Archive</title>
    {{template "styling"}}
</head>
<body>
  <div class="mdl-layout mdl-js-layout">
    <main class="mdl-layout__content">
    <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--1-col"></div>
    <div class="mdl-cell mdl-cell--6-col">
    <i class="material-icons">info</i>No results for [{{.}}]
    </div>
    <div class="mdl-cell mdl-cell--1-col"></div>
    </div>
    </main>
    {{template "footer"}}
  </div>
</body>
</html>
`, htmlStyling, htmlFooter)
	tmplBenchmarks = newTemplate(".benchmarks", `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Benchmark Results Archive</title>
    {{template "styling"}}
</head>
<body class="mdl-demo mdl-color--grey-100 mdl-color-text--grey-700 mdl-base">
  <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <main class="mdl-layout__content">
      <section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
        <div class="mdl-card__supporting-text">
        <form action="#">
            <div class="mdl-textfield mdl-js-textfield">
              <input class="mdl-textfield__input" type="text" id="q" name="q" value="{{.Query}}"/>
              <label class="mdl-textfield__label" for="q">{{.Query}}</label>
            </div>
            <input value="Search" type="submit" class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored"/>
        </form>
        </div>
      </section>
      <section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
      <div class="mdl-card mdl-cell mdl-cell--12-col">
        <div class="mdl-card__supporting-text">
          <h4>Benchmarks</h4>
          <table class="mdl-data-table mdl-js-data-table mdl-shadow--2dp">
            <thead>
            <tr>
              <th>Name</th>
              <th title="nanoseconds per iteration">ns/op</th>
              <th title="operating system on which benchmarks were run">OS</th>
              <th title="CPU architecture on which benchmarks were run">CPU</th>
              <th title="who uploaded results for this benchmark">Uploader</th>
              <th>Label</th>
            </tr>
            </thead>
            <tbody>
                {{range .Items}}
                <tr>
                <td><a href="/?id={{.ID | urlquery}}">{{.Name}}</a></td>
                <td>{{.NanoSecsPerOp}}</td>
                <td><div id="os_{{.ID}}">{{.Scenario.Os.Name}}</div><div class="mdl-tooltip" for="os_{{.ID}}">{{.Scenario.Os.Version}}</div></td>
                <td><div id="cpu_{{.ID}}">{{.Scenario.Cpu.Architecture}}</div><div class="mdl-tooltip" for="cpu_{{.ID}}">{{.Scenario.Cpu.Description}}</div></td>
                <td>{{.Uploader}}</td>
                <td>{{.Scenario.Label}}</td>
                </tr>
                {{end}}
                {{range .Err}}
                <tr><td colspan=5><i class="material-icons">error</i>{{.}}</td></tr>
                {{end}}
            </tbody>
          </table>
        </div>
      </div>
      </section>
    </main>
    {{template "footer"}}
  </div>
</body>
</html>
`, htmlStyling, htmlFooter)
	tmplRuns = newTemplate(".runs", `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Benchmark Results Archive</title>
    {{template "styling"}}
    <style>
    .fixed-width { font-family: monospace; }
    </style>
</head>
<body class="mdl-demo mdl-color--grey-100 mdl-color-text--grey-700 mdl-base">
  <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <main class="mdl-layout__content">
      {{with .Benchmark}}
      <section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
      <div class="mdl-card mdl-cell mdl-cell--12-col">
        <div class="mdl-card__supporting-text">
          <h4>{{.Name}}</h4>
          <table class="mdl-data-table mdl-js-data-table mdl-shadow--2dp">
            <tbody>
              <tr>
                <td class="mdl-data-table__cell--non-numeric">OS</td>
                <td class="mdl-data-table__cell--non-numeric">{{.Scenario.Os.Name}} ({{.Scenario.Os.Version}})</td>
              </tr>
              <tr>
                <td class="mdl-data-table__cell--non-numeric">CPU</td>
                <td class="mdl-data-table__cell--non-numeric">{{.Scenario.Cpu.Architecture}} ({{.Scenario.Cpu.Description}})</td>
              </tr>
              <tr>
                <td class="mdl-data-table__cell--non-numeric">Uploader</td>
                <td class="mdl-data-table__cell--non-numeric">{{.Uploader}}</td>
              </tr>
              {{with .Scenario.Label}}
              <tr>
                <td class="mdl-data-table__cell--non-numeric">Label</td>
                <td class="mdl-data-table__cell--non-numeric">{{.}}</td>
              </tr>
              {{end}}
            </tbody>
          </table>
        </div>
      </div>
      </section>
      {{end}}

      <section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
      <div class="mdl-card mdl-cell mdl-cell--12-col">
        <div class="mdl-card__supporting-text">
          <h4>Runs</h4>
          <table class="mdl-data-table mdl-js-data-table mdl-shadow--2dp fixed-width">
            <thead>
            <tr>
              <th title="nanoseconds per iteration">ns/op</th>
              <th title="number of memory allocations per iteration">allocs/op</th>
              <th title="number of bytes of memory allocations per iteration">allocated bytes/op</th>
              <th title="megabytes processed per second">MB/s</th>
              <th title="timestamp when results were uploaded" class="mdl-data-table__cell--non-numeric">Uploaded</th>
              <th title="description of source code" class="mdl-data-table__cell--non-numeric">SourceCode</th>
              <th>Iterations</th>
              <th title="parallelism used to run benchmark, e.g., GOMAXPROCS for Go benchmarks">Parallelism</th>
            </tr>
            </thead>
            <tbody>
              {{range .Items}}
              <tr>
                <td>{{.Run.NanoSecsPerOp}}</td>
                <td>{{if .Run.AllocsPerOp}}{{.Run.AllocsPerOp}}{{end}}</td>
                <td>{{if .Run.AllocedBytesPerOp}}{{.Run.AllocedBytesPerOp}}{{end}}</td>
                <td>{{if .Run.MegaBytesPerSec}}{{.Run.MegaBytesPerSec}}{{end}}</td>
                <td>{{.UploadTime}}</td>
                <td><a href="?s={{urlquery .SourceCodeID}}">(sources)</a></td>
                <td>{{.Run.Iterations}}</td>
                <td>{{.Run.Parallelism}}</td>
              </tr>
              {{end}}
              {{range .Err}}
              <tr><td colspan=8><i class="material-icons">error</i>{{.}}</td></tr>
              {{end}}
            </tbody>
          </table>
        </div>
      </div>
      </section>
    </main>
    {{template "footer"}}
  </div>
</body>
</html>
`, htmlStyling, htmlFooter)
)
