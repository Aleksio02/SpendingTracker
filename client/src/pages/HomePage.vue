<template>
    <side-bar></side-bar>
    <router-view></router-view>
</template>

<script>
import { checkToken } from '../modules/auth';

export default {
	created() {
		checkToken();
	},
}
</script>
  
<style scoped>
.stats-content {
    margin-left: 360px !important;
}

* {
    box-sizing: content-box;
}
</style>
  
  
  