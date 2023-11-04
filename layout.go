package main

var layout = `
<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
body {
    margin: 0 auto;
    max-width: 800px;
    padding-left: 50px;
    padding-right: 50px;
    padding-top: 50px;
    padding-bottom: 50px;
    hyphens: auto;
    overflow-wrap: break-word;
    text-rendering: optimizeLegibility;
    font-kerning: normal;
    font-size: 1.7em;
}
@media (max-width: 600px) {
    body {
        font-size: 1.7em;
        padding: 12px;
    }
    h1 {
        font-size: 1.8em;
    }
}

* {
  box-sizing: border-box;
}

/* Create two equal columns that floats next to each other */
.column {
  float: left;
  width: 50%;
  padding: 10px;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
}

</style>
</head>
<body>

<div class="row">
  <div class="column">
    {{range $i, $e := .Exercises}}
        {{if (lt $i 15)}}
            <p>{{printf "%2d" .X}} {{.Sign}} {{printf "%2d" .Y}} = </p>
        {{end}}
    {{end}}
  </div>
  <div class="column">
    {{range $i, $e := .Exercises}}
        {{if (gt $i 14)}}
            <p>{{printf "%2d" .X}} {{.Sign}} {{printf "%2d" .Y}} = </p>
        {{end}}
    {{end}}
  </div>
</div>

</body>
</html>
`
