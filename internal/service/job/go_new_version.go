package job

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"golang.org/x/net/html"
	"wlb/internal/biz"
)

type GoNewVersion struct {
	larkClient  *lark.Client
	goVersionUC *biz.GoVersionUsecase
}

func NewGoNewVersion(
	client *lark.Client,
	goVersionUC *biz.GoVersionUsecase,
) *GoNewVersion {
	return &GoNewVersion{
		larkClient:  client,
		goVersionUC: goVersionUC,
	}
}

const (
	GoNewVersion_BIAPP_TOKEN = "Slf4bkZHtaeL3OsH9z9chwxNnzh"
	GoNewVersion_BITABLE_ID  = "tblfqQ8NLEMlRgGM"
)

type GoVersionData struct {
	Version string
	Date    time.Time
}

// Run 运行定时任务
func (job *GoNewVersion) Run() {
	log.Info("run go new version job")
	ctx := context.Background()

	url := "https://github.com/golang/go/tags"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET request failed:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// 解析 HTML
	doc, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("HTML parsing failed:", err)
		os.Exit(1)
	}
	var traverse func(*html.Node)
	var versionList []GoVersionData
	// 分析内容，找到 go 版本和发布时间
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h2" {
			if n.FirstChild != nil && n.FirstChild.FirstChild != nil && strings.HasPrefix(n.FirstChild.FirstChild.Data, "go") {
				findUl := n.Parent.Parent.NextSibling
				for {
					if findUl.Data != "ul" {
						findUl = findUl.NextSibling
					} else {
						break
					}
				}
				findLi := findUl.FirstChild
				for {
					if findLi.Data != "li" {
						findLi = findLi.NextSibling
					} else {
						break
					}
				}
				attr := findLi.LastChild.PrevSibling.Attr
				for _, a := range attr {
					if a.Key == "datetime" {
						// 解析时间字符串
						parsedTime, err := time.Parse(time.RFC3339, a.Val)
						if err != nil {
							fmt.Println("解析时间出错:", err)
							return
						}
						if len(versionList) < 6 {
							versionList = append(versionList, GoVersionData{
								Version: n.FirstChild.FirstChild.Data,
								Date:    parsedTime,
							})
						}
					}
				}

			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	fmt.Println(versionList)

	// 更新到多维表格
	if versionList != nil {
		for _, version := range versionList {
			job.createRecord(ctx, version)
		}
	}
}

func (job *GoNewVersion) createRecord(ctx context.Context, version GoVersionData) (*biz.GoVersion, error) {
	return job.goVersionUC.CreateGoVersion(ctx, &biz.GoVersion{
		Version:        version.Version,
		IsMajorVersion: isMajorVersion(version.Version),
		VersionDate:    version.Date,
	})
}

func isMajorVersion(version string) bool {
	if strings.Contains(version, "beta") || strings.Contains(version, "rc") {
		return false
	}
	return true
}
