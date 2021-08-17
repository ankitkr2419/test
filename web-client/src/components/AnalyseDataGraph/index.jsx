import React from "react";
import PropTypes from "prop-types";
import { GraphCard } from "./GraphCard";
import { LineChart } from "core-components";
import { options } from "./GraphOptions";
import Filters from "./Filters";

const AnalyseDataGraphComponent = (props) => {
  let { data } = props;
  return (
    <div>
      <GraphCard>
        <LineChart data={data} options={options} />
      </GraphCard>
      <Filters />
    </div>
  );
};

AnalyseDataGraphComponent.propTypes = {
  data: PropTypes.object.isRequired,
};

export default React.memo(AnalyseDataGraphComponent);
