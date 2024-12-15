const fs = require('fs');
const https = require('https');

const filePath = process.argv[2];

fs.readFile(filePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  const { cookie, body } = JSON.parse(data);

  const options = {
    hostname: 'www.parentsquare.com',
    path: '/sessions',
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      'Cookie': cookie
    }
  };

  const req = https.request(options, (res) => {
    let responseBody = '';

    res.on('data', (chunk) => {
      responseBody += chunk;
    });

    res.on('end', () => {
      console.log('Response status code:', res.statusCode);
      console.log('Response body:', responseBody);
    });
  });

  req.on('error', (e) => {
    console.error('Error making request:', e);
  });

  req.write(body);
  req.end();
});
