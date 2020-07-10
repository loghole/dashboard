<template>
  <label>
    <!-- level -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput v-model="computedLevelValue" type="level"></TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedLevelOperator"
        name="Level"
      ></LableList>
    </b-field>
    <!-- // level -->

    <!-- namespace -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput v-model="computedNamespaceValue" type="namespace"></TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedNamespaceOperator"
        name="Namespace"
      ></LableList>
    </b-field>
    <!-- // namespace -->

    <!-- source -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput v-model="computedSourceValue" type="source"></TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedSourceOperator"
        name="Source"
      ></LableList>
    </b-field>
    <!-- // source -->

    <!-- traceID -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput
        v-model="computedTraceValue"
        type="trace id"
        :with-suggestions="false">
      </TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedTraceOperator"
        name="Trace ID"
      ></LableList>
    </b-field>
    <!-- // traceID -->
  </label>
</template>

<script lang="ts">
import Vue from 'vue';
import TagInput from '@/components/TagInput.vue';
import LableList from '@/components/LableList.vue';

export default Vue.extend({
  components: {
    TagInput,
    LableList,
  },
  props: {
    levelValue: {
      type: Array,
      required: true,
    },
    levelOperator: {
      type: String,
      required: true,
    },
    namespaceValue: {
      type: Array,
      required: true,
    },
    namespaceOperator: {
      type: String,
      required: true,
    },
    sourceValue: {
      type: Array,
      required: true,
    },
    sourceOperator: {
      type: String,
      required: true,
    },
    traceValue: {
      type: Array,
      required: true,
    },
    traceOperator: {
      type: String,
      required: true,
    },
  },
  computed: {
    computedLevelValue: {
      get() {
        return this.levelValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'level', newValue);
      },
    },
    computedLevelOperator: {
      get() {
        return this.levelOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'level', newValue);
      },
    },
    computedNamespaceValue: {
      get() {
        return this.namespaceValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'namespace', newValue);
      },
    },
    computedNamespaceOperator: {
      get() {
        return this.namespaceOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'namespace', newValue);
      },
    },
    computedSourceValue: {
      get() {
        return this.sourceValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'source', newValue);
      },
    },
    computedSourceOperator: {
      get() {
        return this.sourceOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'source', newValue);
      },
    },
    computedTraceValue: {
      get() {
        return this.traceValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'traceID', newValue);
      },
    },
    computedTraceOperator: {
      get() {
        return this.traceOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'traceID', newValue);
      },
    },
  },
});
</script>
