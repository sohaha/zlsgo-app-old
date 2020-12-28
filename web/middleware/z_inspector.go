package middleware

import (
	"errors"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
)

type (
	Pagination struct {
		Total       int           `json:"total"`
		TotalPage   int           `json:"total_page"`
		CurrentPage int           `json:"current_page"`
		PerPage     int           `json:"per_page"`
		HasNext     bool          `json:"has_next"`
		HasPrev     bool          `json:"has_prev"`
		NextPageUrl string        `json:"next_page_url"`
		PrevPageUrl string        `json:"prev_page_url"`
		Data        []RequestStat `json:"data"`
	}
	RequestStat struct {
		RequestedAt   time.Time   `json:"requested_at"`
		RequestUrl    string      `json:"request_url"`
		HttpMethod    string      `json:"http_method"`
		HttpStatus    int         `json:"http_status"`
		ContentType   string      `json:"content_type"`
		GetParams     interface{} `json:"get_params"`
		PostParams    interface{} `json:"post_params"`
		PostMultipart interface{} `json:"post_multipart"`
		ClientIP      string      `json:"client_ip"`
		Cookies       interface{} `json:"cookies"`
		Headers       interface{} `json:"headers"`
		// Content      []byte   `json:"content"`
		Content        interface{} `json:"content"`
		ProcessingTime string      `json:"processing_time"`
		Raw            string
	}
	AllRequests struct {
		Requets []RequestStat `json:"requests"`
	}
)

