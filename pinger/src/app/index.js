const express = require('express')
const app = express()

app.get('/', (request, response) => {
  response.render('home', {
    name: 'John'
  })
})

app.listen(3000)