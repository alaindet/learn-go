// https://nodejs.org/en/knowledge/HTTP/servers/how-to-serve-static-files/
const fs = require('fs');
const http = require('http');
const path = require('path');

const PORT = 8080;
const ASSETS_DIR = path.join(__dirname, 'assets');

http.createServer(fileServer).listen(PORT);

function fileServer(req, res) {

  const filePath = getFilePath(req.url);

  fs.readFile(filePath, function (err, data) {

    // File not found
    if (err) {
      res.writeHead(404);
      res.end(JSON.stringify(err));
      return;
    }

    // Is it a WASM file?
    if (filePath.endsWith("wasm")) {
      res.setHeader("Content-Type", "application/wasm");
    }

    res.writeHead(200);
    res.end(data);
  });
}

function getFilePath(url) {

  // Is it index.html?
  if (url === '/') {
    return path.join(ASSETS_DIR, 'index.html');
  }

  return path.join(ASSETS_DIR, url);
}
