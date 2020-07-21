
interface Param {
  type: string;
  operator: string;
  key: string;
  value: ParamValue;
}

interface ParamValue {
  item: string;
  list: string[];
}

interface Form {
  startTime: Date;
  endTime: any;
  interval: string,
  namespace: string[];
  source: string[];
  traceID: string[];
  host: string[];
  level: string[];
  buildCommit: string[];
  configHash: string[];
  message: string;
}

interface SearchParam {
  key: string;
  name: string;
  type: string;
}

export { Param, Form, ParamValue, SearchParam }
