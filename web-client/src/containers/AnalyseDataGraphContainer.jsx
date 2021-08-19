import React, { useState } from "react";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";

const AnalyseDataGraphContainer = (props) => {
  //TODO make it dynamic
  const dyeOptions = [
    { label: "FAM", value: "FAM" },
    { label: "VIC", value: "VIC" },
  ];

  //local state to maintain selected dye
  const [selectedDye, setSelectedDye] = useState(dyeOptions[0]);

  //TODO get data from reducer
  const data = {};

  const onDyeChanged = (value) => {
    setSelectedDye(value);
  };

  return (
    <AnalyseDataGraphComponent
      data={data}
      dyeOptions={dyeOptions}
      selectedDye={selectedDye}
      onDyeChanged={onDyeChanged}
    />
  );
};

export default React.memo(AnalyseDataGraphContainer);
