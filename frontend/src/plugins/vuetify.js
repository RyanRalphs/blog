import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify, {
  options: {
    customProperties: true,
  },
});

const vuetify = new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#E69B9C",
        secondary: "#A3FFF3",
        accent: "#EEA4ED",
        error: "#f44336",
        warning: "#ffc107",
        info: "#03a9f4",
        success: "#8bc34a",
      },
      dark: {
        primary: "#E69B9C",
        secondary: "#A3FFF3",
        accent: "#EEA4ED",
        error: "#f44336",
        warning: "#ffc107",
        info: "#03a9f4",
        success: "#8bc34a",
      },
    },
  },
});

export default vuetify;
