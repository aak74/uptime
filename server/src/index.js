'use strict'

const type = process.env.PROCESS_TYPE

if (type === 'web') {
  require('./web')
} else if (type === 'ping-worker') {
  require('./worker/ping')
} else {
  throw new Error(`
    ${type} is an unsupported process type. 
    Use one of: 'web', 'ping-worker'!
  `)
}
