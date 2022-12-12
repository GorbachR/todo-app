/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        open: ['"Open Sans", "Segoe UI", Tahoma, sans-serif'],
        montserrat: ["Montserrat, sans-serif"],
      },
    },
  },
  plugins: [],
};
