import React, { useState } from "react";
import moment from "moment";

import { Table } from "core-components";
import { ButtonIcon } from "shared-components";
import ActivityData from "./ActivityData.json";
import SearchBar from "./SearchBar";
import MlModal from "shared-components/MlModal";
import { MODAL_MESSAGE, MODAL_BTN } from "appConstants";

import "./activity.scss";

const headers = ActivityData.headers;
// const experiments = ActivityData.experiments;//TODO remove if not needed

const ActivityComponent = (props) => {
  let {
    experiments,
    searchText,
    onSearchTextChanged,
    expandLogHandler,
  } = props;

  const [selectedActivity, setSelectedActivity] = useState(null);
  const [showDeleteActivityModal, setShowDeleteActivityModal] = useState(false);

  const deleteActivityClickHandler = (e) => {
    e.stopPropagation();
    toggleDeleteActivityModal();
  };

  const toggleDeleteActivityModal = () => {
    setShowDeleteActivityModal(!showDeleteActivityModal);
  };

  const onConfirmedDeleteActivity = () => {
    const activityIdToDelete = selectedActivity?.id;

    //TODO remove console
    console.log("activity Id confirmed to delete: ", activityIdToDelete);

    toggleDeleteActivityModal();

    //TODO: API call here
  };

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
      {/**Delete activity confirmation modal */}
      {showDeleteActivityModal && (
        <MlModal
          isOpen={showDeleteActivityModal}
          textHead={""}
          textBody={MODAL_MESSAGE.deleteActivityConfirmation}
          handleSuccessBtn={onConfirmedDeleteActivity}
          handleCrossBtn={toggleDeleteActivityModal}
          successBtn={MODAL_BTN.yes}
          failureBtn={MODAL_BTN.no}
        />
      )}
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
                  key={experiment.id}
                  onClick={() => toggleSelectedActivity(experiment)}
                >
                  {/**TODO remove comments once activity log finalized */}
                  <td>
                    {i + 1}
                    {/*experiment.id*/}
                  </td>
                  <td>{experiment.template_name /*{experiment.template}*/}</td>
                  <td>
                    {
                      experiment.created_at &&
                        moment(experiment.created_at).format(
                          "DD/MM/YYYY"
                        ) /*date*/
                    }
                  </td>
                  <td>
                    {experiment.start_time &&
                      moment(experiment.start_time).format("HH:MM A")}
                  </td>
                  <td>
                    {experiment.end_time &&
                      moment(experiment.end_time).format("HH:MM A")}
                  </td>
                  <td>{experiment.well_count /*no_of_wells*/}</td>
                  <td>{experiment.repeat_cycle /*repeat_cycles*/}</td>
                  <td
                    className={
                      experiment.result === "Error"
                        ? "text-danger text-capitalize"
                        : "text-capitalize"
                    }
                  >
                    {experiment.result ? experiment.result : "N/A"}
                  </td>
                  <td className="td-actions">
                    {/* <ButtonIcon size={28} name="expand" />
                    <ButtonIcon
                      size={28}
                      name="trash"
                      onClick={deleteActivityClickHandler}
                    /> */}
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
