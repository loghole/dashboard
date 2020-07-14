<template>
  <div class="columns p-2 pt-4">
    <!-- add param
    <Sidebar v-model="showAddParam" v-on:save="saveParam"></Sidebar> -->
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

      <!-- menu default params -->
      <b-field>
        <Menu
          v-on:setFormField="setFormField"
          v-on:setOperatorField="setOperatorField"
          :form="form"
          :operator="operator"
          :params="defaultParams"
        >
        </Menu>
      </b-field>
      <!-- // menu default params -->

      <!-- params -->
      <JSONValue
        v-for="(param, i) in params"
        :key="`param_${i}`"
        :param="param"
        :index="i"
        v-on:setJSONField="setJSONField"
        v-on:setJSONOperator="setJSONOperator"
        v-on:deleteJSONParam="deleteJSONParam"
      ></JSONValue>
      <!-- // params -->

      <!-- menu additional param -->
      <template v-if="showAdditionalParam">
        <b-field>
          <Menu
            v-on:setFormField="setFormField"
            v-on:setOperatorField="setOperatorField"
            :form="form"
            :operator="operator"
            :params="additionalParams"
          ></Menu>
        </b-field>
      </template>
      <!-- // menu additional param -->

      <!-- add new param  -->
      <b-field v-if="showAddParam">
        <b-input
          v-model="newParamName"
          placeholder="new param name"
          type="text"
          @keydown.native.enter="saveParam"
        >
        </b-input>
        <p class="control">
          <button class="button is-primary" @click="saveParam">Add</button>
        </p>
      </b-field>
      <!-- // add new param  -->

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

        <!-- add param -->
        <button class="button is-small is-outlined" @click="newParam">
          <b-icon icon="plus" size="is-small"> </b-icon>
          <span>param</span>
        </button>
        <!-- // add param -->
      </div>
      <b-button class="button is-primary is-fullwidth" @click="search"
        >Search</b-button
      >
    </div>

    <div class="column">
      <div class="columns">
        <!-- Showed tags -->
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
        <!-- // Showed tags -->

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

      <!-- messages table -->
      <MessagesTable
        :activeTags="showTags"
        :messages="messages"
      ></MessagesTable>
      <!-- // entry table -->
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import DateTime from '@/components/DateTime.vue';
import Menu from '@/components/menu/Menu.vue';
import MessagesTable from '@/components/messages/MessagesTable.vue';
import JSONValue from '@/components/menu/JSONValue.vue';
import {
  Param, Form, ParamValue, SearchParam,
} from '@/types/view';

import { SingleParam } from '@/const/const';
import FilterTags from '@/plugins/filter';

export default Vue.extend({
  components: {
    DateTime,
    Menu,
    MessagesTable,
    JSONValue,
  },
  data() {
    return {
      operator: {
        level: '=' as string,
        namespace: '=' as string,
        source: '=' as string,
        traceID: '=' as string,
        host: '=' as string,
        buildCommit: '=' as string,
        configHash: '=' as string,
      },
      form: {
        startTime: new Date(new Date().getTime() - 1000 * 60 * 60 * 24),
        endTime: null,
        namespace: [] as string[],
        source: [] as string[],
        traceID: [] as string[],
        host: [] as string[],
        level: [] as string[],
        buildCommit: [] as string[],
        configHash: [] as string[],
        message: '',
      } as Form,
      defaultParams: [
        { key: 'level', name: 'Level', type: 'level' },
        { key: 'namespace', name: 'Namespace', type: 'namespace' },
        { key: 'source', name: 'Source', type: 'source' },
        { key: 'traceID', name: 'Trace ID' },
      ] as SearchParam[],
      additionalParams: [
        { key: 'host', name: 'Host', type: 'host' },
        { key: 'buildCommit', name: 'Build commit' },
        { key: 'configHash', name: 'Config hash' },
      ] as SearchParam[],
      params: [] as Param[],
      showAddParam: false,
      newParamName: '',
      tagsInput: '',

      // TODO drop..?

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
      operators: ['=', '!=', '>', '<', '>=', '<=', 'LIKE', 'NOT LIKE'],
      showAdditionalParam: false,
      messages: [],
      showTags: [],
    };
  },
  computed: {
    filteredOperators(): string[] {
      return this.operators.filter(
        (option: string) => option.toLowerCase().indexOf(this.param.operator.toLowerCase()) >= 0,
      );
    },
    filteredTags(): string[] {
      return FilterTags(this.tags, this.showTags, this.tagsInput);
    },
  },
  methods: {
    saveParam(): void {
      this.showAddParam = false;

      if (this.params.findIndex((p) => p.key === this.newParamName) !== -1) {
        // TODO higlight this param // 5 sec
        return;
      }

      this.params.push({
        type: 'json',
        key: this.newParamName,
        value: { item: '', list: [] as string[] },
        operator: '=',
      });
    },
    newParam() {
      this.showAddParam = true;
      this.newParamName = '';
    },
    setStartTime(val: Date): void {
      this.form.startTime = val;
    },
    setEndTime(val: Date): void {
      this.form.endTime = val;
    },
    setFormField(key: string, val: string[]): void {
      this.form[key] = val;
    },
    setOperatorField(key: string, val: string): void {
      this.operator[key] = val;
    },
    setJSONField(idx: number, val: string | string[]): void {
      if (typeof val === 'string') {
        this.params[idx].value.item = val;
        return;
      }

      this.params[idx].value.list = val;
    },
    setJSONOperator(idx: number, val: string): void {
      if (
        !SingleParam.includes(this.params[idx].operator)
        === SingleParam.includes(val)
      ) {
        this.params[idx].value.list = [];
        this.params[idx].value.item = '';
      }

      this.params[idx].operator = val;
    },
    deleteJSONParam(idx: number): void {
      this.$nextTick(() => {
        this.params = this.params.filter((v, i) => i !== idx);
      });
    },
    getFilteredTags(text: string) {
    // console.log(this.tags, this.filteredTags);
      //   this.filteredTags = FilterTags(this.tags, this.filteredTags, text);
      this.tagsInput = text;
    },
    // TODO drop it...?
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
        { mapKey: 'namespace', key: 'namespace', value: this.form.namespace },
        { mapKey: 'level', key: 'level', value: this.form.level },
        { mapKey: 'source', key: 'source', value: this.form.source },
        { mapKey: 'trace_id', key: 'trace_id', value: this.form.traceID },
        { mapKey: 'host', key: 'host', value: this.form.host },
        {
          mapKey: 'buildCommit',
          key: 'build_commit',
          value: this.form.buildCommit,
        },
        {
          mapKey: 'configHash',
          key: 'config_hash',
          value: this.form.configHash,
        },
      ].forEach((h) => {
        if (h.value.length > 0) {
          params.push({
            type: 'column',
            key: h.key,
            operator: this.operator[h.mapKey] || '=',
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

      console.log(JSON.stringify(params));

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
  mounted() {
    Vue.axios
      .post('/api/v1/entry/list', { limit: 100 })
      .then((response) => {
        this.messages = response.data.data;

        this.setTags(response.data.data);

        console.log(response.data);
      })
      .catch((e) => {
        console.error(e);
      });
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
