<template>
  <div>
    <b-taginput
      v-model="val"
      :autocomplete="autocomplete"
      :open-on-focus="autocomplete"
      :allow-new="allowNew"
      :placeholder="placeholder"
      :icon="icon"
      :data="filteredTags"
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
      default: '',
    },
    allowNew: {
      type: Boolean,
      default: true,
    },
    placeholder: {
      type: String,
      default: 'value',
    },
    icon: {
      type: String,
      default: 'label',
    },
  },
  data() {
    return {
      data: [] as string[],
      text: '' as string,
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
    autocomplete() {
      return this.type !== '';
    },
    filteredTags() {
      return this.data.filter(
        (option) => option
          .toString()
          .toLowerCase()
          .indexOf(this.text.toLowerCase()) >= 0,
      );
    },
  },
  methods: {
    getFilteredTags(text: string) {
      this.text = text;
    },
    getFromServer() {
      Vue.axios
        .post(`/api/v1/suggest/${this.type}`, {})
        .then((response) => {
          if (!Array.isArray(response.data.data)) {
            console.error('invalid response type', response);
            return;
          }

          this.data = response.data.data;
        })
        .catch((e) => {
          console.error(e);
        });
    },
  },
  mounted() {
    if (this.autocomplete) {
      this.getFromServer();
    }
  },
});
</script>
