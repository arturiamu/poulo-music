package usecase

import (
	ua "github.com/wux1an/fake-useragent"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Http struct {
	method    string
	url       string
	header    map[string]string
	body      string
	parameter url.Values
}

func NewHttp() *Http {
	return &Http{
		method:    "GET",
		header:    map[string]string{"referer": "https://www.bilibili.com", "User-Agent": ua.Random()},
		parameter: url.Values{},
	}
}

func (h *Http) SetMethod(method string) *Http {
	h.method = method
	return h
}

func (h *Http) SetHeader(key, value string) *Http {
	h.header[key] = value
	return h
}

func (h *Http) SetHeaders(headers map[string]string) *Http {
	for k, v := range headers {
		h.header[k] = v
	}
	return h
}

func (h *Http) SetUrl(url string) *Http {
	h.url = url
	return h
}

func (h *Http) AddParam(k, v string) *Http {
	h.parameter.Add(k, v)
	return h
}

func (h *Http) AddParams(params map[string]string) *Http {
	for k, v := range params {
		h.AddParam(k, v)
	}
	return h
}

func (h *Http) DoDetail() (*http.Response, error) {
	var param = h.parameter.Encode()

	if strings.Contains(h.url, "?") {
		h.url += "&" + param
	} else {
		h.url += "?" + param
	}

	req, err := http.NewRequest(h.method, h.url, strings.NewReader(h.body))
	if err != nil {
		return nil, err
	}

	for k, v := range h.header {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}

func (h *Http) Do() ([]byte, error) {
	resp, err := h.DoDetail()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
