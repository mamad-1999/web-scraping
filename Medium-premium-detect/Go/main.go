package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func isPremium(url string) (bool, error) {
	// Create a new allocator with custom user agent
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	// Create a new context with the allocator
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a new browser context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Set a timeout for the entire operation
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var isPremium bool

	// Run the browser tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"), // Wait for the body to load
		chromedp.Evaluate(`document.querySelectorAll('[aria-label="Close"]').forEach(btn => btn.click());`, nil), // Close popups
		chromedp.Evaluate(`{
			// Check for "Member-only story" text
			const xpathCheck = document.evaluate(
				'//*[contains(text(), "Member-only story")]',
				document,
				null,
				XPathResult.ANY_TYPE,
				null
			);
			const hasMemberText = xpathCheck.iterateNext() !== null;

			// Check for the golden star icon
			const hasGoldenStar = document.querySelector('svg[fill="#FFC017"]') !== null;

			// Return true if either condition is met
			hasMemberText || hasGoldenStar;
		}`, &isPremium),
	)

	if err != nil {
		return false, fmt.Errorf("error checking %s: %v", url, err)
	}

	return isPremium, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a URL")
	}

	url := os.Args[1]
	premium, err := isPremium(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if premium {
		fmt.Printf("🔒 %s is PREMIUM\n", url)
	} else {
		fmt.Printf("🔒 %s is NOT premium\n", url)
	}
}
