module.exports = {
  mode: "jit",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      width: {
        "right-bar": "350px",
        "left-bar": "275px",
        "left-bar-m": "88px",
        "login-aside": "514px",

        main: "600px",
        "main-box": "950px",
      },
      colors: {
        blue: {
          "nav-sky": "#EAF4FE",
          "nav-text": "#4793ee",
          link: "#09f",
        },
      },
    },
  },
  plugins: [],
};
