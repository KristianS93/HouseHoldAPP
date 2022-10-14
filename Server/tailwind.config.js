// npx tailwindcss -o ./public/static/css/styles.css

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: [
    './public/**/*.{html,js}',
  ],
  theme: {
    extend: {
      colors: {
        primary_dark: "#222831",
        secondary_dark: "#393E46",
        text_dark: "#bfc7d5",
        text_dark_hover: "#00ADB5",
        bg_dark_hover: "#EEEEEE",
        
        primary_light: "#E3FDFD",
        secondary_light: "#CBF1F5",
        text_light: "#1a1c23",
        text_light_hover: "#A6E3E9",
        bg_light_hover: "#71C9CE",

        hyperlink: "#82a0d9",
      },
    },
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
