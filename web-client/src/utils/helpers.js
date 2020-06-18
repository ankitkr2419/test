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


export const convertStringToSeconds = (timeString) => {
	if (timeString.indexOf(':') !== -1) {
		const a = timeString.split(':'); // split it at the colons

		// minutes are worth 60 seconds.
		return parseInt(a[0] * 60 + a[1], 10);
	}
	return 0;
};

export const convertSecondsToString = (seconds) => {
	const min = Math.floor(seconds / 60);
	const sec = seconds - min * 60;
	return `${min}:${sec}`;
};
