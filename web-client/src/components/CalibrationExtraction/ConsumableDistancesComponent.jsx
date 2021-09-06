import React, { useState } from "react";
import { useFormik } from "formik";

import { ButtonIcon, Icon } from "shared-components";
import { Table, Button, Card, CardBody } from "core-components";
import EditConsumableModal from "./EditConsumableModal";
import { consumableFormikInitialState } from "./helpers";

const ConsumableDistancesComponent = (props) => {
  const { addNewConsumableDistance, consumableDistanceData } = props;

  const formik = useFormik({
    initialValues: consumableFormikInitialState,
    enableReinitialize: true,
  });

  const [selectedId, setSelectedId] = useState(null);
  const [showEditModal, setShowModal] = useState(false);
  const [isUpdate, setIsUpdate] = useState(false);

  const handleSelect = (id) => {
    if (id === selectedId) {
      setSelectedId(null);
      return;
    }
    setSelectedId(id);
  };

  const handleAddBtn = () => {
    formik.resetForm();
    setShowModal(true);
    setIsUpdate(false);
  };

  const handleEditBtn = ({ id, name, description, distance }) => {
    formik.setFieldValue(`id.value`, id);
    formik.setFieldValue(`name.value`, name);
    formik.setFieldValue(`description.value`, description);
    formik.setFieldValue(`distance.value`, distance);

    setShowModal(true);
    setIsUpdate(true);
  };

  const handleModalBtn = (data) => {
    const { id, name, description, distance } = data;

    const requestBody = {
      id: parseInt(id.value),
      name: name.value,
      description: description.value,
      distance: parseFloat(distance.value),
    };
    addNewConsumableDistance(requestBody, isUpdate);
    setShowModal(false);
  };

  const handleCrossBtn = () => {
    setShowModal(false);
  };

  return (
    <>
      {showEditModal && (
        <EditConsumableModal
          isUpdate={isUpdate}
          isOpen={showEditModal}
          handleModalBtn={handleModalBtn}
          handleCrossBtn={handleCrossBtn}
          formik={formik}
        />
      )}
      <Card>
        <CardBody>
          <div className="table-consumable-wrapper -hold">
            <Table className="table-consumable" size="sm" striped>
              <colgroup>
                <col width="80px" />
                <col width="200px" />
                <col width="350px" />
                <col width="80px" />
                <col />
              </colgroup>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Discription</th>
                  <th>Distance</th>

                  <th>
                    <Button
                      color="primary"
                      icon
                      className="ml-auto"
                      onClick={handleAddBtn}
                    >
                      <Icon size={40} name="plus-2" />
                    </Button>
                  </th>
                </tr>
              </thead>
              <tbody>
                {consumableDistanceData.map((distanceObj, index) => {
                  const { id, name, description, distance } = distanceObj;
                  const classes = selectedId === id ? "active" : "";
                  return (
                    <tr
                      className={classes}
                      key={id}
                      onClick={() => handleSelect(id)}
                    >
                      <td>{id}</td>
                      <td>{name}</td>
                      <td>{description}</td>
                      <td>{distance}</td>

                      <td className="td-actions">
                        <ButtonIcon
                          size={16}
                          name="pencil"
                          onClick={() => {
                            handleEditBtn({ id, name, description, distance });
                          }}
                        />
                        {/* <ButtonIcon
                    size={16}
                    name="trash"
                    onClick={() => {
                      deleteStep(stepId);
                    }}
                  /> */}
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </Table>
          </div>
        </CardBody>
      </Card>
    </>
  );
};

export default ConsumableDistancesComponent;
