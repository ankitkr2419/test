import React from "react";

import { Row, Col } from "core-components";
import { Text } from "shared-components";
import { LABWARE_NAME } from "appConstants";
import PreviewImage from "./PreviewImage";

const Preview = (props) => {
  const { formik } = props;
  const recipeData = formik.values;

  const getSubHead = (key) => {
    const nestedKeys = Object.keys(recipeData[key].processDetails);
    const LEN = nestedKeys.length;
    const previewInfoSubHead = [];

    nestedKeys.forEach((nestedKey) => {
      recipeData[key].processDetails[nestedKey].id &&
        previewInfoSubHead.push(
          <Text>
            {LEN > 1 && (
              <Text Tag="span" className="font-weight-bold">
                {LABWARE_NAME[nestedKey]}
                {key !== "tipPiercing" && ": "}
              </Text>
            )}
            <Text Tag="span" className={LEN === 1 ? "font-weight-bold" : ""}>
              {recipeData[key].processDetails[nestedKey].label}
            </Text>
          </Text>
        );
    });
    return previewInfoSubHead;
  };

  return (
    <>
      <div className="w-100 h-100 preview-box">
        {/* Label */}
        <Row>
          <Col
            md={12}
            className="d-flex align-items-center font-weight-bold text-center top-heading"
          >
            Preview
          </Col>
        </Row>

        <div className="d-flex justify-content-between">
          <div className="labware-selection-info w-100">
            {/* Secondary Label : Refer UI for more clarification */}
            <Text className="setting-info font-weight-bold selected-positions">
              Selected Positions
            </Text>

            {/* Content */}
            <ul className="list-unstyled">
              {Object.keys(recipeData).map((key, index) => {
                return (
                  recipeData[key].isTicked && (
                    <li className="d-flex justify-content-between">
                      <Text className="d-flex w-25 font-weight-bold">
                        {LABWARE_NAME[key]} :{" "}
                      </Text>
                      <div className="w-75">
                        <div className="ml-2 setting-value">
                          <Text>{getSubHead(key)}</Text>
                        </div>
                      </div>
                    </li>
                  )
                );
              })}
            </ul>
          </div>

          {/* SideImage */}
          <PreviewImage formik={formik} />
        </div>
      </div>
    </>
  );
};

export default React.memo(Preview);
