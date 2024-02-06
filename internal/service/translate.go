package service

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	googletrans "cloud.google.com/go/translate"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"wlb/api/translate"
)

type TranslateService struct {
	APIKey string
}

func NewTranslateService() *TranslateService {
	apikey := ""
	for _, envVar := range os.Environ() {
		v := strings.Split(envVar, "=")
		if len(v) > 1 && v[0] == "GOOGLE_API_KEY" {
			apikey = v[1]
			break
		}
	}

	return &TranslateService{
		APIKey: apikey,
	}
}

func (t *TranslateService) Text(ctx context.Context, request *translate.TranslateTextRequest) (*translate.TranslateTextReply, error) {
	output, err := t.trans(ctx, request.Input)
	if err != nil {
		log.Error(err)
	}

	return &translate.TranslateTextReply{
		Output: output,
	}, nil
}

func (t *TranslateService) trans(ctx context.Context, sourceText string) (string, error) {
	lang := targetLanguageDetect(sourceText)

	client, err := googletrans.NewClient(ctx, option.WithAPIKey(t.APIKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{sourceText}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %w", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", sourceText)
	}
	return resp[0].Text, nil
}

func targetLanguageDetect(text string) (lang language.Tag) {
	// 创建一个正则表达式模式，用于匹配中文字符
	pattern := regexp.MustCompile(`\p{Han}`)

	// 计算文字中的中文字符数量
	matches := pattern.FindAllString(text, -1)
	chineseCharCount := len(matches)

	// 计算文字中非中文字符数量
	nonChineseCharCount := 0
	for _, char := range text {
		if !unicode.Is(unicode.Scripts["Han"], char) {
			nonChineseCharCount++
		}
	}

	// 判断大部分是否为中文
	if chineseCharCount > nonChineseCharCount {
		lang = language.English
	} else {
		lang = language.Chinese
	}

	return lang
}
