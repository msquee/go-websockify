const { colors: { ...colors } } = require('tailwindcss/defaultTheme')

module.exports = {
  purge: [],
  theme: {
    extend: {},
    fontFamily: {
      mono: ["Iosevka"],
    },
    colors: {
      grafana: 'var(--color-grafana)',
      prometheus: 'var(--color-prometheus)',
      ...colors,
    },
  },
  variants: {
    backgroundColor: ['responsive', 'hover', 'focus', 'active']
  },
  plugins: [],
};
