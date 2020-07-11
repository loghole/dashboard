<template>
  <div class="table-container" v-if="messages.length > 0">
    <table v-if="hasTags" class="table is-striped is-narrow is-hoverable is-fullwidth">
      <thead>
      <tr>
        <th>Time</th>
        <th v-for="(tag, i) in activeTags" :key="`header_${i}`">
          {{ tag }}
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(message, i) in messages" :key="`message_${i}`">
        <td style="max-width: 166px">{{ prepareText('time', message.time) }}</td>
        <td v-for="(tag, i) in activeTags" :key="`tag_${i}`">
          {{ prepareText(tag, message[tag]) }}
        </td>
      </tr>
      </tbody>
    </table>

    <!-- выводится время и все содержимое, типа такого: level: INFO, nsec: 123, message: text -->
    <!-- а пока костыль -->
    <table v-else class="table is-striped is-narrow is-hoverable is-fullwidth">
      <thead>
        <tr>
          <th>New Time</th>
          <th>_source</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(message, i) in messages" :key="`message_${i}`">
          <td style="max-width: 166px">{{ prepareText('time', message.time) }}</td>
          <td>{{ crutch(message) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
  <b-skeleton
    size="is-large"
    :active="true"
    :count="20"
    v-else
  ></b-skeleton>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  props: {
    activeTags: {
      type: Array,
      required: true,
    },
    messages: {
      type: Array,
      required: true,
    },
  },
  computed: {
    hasTags() {
      return this.activeTags.length !== 0;
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
    crutch(data: object): string {
      return `time: ${this.prepareText('time', data.time)}
              level: ${this.prepareText('level', data.level)}
              message: ${data.message}`;
    },
  },
});
</script>
