package discord

import (
	"fmt"
	"regexp"

	"github.com/valyala/fasthttp"
)

var (
	JS_FILE_REGEX    = regexp.MustCompile(`src="(/assets/[^"]+\.js)"`)
	BUILD_INFO_REGEX = regexp.MustCompile(`build_number:"(\d+)"`)
)

func getLatestBuild() (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI("https://discord.com/app")

	if err := requestClient.Do(req, resp); err != nil {
		return "", err
	}

	matches := JS_FILE_REGEX.FindAllStringSubmatch(string(resp.Body()), -1)
	if len(matches) == 0 {
		fmt.Println("build number not found, falling back to 9999")
		return "9999", nil
	}

	for _, match := range matches {
		if len(match) < 2 || match[1] == "" {
			continue
		}
		asset := match[1]

		jsReq := fasthttp.AcquireRequest()
		jsResp := fasthttp.AcquireResponse()
		jsReq.Header.SetMethod(fasthttp.MethodGet)
		jsReq.SetRequestURI(fmt.Sprintf("https://discord.com%s", asset))
		if err := requestClient.Do(jsReq, jsResp); err != nil {
			fasthttp.ReleaseRequest(jsReq)
			fasthttp.ReleaseResponse(jsResp)
			continue
		}
		m := BUILD_INFO_REGEX.FindStringSubmatch(string(jsResp.Body()))
		fasthttp.ReleaseRequest(jsReq)
		fasthttp.ReleaseResponse(jsResp)

		if len(m) >= 2 {
			return m[1], nil
		}
	}

	fmt.Println("build number not found, falling back to 9999")
	return "9999", nil
}

func mustGetLatestBuild() string {
	if build, err := getLatestBuild(); err != nil {
		panic(err)
	} else {
		return build
	}
}
