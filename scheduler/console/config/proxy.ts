export default {
  '/api': {
      'target': 'http://127.0.0.1:5266',
      'changeOrigin': true,
      'pathRewrite': { '^/api' : '' },
    },
}
