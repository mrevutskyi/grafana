import { DataFrame, Field, FieldType, SortedVector } from '@grafana/data';

export enum SortDirection {
  Ascending,
  Descending,
}

// creates the `index` for the sorting.
// this is needed by the `SortedVector`.
// the index is an array of numbers, and it defines an order.
// at every slot in the index the values is the position of
// the sorted item.
// for example, an index of [3,1,2] means that
// in the dataframe, that has 3 rows, after sorting:
// - the third row will become the first
// - the first row will become the second
// - the second row will become the third
function makeIndex(fieldValues: string[], dir: SortDirection): number[] {
  // we first build an array which is [0,1,2,3....]
  const index = Array(fieldValues.length);
  for (let i = 0; i < index.length; i++) {
    index[i] = i;
  }

  const isAsc = dir === SortDirection.Ascending;

  index.sort((a: number, b: number): number => {
    // we need to answer this question:
    // in the field-used-for-sorting, how would we compare value-at-index-a to value-at-index-b?
    const valA = fieldValues[a];
    const valB = fieldValues[b];
    if (valA < valB) {
      return isAsc ? -1 : 1;
    }

    if (valA > valB) {
      return isAsc ? 1 : -1;
    }

    return 0;
  });

  return index;
}

// builds an array of strings that are exactly like the `tsNs` field's values.
// this builds it based on the timestamp-field, with using the optional
// nanos part
function makeTsNsValues(timeField: Field): string[] {
  const { nanos, values } = timeField;
  const output = Array(values.length);
  for (let i = 0; i < values.length; i++) {
    const nanoPart = nanos == null ? '000000' : nanos[i].toString().padStart(6, '0');
    output[i] = values[i].toString() + nanoPart;
  }

  return output;
}

// sort a dataframe that is in the Loki format ascending or descending,
// based on the nanosecond-timestamp
export function sortDataFrameByTime(frame: DataFrame, dir: SortDirection): DataFrame {
  const { fields, ...rest } = frame;

  // we use the approach used in @grafana/data/sortDataframe.
  // we cannot use it directly, because our tsNs field has a type=time,
  // so we have to build the `index` manually.

  let tsNsValues = fields.find((field) => field.name === 'tsNs')?.values;
  if (tsNsValues === undefined) {
    // if no `tsNs` field, this means we are running in dataplane-dataframes-mode.
    // in that case we will find the timestamp field, because that contains
    // the nanosecond-data, and build an artificial tsNs field.
    const timeField = fields.find((field) => field.name === 'timestamp' && field.type === FieldType.time);
    if (timeField == null) {
      throw new Error('missing timestamp field. should never happen');
    }

    tsNsValues = makeTsNsValues(timeField);
  }

  const index = makeIndex(tsNsValues, dir);

  return {
    ...rest,
    fields: fields.map((field) => ({
      ...field,
      values: new SortedVector(field.values, index).toArray(),
      nanos: field.nanos === undefined ? undefined : new SortedVector(field.nanos, index).toArray(),
    })),
  };

  return frame;
}
