<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  methods: {
    getVersion() {
      this.axios
        .get('/api/v1/info')
        .then((response) => {
          const info = response.data.data;

          Object.keys(info).forEach((key) => {
            console.log(`${key}: ${info[key]}`);
          });
        })
        .catch((e) => {
          console.error(e);
        });
    },
  },
  mounted() {
    this.getVersion();
  },
});
</script>
