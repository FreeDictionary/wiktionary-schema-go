package main_test

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

const (
	RAW_WIKTEXTRACT_DATA_GZ string = "https://kaikki.org/dictionary/raw-wiktextract-data.jsonl.gz"
	SAVE_PATH               string = "./test_data/raw-wiktextract-data.jsonl.gz"
	RAW_DATA_PATH           string = "./test_data/raw-wiktextract-data.jsonl"
)

var FORCE_DOWNLOAD bool = false

func init() {
	// Read `FORCE_DOWNLOAD` from environment variable
	FORCE_DOWNLOAD = (strings.ToLower(os.Getenv("FORCE_DOWNLOAD")) == "true")
	log.Println("FORCE_DOWNLOAD:", FORCE_DOWNLOAD)
}

// Download `RAW_WIKTIONARY_DATA` to `SAVE_PATH`
func downloadRawWiktionaryData() error {
	resp, err := http.Get(RAW_WIKTEXTRACT_DATA_GZ)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s", resp.Status)
	}

	out, err := os.Create(SAVE_PATH)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// decompress the .gz file to `RAW_DATA_PATH`
func decompressRawWiktionaryData() error {
	in, err := os.Open(SAVE_PATH)
	if err != nil {
		return err
	}
	defer in.Close()

	gzReader, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	out, err := os.Create(RAW_DATA_PATH)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, gzReader)
	return err
}

func TestDownloadAndDecompressRawWiktionaryData(t *testing.T) {
	err := os.MkdirAll("./test_data", 0755)
	if err != nil {
		t.Fatalf("failed to create test_data directory: %v", err)
	}

	// Step 1: Check if the decompressed file already exists
	// If it exists and FORCE_DOWNLOAD is false, skip everything
	if _, err := os.Stat(RAW_DATA_PATH); err == nil && !FORCE_DOWNLOAD {
		t.Logf("file %s already exists, skipping download and decompress", RAW_DATA_PATH)
		return
	}

	// Step 2: Ensure the compressed file exists
	// If it doesn't exist, or FORCE_DOWNLOAD is true, download it
	if FORCE_DOWNLOAD {
		t.Logf("FORCE_DOWNLOAD is true, re-downloading...")
	}
	if _, err := os.Stat(SAVE_PATH); os.IsNotExist(err) || FORCE_DOWNLOAD {
		err = downloadRawWiktionaryData()
		if err != nil {
			t.Fatalf("failed to download: %v", err)
		}
		t.Logf("downloaded %s", SAVE_PATH)
	} else {
		t.Logf("file %s already exists, skipping download", SAVE_PATH)
	}

	// Step 3: Decompress the file (always do this since we need fresh decompressed file)
	err = decompressRawWiktionaryData()
	if err != nil {
		t.Fatalf("failed to decompress: %v", err)
	}
	t.Logf("decompressed to %s", RAW_DATA_PATH)

	// Step 4: Verify the decompressed file was created successfully
	if _, err := os.Stat(RAW_DATA_PATH); os.IsNotExist(err) {
		t.Fatalf("decompressed file %s was not created", RAW_DATA_PATH)
	}
	t.Logf("successfully ensured %s exists", RAW_DATA_PATH)
}
