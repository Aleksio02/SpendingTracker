<template>
  <div class="intro">
    <div class="intro_fade"></div>
  </div>

  <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
  created() {
    this.$router.push('/login');
  },
}
</script>

<style>

</style>
