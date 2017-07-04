var express = require('express')
var app = express()

var path = require('path')

app.use(
  express.static(path.join(__dirname, 'public'))
)

var server = app.listen(3000)
console.log('Servidor iniciado na porta: %s', server.address().port)
