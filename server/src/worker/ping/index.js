'use strict'

const rp = require('request-promise')

// const urls = ['https://ya.ru']
const urls = ['https://ya.ru', 'https://portal.arealidea.ru', 'https://google.com', 'https://yandex.ru', 'https://arealidea.ru']
const timeout = 5;

urls.forEach((item, index) => {
  rp({
    uri: item,
    resolveWithFullResponse: true,
    timeout: timeout * 1000,
    time: true,
  })
    .then((res) => {
      // console.log('index', item, index, res.statusCode, res.timings, res.timingPhases)
      console.log('index', item, index, res.statusCode, res.timingPhases.total)
    })
    .catch((err) => {
      if (err.message === 'Error: ETIMEDOUT') {
        console.log(`Timeout error. Response longer than ${timeout} seconds`, item, index)
      } else {
        console.log('error', err, item, index)
      }
    })
})