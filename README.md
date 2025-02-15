# TryHackMe Badge Screenshot Automation

This project automates the process of capturing a high-quality screenshot of a TryHackMe profile badge and updating it in a GitHub repository every 12 hours using GitHub Actions.

## Preview

![2025-02-02_16-15](https://github.com/user-attachments/assets/f295bff7-0a34-41f2-b016-cd732a5dd113)


## Features
- Uses Puppeteer to capture a high-resolution screenshot of the TryHackMe badge.
- Automates the update process with a GitHub Action.
- Runs every 12 hours via a scheduled cron job.

## Prerequisites
Ensure you have the following installed:
- [Node.js](https://nodejs.org/) (Version 18 recommended)
- [Puppeteer](https://pptr.dev/)
- Dependencies required by Puppeteer:
  ```sh
  sudo apt-get install -y libnss3 libxss1 fonts-liberation libatk1.0-0 \
  libatk-bridge2.0-0 libcups2 libdrm2 libgbm1 libnspr4 libxcomposite1 \
  libxdamage1 libxrandr2 libxtst6 xdg-utils
  ```

## Installation & Setup

1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/yourrepository.git
   cd yourrepository
   ```

2. Install Node.js dependencies:
   ```sh
   npm install puppeteer
   ```

3. Create a new file named `generateBadge.js` and paste the following script:
   ```javascript
   const puppeteer = require('puppeteer');

   (async () => {
       try {
           const browser = await puppeteer.launch({
               args: ['--no-sandbox', '--disable-setuid-sandbox'],
           });
           const page = await browser.newPage();

           await page.setViewport({
               width: 1920,
               height: 1080,
               deviceScaleFactor: 2,
           });

           const url = 'https://tryhackme.com/api/v2/badges/public-profile?userPublicId=3153096';
           await page.goto(url, { waitUntil: 'networkidle2' });

           const badgeSelector = '#thm-badge';
           await page.waitForSelector(badgeSelector);

           await page.addStyleTag({
               content: '#thm-badge { transform: scale(1.5); transform-origin: top left; }',
           });

           const badgeElement = await page.$(badgeSelector);
           if (badgeElement) {
               await badgeElement.screenshot({ path: 'tryhackme-badge-high-quality.png' });
               console.log('High-quality screenshot saved as tryhackme-badge-high-quality.png');
           } else {
               console.log('Badge element not found!');
           }

           await browser.close();
       } catch (error) {
           console.error('Error generating screenshot:', error);
       }
   })();
   ```

4. Run the script manually to test it:
   ```sh
   node generateBadge.js
   ```
   This should generate a `tryhackme-badge-high-quality.png` file.

5. Move the generated image to the `assets/` folder:
   ```sh
   mv tryhackme-badge-high-quality.png assets/
   ```

## Automating with GitHub Actions

To automate the badge update process, create a GitHub Action workflow:

1. Inside your repository, navigate to `.github/workflows/` (create these folders if they don't exist).
2. Create a new file named `update_badge.yml` and paste the following content:

   ```yaml
   name: Update TryHackMe Badge

   on:
     schedule:
       - cron: "0 */12 * * *"  # Runs every 12 hours

   jobs:
     update-badge:
       runs-on: ubuntu-latest

       steps:
         - name: Checkout repository
           uses: actions/checkout@v3

         - name: Set up Node.js
           uses: actions/setup-node@v3
           with:
             node-version: "18"

         - name: Install Puppeteer dependencies
           run: sudo apt-get install -y libnss3 libxss1 fonts-liberation libatk1.0-0 libatk-bridge2.0-0 libcups2 libdrm2 libgbm1 libnspr4 libxcomposite1 libxdamage1 libxrandr2 libxtst6 xdg-utils

         - name: Install Puppeteer
           run: npm install puppeteer

         - name: Generate TryHackMe badge
           run: node generateBadge.js

         - name: Move badge to assets folder
           run: mv tryhackme-badge-high-quality.png assets/tryhackme-badge-high-quality.png

         - name: Commit and push changes
           run: |
             git config --global user.name "github-actions[bot]"
             git config --global user.email "github-actions[bot]@users.noreply.github.com"
             git add assets/tryhackme-badge-high-quality.png
             git commit -m "Update TryHackMe badge"
             git push
   ```

## Final Steps
1. Commit and push your changes to GitHub:
   ```sh
   git add .
   git commit -m "Added TryHackMe badge automation"
   git push origin main
   ```
2. The GitHub Action will now run automatically every 12 hours to update your badge.
----

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

# TryHackMe Streak Notifier

This project monitors your TryHackMe streak and sends an alert to a Telegram channel if the streak does not increase.

## Requirements

Install dependencies using:
```bash
pip install -r requirements.txt
```

## Setup

1. Create a `.env` file and add your bot token and channel ID:
   ```ini
   BOT_TOKEN=your_bot_token
   CHANNEL_ID=your_channel_id
   ```

2. Run the script:
   ```bash
   python script.py
   ```

## How It Works
- Fetches your current TryHackMe streak from the API.
- Compares it with the previously stored streak in `streak.txt`.
- If the streak has not increased, it sends an alert to the configured Telegram channel.

## Notes
- Ensure `streak.txt` exists or the script will create it.
- The script should be run periodically (e.g., via a cron job) to check streak updates.

Github Action usage of this script:
```yml
name: Check THM Streak

on:
  schedule:
    - cron: "30 16 * * *"
  workflow_dispatch:

jobs:
  check-streak:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4
        with:
          fetch-depth: 0 # Required to allow committing changes

      - name: Set up Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.x"

      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install requests beautifulsoup4 python-dotenv

      - name: Create .env file
        run: |
          echo "BOT_TOKEN=${{ secrets.BOT_TOKEN }}" > .env
          echo "CHANNEL_ID=${{ secrets.CHANNEL_ID }}" >> .env

      - name: Run streak check
        run: python script.py

      - name: Commit streak changes
        id: commit
        run: |
          git config --global user.name "github-actions[bot]"
          git config --global user.email "github-actions[bot]@users.noreply.github.com"
          git add streak.txt
          git diff --quiet && git diff --staged --quiet || (git commit -m "Update streak.txt [skip ci]" && git push)

```
