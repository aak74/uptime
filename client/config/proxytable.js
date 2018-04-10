const proxy = {
  target: 'http://localhost:80/',
  changeOrigin: true,
}

module.exports = {
  '/api': proxy,
  '/local': proxy,
  '/upload': proxy,
}
