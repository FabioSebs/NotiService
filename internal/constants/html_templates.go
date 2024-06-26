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

	HTML_NEW_EMAIL = `
    <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Fabrzy!</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            padding: 20px 0;
            background-color: #000;
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
        }
        .header img {
            max-width: 150px;
        }
        .content {
            text-align: center;
        }
        .content h1 {
            color: #D32F2F;
        }
        .content p {
            font-size: 16px;
            line-height: 1.5;
        }
        .footer {
            text-align: center;
            padding: 20px 0;
            color: #777;
            font-size: 12px;
            background-color: #000;
            color: #fff;
            border-bottom-left-radius: 10px;
            border-bottom-right-radius: 10px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <img src="https://i.imgur.com/jvgtXlL.png" alt="Fabrzy Logo">
        </div>
        <div class="content">
            <h1>Welcome to Fabrzy!</h1>
            <p>Hi there!</p>
            <p>Thank you for subscribing to my email list. I am excited to have you on board!</p>
            <p>I strive to bring you the best content on tech , life, and fitness updates regularly. You'll be the first to know whenever we have a new post on our website.</p>
            <p>Stay tuned for updates and thank you once again for joining my community.</p>
            <p>Best regards,<br> Fabio Espinoza</p>
        </div>
        <div class="footer">
            &copy; 2024 Fabrzy. All rights reserved.
        </div>
    </div>
</body>
</html>
`
)

//
