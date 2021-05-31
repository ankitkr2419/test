export const getWellsInitialArray = (n, type) => {
  const wellsArray = [];
  for (let i = 0; i < n; i++) {
    wellsArray.push({
      id: i + 1,
      type: type,
      label: `${i + 1}`,
      footerText: "",
      isDisabled: false,
      isSelected: false,
    });
  }
  return wellsArray;
};
