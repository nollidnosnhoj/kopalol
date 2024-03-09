/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,templ,js}",'node_modules/preline/dist/*.js',],
  theme: {
    extend: {},
  },
  plugins: [
    require('preline/plugin'),
  ],
}