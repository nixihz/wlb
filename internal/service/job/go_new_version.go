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
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"golang.org/x/net/html"
)

type GoNewVersionJob struct {
	larkClient *lark.Client
}

func NewGoNewVersionJob(
	client *lark.Client,
) *GoNewVersionJob {
	return &GoNewVersionJob{
		larkClient: client,
	}
}

const (
	GoNewVersion_BIAPP_TOKEN = "Slf4bkZHtaeL3OsH9z9chwxNnzh"
	GoNewVersion_BITABLE_ID  = "tblfqQ8NLEMlRgGM"
)

type GoVersion struct {
	Version string
	Date    time.Time
}

// Run 运行定时任务
func (job *GoNewVersionJob) Run() {
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
	var versionList []GoVersion
	now := time.Now()
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
						timeGap := now.Add(time.Hour * 24 * -1)
						if timeGap.Unix() < parsedTime.Unix() {
							versionList = append(versionList, GoVersion{
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

func (job *GoNewVersionJob) createRecord(ctx context.Context, version GoVersion) (string, error) {
	boolMap := map[bool]string{true: "是", false: "否"}
	appTableRecord := larkbitable.NewAppTableRecordBuilder().
		Fields(map[string]interface{}{
			"版本号":       version.Version,
			"是否Major版本": boolMap[isMajorVersion(version.Version)],
			"发布日期":      version.Date.UnixMilli(),
		}).Build()

	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(GoNewVersion_BIAPP_TOKEN).TableId(GoNewVersion_BITABLE_ID).
		AppTableRecord(appTableRecord).
		Build()
	resp, err := job.larkClient.Bitable.AppTableRecord.Create(ctx, req)
	if err != nil || resp.Data == nil {
		log.Error(err)
		return "", err
	}
	return *resp.Data.Record.RecordId, nil
}

func isMajorVersion(version string) bool {
	if strings.Contains(version, "beta") || strings.Contains(version, "rc") {
		return false
	}
	return true
}
