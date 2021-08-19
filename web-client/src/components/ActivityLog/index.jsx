import React, { useState } from "react";
import { Table } from "core-components";
import { ButtonIcon } from "shared-components";
import "./activity.scss";
import ActivityData from "./ActivityData.json";
import SearchBar from "./SearchBar";
import moment from "moment";

const ActivityComponent = (props) => {
  let { experiments, searchText, onSearchTextChanged, expandLogHandler } =
    props;

  const [selectedActivity, setSelectedActivity] = useState(null);
  const headers = ActivityData.headers;

  const toggleSelectedActivity = (experiment) => {
    if (selectedActivity?.id === experiment.id) {
      setSelectedActivity(null);
    } else {
      setSelectedActivity(experiment);
    }
  };

  let filteredExperiments = experiments?.filter((experiment) =>
    experiment.template_name.toLowerCase().includes(searchText.toLowerCase())
  );

  return (
    <div className="activity-content h-100 py-0">
      <SearchBar
        id="search"
        name="search"
        placeholder="Search"
        value={searchText}
        onChange={(e) => onSearchTextChanged(e.target.value)}
      />
      <div className="table-responsive">
        <Table striped className="table-log">
          <colgroup>
            <col width="9%" />
            <col />
            <col width="12%" />
            <col width="10.5%" />
            <col width="10.5%" />
            <col width="8%" />
            <col width="8%" />
            <col width="12%" />
            <col width="15%" />
          </colgroup>
          <thead>
            <tr>
              {headers.map((header, i) => (
                <th key={i}>{header}</th>
              ))}
              <th />
            </tr>
          </thead>
          <tbody>
            {filteredExperiments &&
              filteredExperiments.map((experiment, i) => (
                <tr
                  className={
                    experiment.id === selectedActivity?.id ? "active" : ""
                  }
                  key={i}
                  onClick={() => toggleSelectedActivity(experiment)}
                >
                  <td>{i + 1}</td>
                  <td>{experiment.template_name}</td>
                  <td>
                    {experiment.created_at &&
                      moment(experiment.created_at).format("DD/MM/YYYY")}
                  </td>
                  <td>
                    {experiment.start_time &&
                      moment(experiment.start_time).format("HH:MM A")}
                  </td>
                  <td>
                    {experiment.end_time &&
                      moment(experiment.end_time).format("HH:MM A")}
                  </td>
                  <td>{experiment.well_count}</td>
                  <td>{experiment.repeat_cycle}</td>
                  <td
                    className={
                      experiment.result === "Error"
                        ? "text-danger text-capitalize"
                        : "text-capitalize"
                    }
                  >
                    {experiment.result ? experiment.result : "NA"}
                  </td>
                  <td className="td-actions">
                    <ButtonIcon
                      size={28}
                      name="expand"
                      onClick={() => expandLogHandler(experiment)}
                    />
                  </td>
                </tr>
              ))}
          </tbody>
        </Table>
      </div>
    </div>
  );
};

ActivityComponent.propTypes = {};

export default ActivityComponent;
