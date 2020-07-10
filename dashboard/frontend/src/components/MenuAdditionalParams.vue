<template v-if="showAdditionalParam">
  <label>
  <!-- host -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput
        v-model="computedHostValue"
        type="host"
        :with-suggestions="true">
      </TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedHostOperator"
        name="Host"
      ></LableList>
    </b-field>
    <!-- // host -->

    <!-- Build commit -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput
        v-model="computedBuildCommitValue"
        type="build commit"
        ></TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedBuildCommitOperator"
        name="Build commit"
      ></LableList>
    </b-field>
    <!-- // Build commit -->

    <!-- Config Hash -->
    <b-field label-position="on-border" class="is-relative">
      <TagInput
        v-model="computedConfigHashValue"
        type="config hash"
        ></TagInput>

      <LableList
        :isMultiple="true"
        v-model="computedConfigHashOperator"
        name="Config hash"
      ></LableList>
    </b-field>
    <!-- // Config Hash -->
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
    hostValue: {
      type: Array,
      required: true,
    },
    hostOperator: {
      type: String,
      required: true,
    },
    buildCommitValue: {
      type: Array,
      required: true,
    },
    buildCommitOperator: {
      type: String,
      required: true,
    },
    configHashValue: {
      type: Array,
      required: true,
    },
    configHashOperator: {
      type: String,
      required: true,
    },
  },
  computed: {
    computedHostValue: {
      get() {
        return this.hostValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'host', newValue);
      },
    },
    computedHostOperator: {
      get() {
        return this.hostOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'host', newValue);
      },
    },
    computedBuildCommitValue: {
      get() {
        return this.buildCommitValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'buildCommit', newValue);
      },
    },
    computedBuildCommitOperator: {
      get() {
        return this.buildCommitOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'buildCommit', newValue);
      },
    },
    computedConfigHashValue: {
      get() {
        return this.configHashValue;
      },
      set(newValue) {
        this.$emit('setFormField', 'configHash', newValue);
      },
    },
    computedConfigHashOperator: {
      get() {
        return this.configHashOperator;
      },
      set(newValue) {
        this.$emit('setOperatorField', 'configHash', newValue);
      },
    },
  },
});
</script>
