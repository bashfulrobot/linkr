package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	URL       string `yaml:"url"`
	RewriteTo string `yaml:"rewrite_to,omitempty"`
	Browser   string `yaml:"browser"`
	App       bool   `yaml:"app,omitempty"` // Optional flag to indicate whether to use --app
}

type Config struct {
	DefaultBrowser string `yaml:"default_browser"`
	Rules          []Rule `yaml:"rules"`
}

var configFilePaths = map[string]string{
	"linux":  os.Getenv("HOME") + "/.config/linkr/config.yaml",
	"darwin": os.Getenv("HOME") + "/Library/Application Support/linkr/config.yaml",
}

func getConfigFilePath() string {
	platform := runtime.GOOS
	path, exists := configFilePaths[platform]
	if !exists {
		panic("Unsupported platform: " + platform)
	}
	return path
}

func loadConfig() (Config, error) {
	configPath := getConfigFilePath()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func normalizeURL(url string) string {
	// Remove trailing slash if present
	url = strings.TrimSuffix(url, "/")

	// Check if the URL starts with "http://" or "https://"
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	// If not, add "https://" prefix
	return "https://" + url
}

func rewriteURL(requestedURL string, config Config) string {
	requestedURL = normalizeURL(requestedURL)

	// Iterate through each rule in the configuration
	for _, rule := range config.Rules {
		// Check if the requested URL matches the rule's URL
		if strings.Contains(requestedURL, rule.URL) {
			if rule.RewriteTo != "" {
				fmt.Printf("Rewriting %s to %s using %s\n", requestedURL, rule.RewriteTo, rule.Browser)

				// Construct the browser command with the original URL
				if rule.App {
					return fmt.Sprintf("%s --new-window --app=%s", rule.Browser, rule.RewriteTo)
				} else {
					return fmt.Sprintf("%s %s", rule.Browser, rule.RewriteTo)
				}

			} else {
				fmt.Printf("Launching %s to %s\n", requestedURL, rule.Browser)

				// Construct the browser command with the original URL
				if rule.App {
					return fmt.Sprintf("%s --new-window --app=%s", rule.Browser, requestedURL)
				} else if !rule.App {
					return fmt.Sprintf("%s %s", rule.Browser, requestedURL)
				}
			}
		}
	}

	// If no match is found, use the default browser without modifications
	fmt.Printf("Launching %s to the default browser\n", requestedURL)
	if config.DefaultBrowser == "" {
		platform := runtime.GOOS
		configPath, exists := configFilePaths[platform]
		if exists {
			fmt.Printf("Error: DefaultBrowser not set in the configuration YAML.\n")
			fmt.Printf("Please set the DefaultBrowser in the configuration YAML.\n")
			fmt.Printf("Configuration file path for %s: %s\n", platform, configPath)
		} else {
			fmt.Printf("Unsupported platform: %s\n", platform)
		}
		os.Exit(1)
	}

	if config.Rules[len(config.Rules)-1].App {
		return fmt.Sprintf("%s --new-window --app=%s", config.DefaultBrowser, requestedURL)
	} else {
		return fmt.Sprintf("%s %s", config.DefaultBrowser, requestedURL)
	}
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: linkr <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	rewrittenURL := rewriteURL(url, config)

	// Split the browser command and arguments
	parts := strings.Fields(rewrittenURL)

	browserCmd := parts[0]
	browserArgs := parts[1:]

	// Execute the browser command with arguments
	cmd := exec.Command(browserCmd, browserArgs...)

	fmt.Printf("Launching browser with URL: %s\n", cmd)

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error launching browser:", err)
		os.Exit(1)
	}
}
