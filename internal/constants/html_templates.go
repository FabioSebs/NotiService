package constants

const (
	HTML_NEW_SCRAPE = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ICCT RSS Feed Update</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333333;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            padding-bottom: 20px;
        }
        .header h1 {
            margin: 0;
            color: #4CAF50;
        }
        .content {
            text-align: left;
        }
        .content p {
            margin: 10px 0;
        }
        .footer {
            text-align: center;
            margin-top: 20px;
            font-size: 12px;
            color: #777777;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>ICCT</h1>
        </div>
        <div class="content">
            <p>Dear User,</p>
            <p>We are excited to inform you that there may be new entries in our RSS feed. Our scraper has just finished its latest run and updated the feed with the latest content.</p>
            <p>Please check our RSS feed to stay updated with the latest entries and information.</p>
            <p>Best regards,</p>
            <p>The ICCT Team</p>
        </div>
        <div class="footer">
            <p>ICCT &copy; 2024. All rights reserved.</p>
        </div>
    </div>
</body>
</html>

	`

	HTML_NEW_ENTRY = `
	
	`
)
