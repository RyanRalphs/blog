<template>
  <v-app id="inspire">
    <v-app-bar id="#nav" app theme="light" flat>
      <v-container class="py-0 fill-height">
        <v-spacer></v-spacer>

        <v-btn v-for="link in links" :key="link" text color="white">
          <button class="btn btn--link" @click="aMethod">{{ link }}</button>
        </v-btn>

        <v-spacer></v-spacer>
      </v-container>
    </v-app-bar>

    <v-main class="grey lighten-3" color="primary">
      <v-parallax
        src="https://digitalsynopsis.com/wp-content/uploads/2017/07/beautiful-color-ui-gradients-backgrounds-pinky.png"
      >
        <v-row align="center" justify="center">
          <v-col class="text-center" cols="12">
            <h1 class="text-h4 font-weight-thin mb-4">Jessica's Blog</h1>
            <h4 class="subheading">Author, Poet and Journalist</h4>
          </v-col>
        </v-row>
      </v-parallax>

      <v-container>
        <v-card elevation="2" height="100"></v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      links: ["Home", "About", "Contact"],
    };
  },

  methods: {
    aMethod() {
      console.log(this.$vuetify);
    },
  },
};
</script>

<style>
.v-toolbar__content {
  background: linear-gradient(
    90deg,
    rgba(221, 94, 128, 1) 10%,
    rgba(247, 187, 151, 1) 93%
  );
}
.v-main__wrap {
  background: url("https://mixkit.imgix.net/art/preview/mixkit-starry-night-sky-over-hills-and-water-85-original-large.png?q=80&auto=format%2Ccompress")
    no-repeat center center fixed !important;
  background-size: cover;
}
</style>
