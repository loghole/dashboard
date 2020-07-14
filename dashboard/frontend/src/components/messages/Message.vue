<template>
  <tr>
    <td style="width: 166px">{{ buildDatetime(message.time) }}</td>
    <template v-if="activeTags.length > 0">
      <td v-for="(tag, i) in activeTags" :key="`tag_${i}`">
        {{ prepareText(tag, message[tag]) }}
      </td>
    </template>
    <template v-else>
      <b-taglist attached>
        <span
          v-for="(value, name, i) in message"
          :key="`messages_${i}`"
          class="field">
          <b-tag
            v-if="showField(name)"
            type="is-info">{{ name }}:
          </b-tag>
          <b-tag
            v-if="showField(name)">
            {{ prepareText(name, value) }}
          </b-tag>
        </span>
      </b-taglist>
    </template>
  </tr>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  props: {
    message: {
      type: Object,
      required: true,
    },
    activeTags: {
      type: Array,
      required: true,
    },
  },
  methods: {
    prepareText(tag: string, text: string): string {
      switch (tag) {
        case 'time':
          return this.buildDatetime(text);
        case 'level':
          return text.toUpperCase();
        default:
          return text;
      }
    },
    buildDatetime(text: string): string {
      const date = new Date(text);

      return `${date.getDay()}.${date.getMonth()}.${date.getFullYear()} @
      ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
    },
    showField(name: string): boolean {
      return name !== 'params';
    },
  },
});
</script>

<style lang="scss" scoped>
  .field{
    margin: 0 0 0 10px;
  }
</style>
