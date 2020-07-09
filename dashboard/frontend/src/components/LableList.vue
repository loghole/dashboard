<template>
  <b-dropdown aria-role="list">
    <span
      slot="trigger"
      role="button"
      style="cursor: pointer"
      class="has-text-success"
    >{{operator}}
    </span>

    <b-dropdown-item aria-role="listitem" @click="setOperator('=')"
      >=</b-dropdown-item
    >
    <b-dropdown-item aria-role="listitem" @click="setOperator('!=')"
      >!=</b-dropdown-item
    >
    <b-dropdown-item aria-role="listitem" @click="setOperator('LIKE')"
      >LIKE</b-dropdown-item
    >
  </b-dropdown>
</template>

<script lang="ts">
import Vue from 'vue';

// Тут наверное нужны будут эвенты в родителя что бы передать оператор?
export default Vue.extend({
  data() {
    return {
      operator: '=',
    };
  },
  methods: {
    setOperator(value: string) {
      this.operator = value;
    },
  },
});
</script>
