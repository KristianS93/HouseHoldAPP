// npx tailwindcss -o ./public/static/css/styles.css

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: [
    './public/**/*.{html,js}',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwindcss'),
    require('autoprefixer'),
  ],
}


// Color scheme should most likely have:
// Primary Dark - Light
// Secondary Dark - Light
// Third Dark - Light
// Text Dark - Light
// Hover Text Dark - Light
// Hover BG Dark - Light
