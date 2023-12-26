package services_utils

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// URLCheckError represents custom error types for URL validation.
type URLCheckError int

// URLCheckConfig holds the configuration for URL checking.
type URLCheckConfig struct {
	MaxRedirects      int
	MaxSize           int
	MaxURLLength      int
	CheckForFile      bool
	CheckReachable    bool
	HTTPClientTimeout time.Duration
}

const (
	ErrInvalidSchema URLCheckError = iota
	ErrFileNotAllowed
	ErrMaxRedirectsExceeded
	ErrSizeExceeded
	ErrURLLengthExceeded
	ErrLoginKeywordsFound
	ErrXSSDetected
	ErrUnreachable
	ErrMustHaveTLD
	ErrURLCannotContainCreds
	ErrHTTPClientTimeout
)

var urlCheckErrorMessages = map[URLCheckError]string{
	ErrInvalidSchema:         "Invalid URL schema. Only 'http' or 'https' are allowed.",
	ErrFileNotAllowed:        "Files are not allowed.",
	ErrMaxRedirectsExceeded:  "Maximum number of redirects exceeded.",
	ErrSizeExceeded:          "URL size exceeds the maximum allowed size.",
	ErrURLLengthExceeded:     "URL length exceeds the maximum allowed length.",
	ErrUnreachable:           "URL is unreachable.",
	ErrMustHaveTLD:           "URL must have a valid top-level domain (TLD).",
	ErrURLCannotContainCreds: "URL should not contain credentials.",
	ErrHTTPClientTimeout:     "Website is taking too long to respond",
}

var DefaultURLCheckConfig = URLCheckConfig{
	MaxRedirects:      2,       // Set your desired default values
	MaxSize:           4194304, // 4MB
	MaxURLLength:      2048,    // 2048 characters
	CheckForFile:      true,
	CheckReachable:    true,
	HTTPClientTimeout: 10 * time.Second,
}

// CheckURL validates a URL based on the default configuration and returns custom errors.
func ValidateURL(inputURL string) error {
	return ValidateURLWithConfig(inputURL, DefaultURLCheckConfig)
}

// CheckURL validates a URL based on the given configuration and returns custom errors.
func ValidateURLWithConfig(inputURL string, config URLCheckConfig) error {
	// Validate URL
	validationStartTime := time.Now()
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return err
	}

	if !parsedURL.IsAbs() {
		return errors.New(urlCheckErrorMessages[ErrInvalidSchema])
	}

	// Check schema (http or https)
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New(urlCheckErrorMessages[ErrInvalidSchema])
	}

	// Check for TLD (top-level domain)
	if !HasValidTLD(parsedURL.Host) {
		return errors.New(urlCheckErrorMessages[ErrMustHaveTLD])
	}

	// Check for username and password in the URL http://username:password@yoursitename.com
	if parsedURL.User != nil {
		return errors.New(urlCheckErrorMessages[ErrURLCannotContainCreds])
	}

	// Check for file
	if !config.CheckForFile && strings.HasPrefix(parsedURL.Path, "/") {
		return errors.New(urlCheckErrorMessages[ErrFileNotAllowed])
	}

	// Check URL length
	if config.MaxURLLength > 0 {
		if len(inputURL) > config.MaxURLLength {
			return errors.New(urlCheckErrorMessages[ErrURLLengthExceeded])
		}
	}

	log.Printf("ValidateURL took: %v ", time.Since(validationStartTime))

	// Check URL reachability
	if config.CheckReachable {
		reachabilityStartTime := time.Now()
		_err := CheckIsReachableURLWithConfig(inputURL, config)
		log.Printf("CheckReachableURL took: %v", time.Since(reachabilityStartTime))
		if _err != nil {
			return _err
		}
	}

	return nil
}

// ? This function when used directly panics when tries to return error.
func CheckIsReachableURL(inputURL string) error {
	return CheckIsReachableURLWithConfig(inputURL, DefaultURLCheckConfig)
}

func CheckIsReachableURLWithConfig(inputURL string, config URLCheckConfig) error {
	// Check URL for size, length, and reachability
	if config.MaxSize > 0 || config.MaxURLLength > 0 || config.CheckReachable {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= config.MaxRedirects {
					return errors.New(urlCheckErrorMessages[ErrMaxRedirectsExceeded])
				}

				// Todo Save somehow redirects
				if len(via) > 0 {
					log.Printf("Redirected to: %s", req.URL)
				}
				return nil
			},
			Timeout: config.HTTPClientTimeout,
		}

		resp, err := client.Head(inputURL)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				// Handle timeout error
				return errors.New(urlCheckErrorMessages[ErrHTTPClientTimeout])
			}
			return errors.New(urlCheckErrorMessages[ErrUnreachable])
		}

		// Check Content Size
		if config.MaxSize > 0 {
			contentLength := resp.Header.Get("Content-Length")
			if contentLength != "" {
				size, err := strconv.Atoi(contentLength)
				if err != nil {
					return err
				}
				if size > config.MaxSize {
					return errors.New(urlCheckErrorMessages[ErrSizeExceeded])
				}
			}
		}
	}

	return nil
}

func HasValidTLD(host string) bool {
	// Define a regular expression to match common TLD patterns.
	tldPattern := `(\.[a-zA-Z]{2,63})$`
	matched, err := regexp.MatchString(tldPattern, host)
	return matched && err == nil
}
