const fs = require('fs');
const path = require('path');

const INPUT_PATH = path.join(__dirname, 'raw_repos.json');
const OUTPUT_PATH = path.join(__dirname, 'repos.json');

const rawData = fs.readFileSync(INPUT_PATH, { encoding: 'utf-8' });

const repos = JSON.parse(rawData).map(item => {
  const { name, description, html_url } = item;
  return { name, description, html_url };
});

fs.writeFileSync(OUTPUT_PATH, JSON.stringify({ repos }));
