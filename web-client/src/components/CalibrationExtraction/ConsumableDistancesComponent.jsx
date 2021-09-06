import React, { useState } from "react";
import { ButtonIcon, Icon, MlModal } from "shared-components";
import { Table, Button, Card, CardBody } from "core-components";

const ConsumableDistancesComponent = (props) => {
  const { consumableDistanceData } = props;

  const [selectedId, setSelectedId] = useState(null);
  const [showEditModal, setShowModal] = useState(false);

  const handleSelect = (id) => {
    if (id === selectedId) {
      setSelectedId(null);
      return;
    }
    setSelectedId(id);
  };

  const handleAddBtn = () => {};

  return (
    <>
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
                            // editStep(step.toJS());
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
