// footerText can be: "aspire-from" or "selected"

export const getArray = (length, type) => {
  const array = [];
  for (let i = 0; i < length; i++) {
    array.push({
      id: i + 1,
      type: type,
      label: `${i + 1}`,
      footerText: "",
      isSelected: false,
    });
  }
  return array;
};
