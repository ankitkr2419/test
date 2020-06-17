/**
 * Purpose of this function is to convert array of elements as required by react-select component
 * It requires array of objects as [{label : <label>, value: <value>}]
 * @param {*} list => immutable list
 * @param {*} labelKey => stting
 * @param {*} valueKey =>string
 */
export const covertToSelectOption = (list, labelKey, valueKey) => {
	const arr = [];
	list.map(ele => arr.push({ label : ele.get(labelKey), value: ele.get(valueKey) }));
	return arr;
};
