<template>
  <AddPenyakit />
</template>

<script>
import AddPenyakit from "./components/AddPenyakit.vue";

export default {
  name: "App",
  components: {
    AddPenyakit,
  },
};
</script>

<style>
#app {
  font-family: "Quicksand", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 40px;
}
</style>
