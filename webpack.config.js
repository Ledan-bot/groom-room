const {resolve} = require("path")


module.exports = {
  entry: {
    index: "./src/index.tsx"
  },
  output: {
    path: resolve(__dirname, "public"),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"]
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "babel-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.css$/i,
        include: resolve(__dirname, 'src'),
        use: ['style-loader', 'css-loader']
      }
    ]
  }
}