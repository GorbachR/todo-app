/** @type {import('tailwindcss').Config} */
/* eslint-disable @typescript-eslint/no-var-requires */

const defaultTheme = require("tailwindcss/defaultTheme");

export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        "Josefin-Sans": ["Josefin Sans", ...defaultTheme.fontFamily.sans],
      },
      colors: {
        "todo-text": "var(--todo-text)",
        "todo-text-hover": "var(--todo-text-hover)",
        "todo-text-crossout": "var(--todo-text-crossout)",
        "todo-bright-blue": "var(--todo-bright-blue)",
        "todo-circle-hover1": "var(--todo-circle-hover1)",
        "todo-circle-hover2": "var(--todo-circle-hover2)",
        "todo-bg": "var(--todo-bg)",
        "todo-bg2": "var(--todo-bg2)",
        "todo-outline": "var(--todo-outline)",
        "todo-noidea": "var(--todo-noidea)",
      },
      backgroundImage: {
        "bg-desktop-dark": "url(./assets/bg-desktop-dark.jpg)",
        "bg-desktop-light": "url(./assets/bg-desktop-light.jpg)",
        "bg-mobile-dark": "url(./assets/bg-mobile-dark.jpg)",
        "bg-mobile-light": "url(./assets/bg-mobile-light.jpg)",
      },
    },
  },
  plugins: [],
};
