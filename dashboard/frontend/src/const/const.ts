const OperatorMultiple = ['=', '!=', 'LIKE', 'NOT LIKE'] as string[];

const SingleParam = ['<', '<=', '>', '>='] as string[];

const OperatorsSingle = OperatorMultiple.concat(SingleParam) as string[];

export { OperatorsSingle, OperatorMultiple, SingleParam };
