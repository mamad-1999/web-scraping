const puppeteer = require('puppeteer');

async function isPremium(url) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    try {

        await page.setUserAgent('Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36');
        await page.setExtraHTTPHeaders({
            'Accept-Language': 'en-US,en;q=0.9'
        });
        // Bypass cookie dialogs and overlays
        await page.goto(url, {
            waitUntil: 'networkidle2',
            timeout: 60000
        });

        // Close any popups/modals
        await page.evaluate(() => {
            document.querySelectorAll('[aria-label="Close"]').forEach(btn => btn.click());
        });

        // Add inside the try block after page.goto()
        await page.screenshot({ path: 'debug.png' });

        // Wait for article content to load
        await page.waitForSelector('article', { timeout: 5000 });

        // Check using XPath for precise text matching
        const isPremium = await page.evaluate(() => {
            const xpathCheck = document.evaluate(
                '//*[contains(text(), "Member-only story")]',
                document,
                null,
                XPathResult.ANY_TYPE,
                null
            );

            return xpathCheck.iterateNext() !== null ||
                document.querySelector('svg[fill="#FFC017"]') !== null;
        });

        await browser.close();
        return isPremium;
    } catch (error) {
        console.error(`Error checking ${url}:`, error);
        await browser.close();
        return false;
    }
}

// Usage
const url = process.argv[2];
if (!url) {
    console.log('Please provide a URL');
    process.exit(1);
}

isPremium(url)
    .then(premium => {
        console.log(`🔒 ${url} is ${premium ? 'PREMIUM' : 'NOT premium'}`);
        process.exit(0);
    })
    .catch(err => {
        console.error(err);
        process.exit(1);
    });