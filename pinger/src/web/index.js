'use strict'

const express = require('express')
// const bodyParser = require('body-parser')
// const cors = require('cors')
// const morgan = require('morgan')
const config = require('./config')
const app = express()
const projects = require('../mock/sites')

// app.use(morgan('combined'))
// app.use(bodyParser.json())
// app.use(cors())

app.get('/', (request, response) => {
  console.log('/', request);
  response.send('Hello from Express!')
  
})

app.get('/api/v1/projects', (request, response) => {
  console.log('/projects', request.url);
  response.json(projects)
})

app.listen(process.env.PORT || config.port,
  () => console.log(`Server start listen on port ${config.port} ...`)
)