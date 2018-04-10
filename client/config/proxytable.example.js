const proxy = {
  target: 'http://example.com/',
  changeOrigin: true,
  auth: 'api:StrongPa$$W0RD',
}

module.exports = {
  '/api': proxy,
  '/local': proxy,
  '/upload': proxy,
}
