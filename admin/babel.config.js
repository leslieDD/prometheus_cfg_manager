// {
//   "presets": [
//     "@vue/cli-plugin-babel/preset",
//     "@vue/babel-preset-jsx"
//   ]
// }

module.exports = {
  "parser": "babel-eslint",
  presets: [
    '@vue/app', '@vue/babel-preset-jsx', "@vue/cli-plugin-babel/preset"
  ],
  "plugins": ["@vue/cli-plugin-babel/preset",
    "@vue/babel-preset-jsx"]
  // 'plugins': [
  //   ['import', {
  //     'libraryName': 'iview',
  //     'libraryDirectory': 'src/components'
  //   }]
  // ]
}
