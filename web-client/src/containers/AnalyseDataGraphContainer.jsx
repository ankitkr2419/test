import React from "react";
import AnalyseDataGraphComponent from "components/AnalyseDataGraph";

const AnalyseDataGraphContainer = (props) => {
  //TODO get data from reducer
  const data = {};

  return <AnalyseDataGraphComponent data={data} />;
};

export default React.memo(AnalyseDataGraphContainer);
