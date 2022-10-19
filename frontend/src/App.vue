<script>
import Home from "./Home.vue";
import About from "./About.vue";

const routes = {
  "/": Home,
  "/about": About,
};

export default {
  data() {
    return {
      currentPath: window.location.hash,
    };
  },
  computed: {
    currentView() {
      return routes[this.currentPath.slice(1) || "/"] || 404;
    },
  },
  mounted() {
    window.addEventListener("hashchange", () => {
      this.currentPath = window.location.hash;
    });
  },
};
</script>

<template>
  <v-app id="inspire">
    <v-app-bar id="#nav" app theme="light" flat>
      <v-container class="py-0 fill-height">
        <v-spacer></v-spacer>

        <a class="btn btn--link" style="color: white" href="#/">Home</a>

        <a class="btn btn--link" style="color: white" href="#/about">About</a>

        <v-spacer></v-spacer>
      </v-container>
    </v-app-bar>
    <component :is="currentView" />
  </v-app>
</template>
