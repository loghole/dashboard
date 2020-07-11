<template>
  <b-field label-position="on-border" class="is-relative">
    <TagInput
      v-if="isMultiple"
      :value="param.value.list"
      @input="setField"
      v-on:keydown.native.delete="keyDown"
      v-on:keyup.native.delete="keyUp"
      :placeholder="param.name"
    >
    </TagInput>

    <b-input
      v-else
      :placeholder="param.key"
      :value="param.value.item"
      @input="setField"
      type="text"
      v-on:keydown.native.delete="keyDown"
      v-on:keyup.native.delete="keyUp"
      class="no-z-index"
    >
    </b-input>

    <LableList
      :isMultiple="false"
      :value="param.operator"
      @input="setOperator"
      :name="param.key"
    ></LableList>
  </b-field>
</template>

<script lang="ts">
import Vue from 'vue';
import LableList from '@/components/LableList.vue';
import TagInput from '@/components/TagInput.vue';

import { SingleParam } from '@/const/const';

export default Vue.extend({
  components: {
    TagInput,
    LableList,
  },
  props: {
    param: {
      type: Object,
      required: true,
    },
    index: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      canDelete: true,
    };
  },
  computed: {
    isMultiple() {
      console.log(this.param);
      return !SingleParam.includes(this.param.operator);
    },
    isEmpty() {
      return this.param.value.list.length === 0 && this.param.value.item === '';
    },
  },
  methods: {
    setField(value: any) {
      this.canDelete = false;
      this.$emit('setJSONField', this.index, value);
    },
    setOperator(value: any) {
      this.$emit('setJSONOperator', this.index, value);
    },
    deleteJSONParam() {
      this.$emit('deleteJSONParam', this.index);
    },
    keyDown() {
      if (this.canDelete === true) {
        this.deleteJSONParam();
      }
    },
    keyUp() {
      if (this.isEmpty) {
        this.canDelete = true;
      }
    },
  },
});
</script>

<style lang="scss">
.field.has-addons .no-z-index.control .input:not([disabled]):hover {
  z-index: 0;
}

.field.has-addons .no-z-index.control .input:not([disabled]):focus {
  z-index: 0;
}

.field.has-addons .no-z-index.control .input:not([disabled]):active {
  z-index: 0;
}

.field.has-addons .no-z-index.control .input:not([disabled]):focus:hover {
  z-index: 0;
}
</style>
