package tools

import (
	"fmt"
	"net/http"
)

func Resend(r *http.Request, to string) (resp *http.Response, err error) {
	req, err := http.NewRequest(r.Method, fmt.Sprintf("http://%s%s", to, r.RequestURI), r.Body)
	if err != nil {
		return resp, err
	}
	req.Header = r.Header.Clone()
	cli := &http.Client{}

	return cli.Do(req)
}
