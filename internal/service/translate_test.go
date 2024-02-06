package service

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func Test_languageDetect(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name     string
		args     args
		wantLang language.Tag
	}{
		{"1", args{
			text: "这是一段文本，包含中文字符。れています。",
		}, language.English},
		{"1", args{
			text: "这This is a text with English characters.",
		}, language.Chinese},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLang := targetLanguageDetect(tt.args.text); !reflect.DeepEqual(gotLang, tt.wantLang) {
				t.Errorf("targetLanguageDetect() = %v, want %v", gotLang, tt.wantLang)
			}
		})
	}
}
