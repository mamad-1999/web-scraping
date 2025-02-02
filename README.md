# Medium Premium Checker

This is a Go-based tool that checks whether a given Medium article URL is a premium (member-only) story. It uses the `chromedp` package to automate a headless Chrome browser, navigate to the provided URL, and determine if the article is premium by checking for specific indicators such as "Member-only story" text or a golden star icon.

## Installation

1. **Install Go**: Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/medium-premium-checker.git
   cd medium-premium-checker
   ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build the Tool**:
   ```bash
   go build -o medium-premium-checker
   ```

## Usage

Run the tool from the command line by providing a Medium article URL as an argument:

```bash
./medium-premium-checker https://medium.com/your-article-url
```

### Example Output

- If the article is premium:
  ```bash
  🔒 https://medium.com/your-article-url is PREMIUM
  ```

- If the article is not premium:
  ```bash
  🔒 https://medium.com/your-article-url is NOT premium
  ```
