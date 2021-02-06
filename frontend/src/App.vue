<template>
  <TheFooter/>
  <hr class="z-40">
  <router-view/>
  <TheFooter/>
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
export default {
  components: {TheFooter}
}
</script>