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
      <template v-for="(message, i) in messages">
        <tr :key="`message_${i}`">
          <td style="width: 166px">{{ buildDatetime(message.time) }}</td>
          <td v-for="(tag, i) in activeTags" :key="`tag_${i}`">
            {{ prepareText(tag, message[tag]) }}
          </td>
        </tr>
        <!-- тут должен быть вывод json, по этому template
         <tr>
           <td colspan="3">
              тут json
           </td>
         </tr>
         -->
      </template>
      </tbody>
    </table>

    <!-- table with all message fields without params -->
    <table v-else class="table is-striped is-narrow is-hoverable is-fullwidth">
      <thead>
        <tr>
          <th>Time</th>
          <th>_source</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(message, i) in messages" :key="`message_${i}`">
          <td style="width: 166px">{{ buildDatetime(message.time) }}</td>
          <td>
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
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <section v-else class="section">
    <div class="content has-text-grey has-text-centered">
      <p>
        <b-icon
          icon="emoticon-sad"
          size="is-large">
        </b-icon>
      </p>
      <p>Nothing here.</p>
    </div>
  ></section>
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
    showField(name: string): boolean {
      return name !== 'params';
    },
  },
});
</script>
<style lang="scss" scoped>
  .field{
    margin-right: 10px;
  }
</style>
