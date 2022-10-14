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
        primary_dark: "#292d3e",
        secondary_dark: "#8723ec",
        text_dark: "#bfc7d5",
        text_dark_hover: "#ffffff",
        bg_dark_hover: "#4c1385",
        
        primary_light: "#f6f6f8",
        secondary_light: "#8a90b7",
        text_light: "#1a1c23",
        text_light_hover: "#000125",
        bg_light_hover: "#ffffff",

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
