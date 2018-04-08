<template>
  <div id="app">
    <div v-if="loading">
      Идет загрузка
    </div>
    <sites v-if="!loading"/>
  </div>
</template>

<script>
import Sites from './components/Sites';

export default {
  name: 'app',
  components: {
    Sites,
  },
  computed: {
    loading() {
      return this.$store.getters.status.loading;
    },
  },
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
