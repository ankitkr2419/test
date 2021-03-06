import React, { useState } from "react";
import { useFormik } from "formik";

import { ButtonIcon, Icon, Text } from "shared-components";
import { Table, Button, Card, CardBody } from "core-components";
import EditConsumableModal from "./EditConsumableModal";
import { consumableFormikInitialState } from "./helpers";

const ConsumablesAndCalibrationsComponent = (props) => {
  const { addNewData, data, heading, isReadOnly } = props;

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
    addNewData(requestBody, isUpdate);
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
      <Card className="mb-4" style={{ height: 370 }}>
        <CardBody>
          <Text
            Tag="h4"
            size={24}
            className="text-center text-gray text-bold mt-3 mb-4"
          >
            {heading}
          </Text>
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
                  <th>Description</th>
                  <th>Distance</th>

                  {!isReadOnly && (
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
                  )}
                </tr>
              </thead>
              <tbody>
                {data.map((distanceObj, index) => {
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

export default ConsumablesAndCalibrationsComponent;
