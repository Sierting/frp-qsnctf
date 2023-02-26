// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io"
	"net/http"
	"os"

	frpLog "github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/version"
)

var NotFoundPagePath = ""

const (
	NotFound = `<!DOCTYPE html>
<html>
<head>
  <title>青少年CTF-GameBox提醒</title>
  <meta charset="utf-8">
  <meta http-equiv="refresh" content="30">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://cdn.bootcss.com/twitter-bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
  <script src="https://cdn.bootcss.com/popper.js/1.15.0/umd/popper.min.js"></script>
  <script src="https://cdn.bootcss.com/twitter-bootstrap/4.3.1/js/bootstrap.min.js"></script>
</head>
<body  style="background-image:url(https://tools.qsnctf.com/tools/api/qsn_wallpaper.php); background-size: cover;">
<div class="container-fluid">
  <div class="jumbotron" style="margin-top: 150px; background-color: rgba(255, 255, 255, 0.5);">
  <h2 >错误信息</h2>
  <div class="alert alert-info alert-dismissible" style="margin-top: 20px;">
    <button type="button" class="close" data-dismiss="alert">&times;</button>
    <strong>错误：</strong> 你访问的Game-Box未找到，可能原因有下：
  </div>
    <div class="alert alert-primary">
       Game-Box<strong>初始化中</strong>，请稍等片刻即可访问。
    </div>
    <div class="alert alert-warning">
       Game-Box<strong>已到期</strong>，请重新启动环境即可。
    </div>
    <div class="alert alert-danger">
       Game-Box<strong>已销毁</strong>，请重新开启题目环境。
    </div>
    <h2>操作</h2>
  <a href="https://www.qsnctf.com/" class="btn btn-info" role="button">青少年CTF训练平台</a>
  <a href="https://docs.qsnctf.com/" class="btn btn-info" role="button">青少年CTF文库</a>
  <a href="https://bbs.qsnctf.com/" class="btn btn-info" role="button">青少年CTF论坛</a>
  <a href="https://tools.qsnctf.com/" class="btn btn-info" role="button">青少年CTF在线工具</a>
  <a href="https://jq.qq.com/?_wv=1027&k=yRnKW3uX" class="btn btn-info" role="button">QQ群</a>
  <a href="https://space.bilibili.com/2066710972" class="btn btn-info" role="button">Bilibili</a>
</div>
</div>
<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?10309f8528ef7f3bdd779aa12ad6dc7e";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>
</html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = os.ReadFile(NotFoundPagePath)
		if err != nil {
			frpLog.Warn("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	content := getNotFoundPageContent()
	res := &http.Response{
		Status:        "Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	return res
}

func noAuthResponse() *http.Response {
	header := make(map[string][]string)
	header["WWW-Authenticate"] = []string{`Basic realm="Restricted"`}
	res := &http.Response{
		Status:     "401 Not authorized",
		StatusCode: 401,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	return res
}
