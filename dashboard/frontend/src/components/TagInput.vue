<template>
  <div>
    <b-taginput
      v-model="val"
      autocomplete
      :allow-new="true"
      placeholder="Value"
      icon="label"
      :data="data"
      @typing="getFilteredTags"
    >
    </b-taginput>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  props: {
    value: {
      type: Array,
      required: true,
    },
    type: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      data: [] as string[],
      filteredTags: [] as string[],
    };
  },
  computed: {
    val: {
      get() {
        return this.value;
      },
      set(newValue: string[]): void {
        this.$emit('input', newValue);
      },
    },
  },
  methods: {
    getFilteredTags(text: string) {
      this.filteredTags = this.data.filter(
        (option) => option
          .toString()
          .toLowerCase()
          .indexOf(text.toLowerCase()) >= 0,
      );
    },
    getFromServer() {
      Vue.axios
        .post(`/api/v1/suggest/${this.type}`, {})
        .then((response) => {
          console.log(response.data.data);
        })
        .catch((e) => {
          console.error(e);
        });
    },
  },
  mounted() {
    this.getFromServer();
  },
});
</script>
