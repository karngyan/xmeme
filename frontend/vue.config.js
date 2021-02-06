module.exports = {
  configureWebpack: {
    module: {
      rules: [
        {
          parser: {
            amd: false
          }
        }
      ]
    }
  }
}