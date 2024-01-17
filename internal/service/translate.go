package service

import (
	"context"
	"fmt"
	"os"
	"strings"

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
	lang, err := language.Parse("zh-CN")
	if err != nil {
		return "", fmt.Errorf("language.Parse: %w", err)
	}

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
