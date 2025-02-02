# Medium Premium Checker (Golang)

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

# Medium Premium Checker (Node.js Puppeteer)

This is a Node.js-based tool that checks whether a given Medium article URL is a premium (member-only) story. It uses the `puppeteer` package to automate a headless Chrome browser, navigate to the provided URL, and determine if the article is premium by checking for specific indicators such as "Member-only story" text or a golden star icon.

## Installation

1. **Install Node.js**: Ensure you have Node.js installed on your system. You can download it from [here](https://nodejs.org/).

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/medium-premium-checker-puppeteer.git
   cd medium-premium-checker-puppeteer
   ```

3. **Install Dependencies**:
   ```bash
   npm install puppeteer
   ```

## Usage

Run the tool from the command line by providing a Medium article URL as an argument:

```bash
node index.js https://medium.com/your-article-url
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