const (
	tmp = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <title>{{.title}}</title>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/css/bootstrap.min.css" integrity="sha256-aAr2Zpq8MZ+YA/D6JtRD3xtrwpEz2IqOS+pWD/7XKIw=" crossorigin="anonymous">
  <style>
    body {
      font-family: 'Poppins', sans-serif;
      font-size: 15px;
      background: #f0f5fb !important;
      -webkit-font-smoothing: antialiased;
      text-rendering: optimizeLegibility;
      -moz-osx-font-smoothing: grayscale;
      font-weight: 300;
      overflow-y: scroll;
    }
    .navbar {
      background-color: #37474f;
    }
    .table thead th {
      text-transform: uppercase;
      font-size: 13px;
      font-weight: 500;
      color: #607d8b;
    }
    .table td, .table th {
      border-top: none;
    }
    .table tr th {
      font-weight: 500;
    }
    .table thead th, .table tr {
      border-bottom: 1px solid #eee;
    }
    .table tr:last-child {
      border-bottom: none;
    }
    .badge {
      padding: 6px 12px;
      font-size: 13px;
      font-weight: 500;
    }
    .btn-detail {
      font-size: 13px;
      font-weight: 500;
      padding: 5px 20px;
      background: #00bcd4;
      border: none;
    }
    .btn-detail:hover, .btn-detail:active {
      background-color: #4dd0e1 !important;
    }
    .shadow-sm {
      box-shadow: 0 .125rem 1.25rem rgba(0, 0, 0, .02) !important;
    }
    .page-link {
      border-color: #01bcd4;
      color: #01bcd4;
    }
    .page-link:hover {
      background-color: #00bcd4;
      color: #fff;
    }
    .nav-tabs {
      border: none;
    }
    .nav-tabs .nav-link {
      border-radius: 40px;
      background-color: #eee;
      color: #999;
      border: none;
      font-size: 15px;
      font-weight: 500;
      padding: 10px 30px;
      margin-right: 10px;
    }
    .nav-tabs .nav-link.active {
      background-color: #00bcd4;
      color: #fff;
    }
    .tab-content>.active,.rounded {
        overflow: auto;
    }
    @media (min-width: 576px) {
      .modal-dialog {
        max-width: none !important;
      }
    }
  </style>
</head>
<body class="bg-light">
<nav class="navbar navbar-dark">
  <a class="navbar-brand mx-auto" href="{{.path}}">
    <div style="display: inline-block;margin-top: 10px;margin-left: 10px;">{{.title}}</div>
  </a>
</nav>
<main role="main" class="container">
  <div class="my-3 p-4 bg-white rounded shadow-sm">
    <table class="table m-0">
      <thead>
      <tr>
        <th scope="col">Method</th>
        <th scope="col">Url</th>
        <th scope="col">Status</th>
        <th scope="col">Date</th>
        <th scope="col">Processing</th>
        <th scope="col">Inspect</th>
      </tr>
      </thead>
      <tbody>
      {{ range $i,$value :=.pagination.Data }}
        <tr>
          <th scope="row">{{ $value.HttpMethod }}</th>
          <td><span>{{$value.RequestUrl}}</span></td>
          <td>
            <span class="badge badge-secondary badge-{{$value.HttpStatus}}">{{ $value.HttpStatus }}</span>
          </td>
          <td>{{$value.RequestedAt | formatDate }}</td>
          <td>{{ $value.ProcessingTime }}</td>
          <td>
            <button type="button" class="btn btn-primary btn-detail" data-toggle="modal" data-target="#modal-{{$i}}">
              Detail
            </button>
            <div class="modal fade" id="modal-{{$i}}" tabindex="-1" role="dialog" aria-labelledby="exampleModalLongTitle" aria-hidden="true">
              <div class="container">
                <div class="modal-dialog" role="document">
                  <div class="modal-content">
                    <div class="modal-header pl-4 pr-4">
                      <h5 class="modal-title" id="exampleModalLongTitle">{{.RequestUrl}}
                        <span class="badge badge-secondary badge-{{.HttpStatus}}">{{.HttpStatus}}</span>
                      </h5>
                      <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                      </button>
                    </div>
                    <div class="modal-body p-4">
                      <ul class="nav nav-tabs mb-4">
                        <li class="nav-item">
                          <a class="nav-link active" data-toggle="tab" href
="#modal-{{$i}}-tab1">Response</a>
                        </li>
                        <li class="nav-item">
                          <a class="nav-link" data-toggle="tab" href="#modal-{{$i}}-tab2">Parameters</a>
                        </li>
                        <li class="nav-item">
                          <a class="nav-link" data-toggle="tab" href="#modal-{{$i}}-tab3">Headers</a>
                        </li>
                      </ul>
                      <div class="tab-content">
                        <div id="modal-{{$i}}-tab1" class="tab-pane active">
                          <h3>Response</h3>
                          <code>
                            {{ .Content }}
                          </code>
                        </div>
                        <div id="modal-{{$i}}-tab2" class="tab-pane fade">
                            {{ if .GetParams }}
                              <h3>Query Parameters</h3>
                              <table class="table table-hover">
                                <thead>
                                <tr>
                                  <th scope="col">Key</th>
                                  <th scope="col">Value</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{ range $key, $value :=  .GetParams}}
                                  <tr>
                                    <th scope="row">{{$key}}</th>
                                    <td>{{$value}}</td>
                                  </tr>
                                {{end}}
                                </tbody>
                              </table>
                            {{end}}
                            {{ if .PostParams }}
                              <h3>Post Parameters</h3>
                              <table class="table table-hover">
                                <thead>
                                <tr>
                                  <th scope="col">Key</th>
                                  <th scope="col">Value</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{ range $key, $value :=  .PostParams}}
                                  <tr>
                                    <th scope="row">{{$key}}</th>
                                    <td>{{$value}}</td>
                                  </tr>
                                {{end}}
                                </tbody>
                              </table>
                            {{end}}
                            {{ if .Raw }}
                              <h3>RawData</h3>
                              <table class="table table-hover">
                                <tbody>
                                  <tr>
                                    <td>{{.Raw}}</td>
                                  </tr>
                                </tbody>
                              </table>
                            {{end}}
                            {{ if .PostMultipart }}
                                {{if .PostMultipart.TmpFile}}
                                  <h3>Post Multipart Files</h3>
                                  <table class="table table-hover">
                                    <thead>
                                    <tr>
                                      <th scope="col">Key</th>
                                      <th scope="col">Value</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {{ range $key, $value :=  .PostMultipart.TmpFile}}
                                      <tr>
                                        <th scope="row">{{$key}}</th>
                                        <td>{{$value}}</td>
                                      </tr>
                                    {{end}}
                                    </tbody>
                                  </table>
                                {{end}}
                            {{ end }}
                        </div>
                        <div id="modal-{{$i}}-tab3" class="tab-pane fade">
                          <h3>Headers</h3>
                          <table class="table table-hover">
                            <thead>
                            <tr>
                              <th scope="col">Key</th>
                              <th scope="col">Value</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{ range $key, $value :=  .Headers}}
                              <tr>
                                <th scope="row">{{$key}}</th>
                                <td>{{$value}}</td>
                              </tr>
                            {{end}}
                            </tbody>
                          </table>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </td>
        </tr>
      {{ end }}
      </tbody>
    </table>
    <nav aria-label="...">
      <ul class="pagination mt-3 mb-0">
          {{ if .pagination.HasPrev }}
            <li class="page-item">
              <a class="page-link" href="{{.pagination.PrevPageUrl}}" tabindex="-1">Previous</a>
            </li>
          {{ end }}
          {{ if .pagination.HasNext }}
            <li class="page-item">
              <a class="page-link" href="{{.pagination.NextPageUrl}}">Next</a>
            </li>
          {{ end }}
      </ul>
    </nav>
  </div>
  </div>
