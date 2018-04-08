<template>
  <div class="sites">
    <div class="controls">
      <button @click="start" v-if="!pinging">Start</button>
      <button @click="stop" v-if="pinging">Stop</button>
    </div>
    <div
      class="site"
      :class="getClass(index)"
      v-for="(site, index) in sites"
      :key="index"
      :title="site.domain"
    >
      {{ site.id }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'Sites',
  computed: {
    sites() {
      return this.$store.getters.sites;
    },
    current() {
      return this.$store.getters.current;
    },
    pinging() {
      return this.$store.getters.pinging;
    },
  },
  methods: {
    start() {
      this.$store.dispatch('start');
    },
    stop() {
      this.$store.dispatch('stop');
    },
    getClass(index) {
      return {
        active: index === this.current,
        success: this.sites[index].status === 'success',
        warning: this.sites[index].status === 'warning',
        danger: this.sites[index].status === 'danger',
        pinging: this.sites[index].status === 'pinging',
      };
    },
  },
};
</script>

<style lang="sass" scoped>
.controls button
  font-size: 150%

.sites div.site
  float: left
  height: 40px
  width: 40px
  border: 1px solid #eee

.site.active
  border: 1px solid green !important

.site.pinging
  background-color: #9dd1a1

.site.success
  background-color: lime

.site.warning
  background-color: yellow

.site.danger
  background-color: red

</style>
