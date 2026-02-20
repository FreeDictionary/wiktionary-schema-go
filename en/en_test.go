package en_test

import (
	"bufio"
	"bytes"
	"encoding/json/v2"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/FreeDictionary/wiktionary-schema-go/en"
)

const RAW_DATA_PATH string = "../test_data/raw-wiktextract-data.jsonl"

func TestCurrentDir(t *testing.T) {
	// print current directory
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Current directory:", wd)
}

func TestMarshalUnmarshalEnglish(t *testing.T) {
	file, err := os.Open(RAW_DATA_PATH)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close() // ⚠️ 记得关闭文件

	reader := bufio.NewReader(file)
	for {
		// ✅ 自动处理长行，返回完整一行（包含 \n）
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			t.Log("EOF")
			break
		} else if err != nil {
			t.Fatal(err)
		}

		// 去掉末尾的 \n 或 \r\n
		line = bytes.TrimRight(line, "\n")

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		var schema map[string]any
		if err := json.Unmarshal(line, &schema); err != nil {
			t.Fatal(err)
		}

		if lang_code, ok := schema["lang_code"].(string); !ok || lang_code != "en" {
			if lang, ok := schema["lang"].(string); !ok || strings.ToLower(lang) != "english" {
				continue
			}
		}

		// test marshal
		var enSchema en.WordData
		if err := json.Unmarshal(line, &enSchema); err != nil {
			t.Fatal(err)
		}

		// test unmarshal
		_, err = json.Marshal(enSchema)
		if err != nil {
			t.Fatal(err)
		}

	}
}
