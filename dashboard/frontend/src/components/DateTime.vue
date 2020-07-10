<template>
  <label>
    <b-field>
      <section>
        <b-field class="center">
          <b-radio-button
            v-model="type"
            native-value="relative"
            @input="setTime"
            size="is-small"
          >
            <b-icon icon="clock"></b-icon>
            <span>Relative</span>
          </b-radio-button>

          <b-radio-button
            v-model="type"
            native-value="absolute"
            size="is-small"
          >
            <b-icon icon="calendar"></b-icon>
            <span>Absolute</span>
          </b-radio-button>
        </b-field>
      </section>
    </b-field>

    <b-field
      label="Interval "
      label-position="on-border"
      :type="hasError ? 'is-danger' : ''"
      :message="{ 'invalid value': hasError }"
      v-if="type === 'relative'"
    >
      <b-autocomplete
        :open-on-focus="true"
        v-model="name"
        :data="filteredDataArray"
        placeholder="e.g. 15s"
        icon="clock"
        clearable
      >
      </b-autocomplete>
    </b-field>
    <template v-else>
      <b-field label="Start time" label-position="on-border">
        <b-datetimepicker
          placeholder="Click to select..."
          :max-datetime="maxDate"
          :timepicker="{ enableSeconds: true }"
          editable
          v-model="start"
        ></b-datetimepicker>
      </b-field>

      <b-field label="End time" label-position="on-border">
        <b-datetimepicker
          placeholder="Click to select..."
          :timepicker="{ enableSeconds: true }"
          editable
          v-model="end"
        ></b-datetimepicker>
      </b-field>
    </template>
  </label>
</template>

<script lang="ts">
import Vue from 'vue';

const reg = new RegExp('^([0-9]+)(s|sec|m|min|h|hr|hour|d|day)?$', 'i');

export default Vue.extend({
  props: {
    startTime: {
      type: Date,
    },
    endTime: {
      type: Date,
    },
  },
  data() {
    return {
      maxDate: new Date(),
      type: 'relative',
      data: [
        '5s',
        '15s',
        '30s',
        '10m',
        '30m',
        '1h',
        '3h',
        '12h',
        '1d',
        '7d',
        '14d',
      ],
      name: '15s',
      hasError: false,
    };
  },
  watch: {
    name() {
      this.setTime();
    },
  },
  methods: {
    setTime() {
      if (!reg.test(this.name)) {
        this.hasError = true;
        return;
      }

      this.hasError = false;
      this.end = null;

      const values = reg.exec(this.name);
      const num = parseInt(values[1], 10);
      const t = values[2];

      let offset = 0;

      switch (t) {
        case 'm':
        case 'min':
          offset = num * 60;
          break;

        case 'h':
        case 'hr':
        case 'hour':
          offset = num * 3600;
          break;

        case 'd':
        case 'day':
          offset = num * 3600 * 24;
          break;

        default:
          offset = num;
      }

      offset *= 1000;

      const d = new Date(new Date().getTime() - offset);

      this.start = d;
    },
  },
  computed: {
    start: {
      get() {
        return this.startTime;
      },
      set(newValue) {
        this.$emit('setStartTime', newValue);
      },
    },
    end: {
      get() {
        return this.endTime;
      },
      set(newValue) {
        this.$emit('setEndTime', newValue);
      },
    },
    filteredDataArray() {
      const arr = this.data.filter(
        (option) => option
          .toString()
          .toLowerCase()
          .indexOf(this.name.toLowerCase()) >= 0,
      );

      if (arr.length === 1 && this.name === arr[0]) {
        return [];
      }

      return arr;
    },
  },
  mounted() {
    this.setTime();
  },
});
</script>

<style lang="scss" scoped>
.select {
  width: 100%;
}
.center {
    justify-content: center;

}
</style>
