/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/**/*.{html,js,go,templ}"],
  theme: {
    container: {
      center: true,
      padding: "36px",
      screens: {
        "2xl": "1200px",
      },
    },
  },
  plugins: [],
};
