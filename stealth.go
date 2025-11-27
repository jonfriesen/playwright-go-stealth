package stealth

import (
	_ "embed"

	"github.com/playwright-community/playwright-go"
)

//go:embed "stealth.min.js"
var StealthJS string

//go:embed "chrome_stealth.js"
var ChromeStealthJS string

// Options configures which stealth scripts to inject.
type Options struct {
	// ChromeStealth enables additional Chrome-specific evasions including:
	// - Chrome app/runtime object patching
	// - Console method silencing
	// - Permissions query modification (changes "prompt" to "denied")
	// - Event listener isTrusted spoofing
	// - Automatic Cloudflare challenge checkbox clicking
	//
	// This is useful for bypassing additional bot detection mechanisms
	// that the base stealth script doesn't cover.
	// See: https://github.com/jonfriesen/playwright-go-stealth/issues/2
	ChromeStealth bool
}

// DefaultOptions returns the default stealth options with all features disabled.
// The base stealth script is always injected by InjectWithOptions.
func DefaultOptions() Options {
	return Options{
		ChromeStealth: false,
	}
}

// Inject adds a stealth JavaScript script to the given Playwright page to make it less detectable by automated detection mechanisms.
//
// The stealth script modifies various browser characteristics to evade detection by websites that block or track bots. The script is embedded in the Go binary using the //go:embed directive, which includes the contents of the "stealth.min.js" file as the StealthJS variable. Get the latest version of playwright-go-stealth for more up to version os the evasions.
//
// Usage example:
//
//	page, err := browser.NewPage()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	err = stealth.Inject(page)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// The Inject function uses the AddInitScript method to ensure the stealth script is executed before any other scripts on the page. This method is effective in bypassing many bot detection systems that run their detection logic early in the page lifecycle.
func Inject(page playwright.Page) error {
	return page.AddInitScript(playwright.Script{
		Content: playwright.String(StealthJS),
	})
}

// InjectWithOptions adds stealth JavaScript scripts to the given Playwright page
// with configurable options for additional evasion techniques.
//
// The base stealth script (stealth.min.js) is always injected. Additional scripts
// can be enabled via the Options struct.
//
// Usage example:
//
//	page, err := browser.NewPage()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	opts := stealth.Options{
//	    ChromeStealth: true, // Enable additional Chrome evasions
//	}
//	err = stealth.InjectWithOptions(page, opts)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// The scripts are injected using AddInitScript to ensure they execute before
// any other scripts on the page.
func InjectWithOptions(page playwright.Page, opts Options) error {
	// Always inject the base stealth script
	if err := page.AddInitScript(playwright.Script{
		Content: playwright.String(StealthJS),
	}); err != nil {
		return err
	}

	// Inject optional Chrome stealth script
	if opts.ChromeStealth {
		if err := page.AddInitScript(playwright.Script{
			Content: playwright.String(ChromeStealthJS),
		}); err != nil {
			return err
		}
	}

	return nil
}