</main>
<script src="//cdn.jsdelivr.net/npm/jquery@3.5.1/dist/jquery.slim.min.js" integrity="sha256-4+XzXVhsDmqanXGHaHvgh1gMQKX40OUvDEBTu8JcmNs=" crossorigin="anonymous"></script>
<script src="//cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js" integrity="sha256-/ijcOLwFf26xEYAjW75FizKVo5tnTYiQddPZoLUHHZ8=" crossorigin="anonymous"></script>
<script src="//cdn.jsdelivr.net/npm/bootstrap@4.5.0/dist/js/bootstrap.min.js" integrity="sha256-OFRAJNoaD8L3Br5lglV7VyLRf0itmoBzWUoM+Sji4/8=" crossorigin="anonymous"></script>
</html>`
)

var (
	allRequests = AllRequests{}
	pagination  = Pagination{}
)

func paginate(s []RequestStat, offset, length int) []RequestStat {
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}

	return s[offset:]
}

func inspector(r *znet.Engine, path string) func(c *znet.Context) {
	r.GET(path, func(c *znet.Context) {
		urlPath := c.Request.URL.Path
		page, _ := strconv.ParseFloat(c.DefaultQuery("page", "1"), 64)
		perPage, _ := strconv.ParseFloat(c.DefaultQuery("per_page", "20"), 64)
		total := float64(len(allRequests.Requets))
		totalPage := math.Ceil(total / perPage)
		offset := (page - 1) * perPage
		if offset < 0 {
			offset = 0
		}
		pagination.HasPrev = false
		pagination.HasNext = false
		pagination.CurrentPage = int(page)
		pagination.PerPage = int(perPage)
		pagination.TotalPage = int(totalPage)
		pagination.Total = int(total)
		pagination.Data = paginate(allRequests.Requets, int(offset), int(perPage))
		if pagination.CurrentPage > 1 {
			pagination.HasPrev = true
			pagination.PrevPageUrl = urlPath + "?page=" + strconv.Itoa(pagination.CurrentPage-1) + "&per_page=" + strconv.Itoa(pagination.PerPage)
		}
		if pagination.CurrentPage < pagination.TotalPage {
			pagination.HasNext = true
			pagination.NextPageUrl = urlPath + "?page=" + strconv.Itoa(pagination.CurrentPage+1) + "&per_page=" + strconv.Itoa(pagination.PerPage)
		}
		c.Template(http.StatusOK, tmp, map[string]interface{}{
			"title":      `Inspdeector`,
			"pagination": pagination,
			"path":       urlPath,
		}, map[string]interface{}{
			"formatDate": func(t time.Time) string {
				return ztime.FormatTime(t)
			},
		})
	})

	return func(c *znet.Context) {
		start := time.Now()
		_ = c.Request.ParseForm()
		_ = c.Request.ParseMultipartForm(10000)
		defer func() {
			if err := recover(); err != nil {
				fn := c.Engine.GetPanicHandler()
				if fn != nil {
					errMsg, ok := err.(error)
					if !ok {
						errMsg = errors.New(fmt.Sprint(err))
					}
					fn(c, errMsg)
				}
			}
			p := c.PrevContent()
			raw, _ := c.GetDataRaw()
			var content interface{}
			if p.Type == znet.ContentTypeJSON {
				content = template.HTML(p.Content)
			} else {
				content = zstring.Bytes2String(p.Content)
			}
			request := RequestStat{
				RequestedAt:    start,
				ProcessingTime: time.Since(start).String(),
				RequestUrl:     c.Request.URL.Path,
				HttpMethod:     c.Request.Method,
				HttpStatus:     p.Code,
				ContentType:    p.Type,
				Content:        content,
				Headers:        c.Request.Header,
				Cookies:        c.Request.Cookies(),
				GetParams:      c.Request.URL.Query(),
				PostParams:     c.Request.PostForm,
				Raw:            raw,
				PostMultipart:  c.Request.MultipartForm,
				ClientIP:       c.GetClientIP(),
			}
			allRequests.Requets = append([]RequestStat{request}, allRequests.Requets...)
		}()
		c.Next()
	}
}
