<template>
  <label>
    <b-field>
      <section>
        <b-field class="center">
          <b-radio-button
            v-model="type"
            native-value="relative"
            size="is-small"
          ><b-icon icon="clock"></b-icon>
            <span>Relative</span>
          </b-radio-button>

          <b-radio-button
            v-model="type"
            native-value="absolute"
            size="is-small"
          ><b-icon icon="calendar"></b-icon>
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
import { IntervalRegexp } from '@/const/const';

export default Vue.extend({
  props: {
    startTime: {
      type: Date,
    },
    endTime: {
      type: Date,
    },
    interval: {
      type: String,
      default: '15s',
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
      hasError: false,
    };
  },
  watch: {
    type(newValue) {
      if (newValue === 'relative') {
        this.start = null;
        this.end = null;

        this.name = '15s';

        return;
      }

      this.name = '0';
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
    name: {
      get() {
        return this.interval;
      },
      set(newValue: string) {
        if (!IntervalRegexp.test(newValue)) {
          this.hasError = true;
          return;
        }

        this.hasError = false;

        this.$emit('setInterval', newValue);
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
    if (this.interval === '0') {
      this.type = 'absolute';
    }
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
