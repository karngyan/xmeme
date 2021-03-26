<template>
  <div>
    <TheNavbar v-if="$route.path !== '/configure'"/>
    <hr class="z-40">
    <router-view/>
    <TheFooter  v-if="$route.path !== '/configure'"/>
  </div>
</template>

<style lang="scss">
/*noinspection CssInvalidAtRule*/
@tailwind base;
/*noinspection CssInvalidAtRule*/
@tailwind components;
/*noinspection CssInvalidAtRule*/
@tailwind utilities;

@import url('https://fonts.googleapis.com/css2?family=Oxygen:wght@300;400;700&family=Ubuntu&display=swap');
</style>

<script>
import TheFooter from "@/components/TheFooter";
import TheNavbar from "@/components/TheNavbar";

export default {
  components: {TheNavbar, TheFooter}
}
</script>