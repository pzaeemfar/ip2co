package geoip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"

	ip2location "github.com/ip2location/ip2location-go/v9"
)

const dbURL = "https://github.com/pzaeemfar/oip2co/raw/refs/heads/main/database/database.BIN"
const dbFileName = "database-1704f38bf0b916536afc7712c14da229.BIN"

func getDatabasePath() (string, error) {
	tmpDir := os.TempDir()
	dbPath := filepath.Join(tmpDir, dbFileName)

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err := downloadFile(dbURL, dbPath); err != nil {
			return "", err
		}
	}

	return dbPath, nil
}

func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download database: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func GetCountry(ipStr string, debug bool) (string, error) {
	dbPath, err := getDatabasePath()
	if err != nil {
		return "", err
	}

	db, err := ip2location.OpenDB(dbPath)
	if err != nil {
		return "", fmt.Errorf("failed to open IP2Location DB: %w", err)
	}
	defer db.Close()

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", fmt.Errorf("invalid IP: %s", ipStr)
	}

	results, err := db.Get_all(ip.String())
	if err != nil {
		return "", fmt.Errorf("lookup error: %w", err)
	}

	code := results.Country_short
	if code == "-" {
		code = "Unknown"
	}

	if debug {
		fmt.Printf("IP: %s → Country: %s (%s)\n", ipStr, results.Country_long, code)
	}

	return code, nil
}
