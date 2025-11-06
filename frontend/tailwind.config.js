/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#006493',
          light: '#cae6ff',
          dark: '#001e30',
        },
        secondary: {
          DEFAULT: '#50606e',
          light: '#d3e5f5',
          dark: '#0c1d29',
        },
      },
    },
  },
  plugins: [],
}
