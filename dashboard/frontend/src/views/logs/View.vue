<template>
  <div class="columns p-2 pt-4">
    <!-- add param -->
    <Sidebar v-model="showAddParam" v-on:save="saveParam"></Sidebar>
    <!-- add param -->

    <div class="column page-menu">
      <!-- date -->
      <b-field>
        <DateTime
          :startTime="form.startTime"
          :end-time="form.endTime"
          v-on:setStartTime="setStartTime"
          v-on:setEndTime="setEndTime"
        ></DateTime>
      </b-field>
      <!-- // date -->

      <BaseMenuFields
        v-on:setFormField="setFormField"
        v-on:setOperatorField="setOperatorField"
        :level-value="form.level"
        :level-operator="operator.level"
        :namespace-value="form.namespace"
        :namespace-operator="operator.namespace"
        :source-value="form.source"
        :source-operator="operator.source"
        :trace-value="form.traceID"
        :trace-operator="operator.traceID">
      </BaseMenuFields>

      <!-- params -->
      <b-field
        v-for="(param, i) in params"
        :label="`${param.key} ${param.operator}`"
        :key="`param_${i}`"
        label-position="on-border"
      >
        <b-taginput
          v-if="isListValue(param.operator)"
          v-model="param.value.list"
          autocomplete
          :allow-new="true"
          placeholder="Value"
          icon="label"
          icon-right="close-circle"
          icon-right-clickable
          @icon-right-click="removeParam(i)"
        >
        </b-taginput>
        <b-input
          v-else
          :placeholder="param.key"
          v-model="param.value.item"
          type="text"
          icon-right="close-circle"
          icon-right-clickable
          @icon-right-click="removeParam(i)"
        >
        </b-input>
      </b-field>
      <!-- // params -->

      <template v-if="showAdditionalParam">
        <!-- host -->
        <b-field label="Host" label-position="on-border">
          <TagInput v-model="form.host" type="host"> </TagInput>
        </b-field>
        <!-- // host -->

        <!-- Build commit -->
        <b-field label="Build commit" label-position="on-border">
          <b-taginput
            v-model="form.buildCommit"
            autocomplete
            :allow-new="true"
            placeholder="Build commit"
            icon="label"
          >
          </b-taginput>
        </b-field>
        <!-- // Build commit -->

        <!-- Config Hash -->
        <b-field label="Config hash" label-position="on-border">
          <b-taginput
            v-model="form.configHash"
            autocomplete
            :allow-new="true"
            placeholder="Config hash"
            icon="label"
          >
          </b-taginput>
        </b-field>
        <!-- // Config Hash -->
      </template>

      <div class="buttons is-centered">
        <button
          class="button is-small is-outlined"
          @click="showAdditionalParam = !showAdditionalParam"
        >
          <b-icon
            :icon="showAdditionalParam ? 'eye-off' : 'eye'"
            size="is-small"
          >
          </b-icon>
          <span>other</span>
        </button>
        <button
          class="button is-small is-outlined"
          @click="showAddParam = true"
        >
          <b-icon icon="plus" size="is-small"> </b-icon>
          <span>param</span>
        </button>
      </div>
      <b-button class="button is-primary is-fullwidth" @click="search"
        >Search</b-button
      >
    </div>

    <div class="column">
      <div class="columns">
        <div class="column">
          <b-taginput
            v-model="showTags"
            :data="filteredTags"
            autocomplete
            :allow-new="true"
            :open-on-focus="true"
            placeholder="Showed tags"
            icon="label"
            @typing="getFilteredTags"
          >
          </b-taginput>
        </div>
        <div class="column">
          <b-field label="Search" label-position="on-border">
            <b-input
              placeholder="Search..."
              type="search"
              icon="magnify"
              icon-clickable
              class="w100"
              v-model="form.message"
            ></b-input>
            <p class="control">
              <b-button class="button is-primary">Search</b-button>
            </p>
          </b-field>
        </div>
      </div>

      <div class="table-container" v-if="messages.length > 0">
        <table class="table is-striped is-narrow is-hoverable is-fullwidth">
          <thead>
            <tr>
              <th>Time</th>
              <th>Level</th>
              <th>Message</th>
              <th v-for="(tag, i) in showTags" :key="`header_${i}`">
                {{ tag }}
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(message, i) in messages" :key="`message_${i}`">
              <td style="max-width: 166px">
                {{ new Date(message.time).toLocaleString() }}
              </td>
              <td>{{ message.level.toUpperCase() }}</td>
              <td>{{ message.message }}</td>
              <td v-for="(tag, i) in showTags" :key="`tag_${i}`">
                {{ message[tag] }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <b-skeleton
        size="is-large"
        :active="loading"
        :count="20"
        v-else
      ></b-skeleton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Sidebar from '@/components/Sidebar.vue';
import TagInput from '@/components/TagInput.vue';
import DateTime from '@/components/DateTime.vue';
import BaseMenuFields from '@/components/BaseMenuFields.vue';
import { Param, Form, ParamValue } from '@/types/view';

export default Vue.extend({
  components: {
    Sidebar,
    TagInput,
    DateTime,
    BaseMenuFields,
  },
  data() {
    return {
      loading: true,
      operator: {
        level: '=' as string,
        namespace: '=' as string,
        source: '=' as string,
        traceID: '=' as string,
      },
      form: {
        startTime: new Date(new Date().getTime() - 1000 * 60 * 60 * 24),
        endTime: null,
        namespace: [] as string[],
        source: [] as string[],
        traceID: [] as string[],
        host: [] as string[],
        level: [] as string[],
        buildCommit: '',
        configHash: '',
        message: '',
      } as Form,
      params: [] as Param[],
      param: {
        operator: '',
        type: '',
        key: '',
        value: {
          item: '',
          list: [] as string[],
        } as ParamValue,
      } as Param,
      maxDatetime: new Date(),
      sources: [],
      hosts: [],
      namespaces: [],
      levels: [],
      tags: [] as string[],
      filteredTags: [] as string[],
      operators: ['=', '!=', '>', '<', '>=', '<=', 'LIKE', 'NOT LIKE'],
      showAddParam: false,
      showAdditionalParam: false,
      messages: [],
      showTags: ['trace_id'],
    };
  },
  computed: {
    filteredOperators(): string[] {
      return this.operators.filter(
        (option: string) => option.toLowerCase().indexOf(this.param.operator.toLowerCase()) >= 0,
      );
    },
  },
  methods: {
    setStartTime(val: Date): void {
      this.form.startTime = val;
    },
    setEndTime(val: Date): void {
      this.form.endTime = val;
    },
    setFormField(key: string, val: any): void {
      console.log(key, val);

      this.form[key] = val;
    },
    setOperatorField(key: string, val: string): void {
      console.log(key, val);

      this.operator[key] = val;
    },
    getSourceList(val: string): void {
      console.log(val);
    },
    getHostList(val: string): void {
      console.log(val);
    },
    getNamespaceList(val: string): void {
      console.log(val);
    },
    getLevelList(val: string): void {
      console.log(val);
    },
    getFilteredTags(text: string) {
      this.filteredTags = this.tags.filter(
        (option) => option
          .toString()
          .toLowerCase()
          .indexOf(text.toLowerCase()) >= 0,
      );
    },
    saveParam(param: Param): void {
      this.showAddParam = false;

      this.params.push({
        type: param.type,
        key: param.key,
        value: param.value,
        operator: param.operator,
      });
    },
    removeParam(idx: number): void {
      this.$nextTick(() => {
        this.params = this.params.filter((v, i) => i !== idx);
      });
    },
    isListValue(operator: string): boolean {
      return ['=', '!=', 'LIKE', 'NOT LIKE'].includes(operator);
    },
    search(): void {
      const params = [
        {
          type: 'column',
          key: 'time',
          operator: '>=',
          value: {
            item: parseInt(
              (this.form.startTime.getTime() / 1000).toString(),
              10,
            ).toString(),
          } as ParamValue,
        },
      ] as Param[];

      if (this.form.endTime !== null) {
        params.push({
          type: 'column',
          key: 'time',
          operator: '<=',
          value: { item: this.form.endTime } as ParamValue,
        });
      }

      if (this.form.message !== '') {
        params.push({
          type: 'column',
          key: 'message',
          operator: 'LIKE',
          value: { item: this.form.message } as ParamValue,
        });
      }

      [
        { key: 'namespace', value: this.form.namespace },
        { key: 'level', value: this.form.level },
        { key: 'source', value: this.form.source },
        { key: 'trace_id', value: this.form.traceID },
        { key: 'host', value: this.form.host },
        { key: 'build_commit', value: this.form.buildCommit },
        { key: 'config_hash', value: this.form.configHash },
      ].forEach((h) => {
        if (h.value.length > 0) {
          params.push({
            type: 'column',
            key: h.key,
            operator: this.operator[h.key] || '=',
            value: { list: h.value } as ParamValue,
          });
        }
      });

      this.params.forEach((param) => {
        params.push({
          type: 'key',
          key: param.key,
          operator: param.operator,
          value: param.value,
        });
      });

      console.log(this.form.startTime);

      Vue.axios
        .post('/api/v1/entry/list', { params, limit: 100 })
        .then((response) => {
          this.messages = response.data.data;

          this.setTags(response.data.data);

          console.log(response.data);
        })
        .catch((e) => {
          console.error(e);
        });
    },
    setTags(list: Array<any>): void {
      const h = {} as Record<string, boolean>;

      list.forEach((l: Record<string, any>) => {
        Object.keys(l).forEach((k) => {
          h[k] = true;
        });
      });

      this.tags = Object.keys(h);
    },
  },
});
</script>

<style lang="scss" scoped>
.p-2 {
  padding: 0.5rem;
}
.pt-4 {
  padding-top: 1rem;
}

.w100 {
  width: 100%;
}

.page {
  &-menu {
    max-width: 210px;
    min-width: 150px;
  }
}

.is-relative {
  position: relative;
}
</style>
