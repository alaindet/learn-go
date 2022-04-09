package cyoaweb

var defaultHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>
<body>

  <h1>{{ .Title }}</h1>

  {{ range .Paragraphs }}
    <p>{{ . }}</p>
  {{ end }}

  <ul>
    {{ range .Options }}
      <li>
        <a href="/{{ .Chapter }}">{{ .Text }}</a>
      </li>
    {{ end }}
  </ul>

</body>
</html>`
