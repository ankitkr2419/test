import React, { useEffect, useCallback } from "react";
import { useHistory } from "react-router";

import {
  Button,
  Form,
  FormGroup,
  Input,
  Label,
  Card,
  CardBody,
  Row,
  Col,
} from "core-components";
import { Icon, Text } from "shared-components";

import {
  isValueValid,
  formikInitialState,
  validateAllFields,
  getRequestBody,
} from "./helper";

import { HeadingTitle } from "components/CalibrationExtraction/HeadingTitle";
import CommonFieldsComponent from "components/CalibrationExtraction/CommonFieldsComponent";

const CalibrationComponent = (props) => {
  let { handleOnChange, handleSaveDetailsBtn, formik, isAdmin } = props;
  const history = useHistory();

  const handleBack = () => {
    history.goBack();
  };

  return (
    <div className="calibration-content px-5">
      <div className="d-flex align-items-center">
        {isAdmin && (
          <div style={{ cursor: "pointer" }} onClick={handleBack}>
            <Icon name="angle-left" size={32} className="text-white" />
            <HeadingTitle
              Tag="h5"
              className="text-white font-weight-bold ml-3 mb-0"
            >
              Go back to template page
            </HeadingTitle>
          </div>
        )}
      </div>
      <Card default className="my-5">
        <CardBody className="px-5 py-4">
          {/* Common fields - name, email, room temperature */}
          <CommonFieldsComponent
            formik={formik}
            handleSaveDetailsBtn={handleSaveDetailsBtn}
          />
        </CardBody>
      </Card>
    </div>
  );
};

export default React.memo(CalibrationComponent);

//old code for references
//TODO remove this if not needed for reference

// import React, { useEffect, useCallback } from "react";
// import { useFormik } from "formik";

// import {
//   Button,
//   Form,
//   FormGroup,
//   Input,
//   Label,
//   Card,
//   CardBody,
//   Row,
//   Col,
// } from "core-components";
// import { Text } from "shared-components";

// import {
//   isValueValid,
//   formikInitialState,
//   validateAllFields,
//   getRequestBody,
// } from "./helper";

// const CalibrationComponent = (props) => {
//   let { configs, saveButtonClickHandler } = props;

//   const formik = useFormik({
//     initialValues: formikInitialState,
//     enableReinitialize: true,
//   });

//   //store new data in local state
//   useEffect(() => {
//     if (configs) {
//       Object.keys(formik.values).map((element) => {
//         const { apiKey, name } = formik.values[element];

//         const newValue = configs[apiKey] ? configs[apiKey] : "";
//         const isValid = isValueValid(name, newValue);

//         // set formik fields
//         formik.setFieldValue(`${name}.isInvalid`, !isValid);
//         formik.setFieldValue(`${name}.value`, newValue);
//       });
//     }
//   }, [configs]);

//   //validations and api call
//   const onSubmit = (e) => {
//     e.preventDefault();

//     if (validateAllFields(formik.values) === true) {
//       const requestBody = getRequestBody(formik.values);
//       saveButtonClickHandler(requestBody);
//     }
//   };

//   const handleBlurChange = useCallback((name, value) => {
//     const isValid = isValueValid(name, value);
//     formik.setFieldValue(`${name}.isInvalid`, !isValid);
//   }, []);

//   const handleOnChange = (event, name) => {
//     formik.setFieldValue(`${name}.value`, event.target.value);
//   };

//   const handleOnFocus = (name) => {
//     formik.setFieldValue(`${name}.isInvalid`, false);
//   };

//   return (
//     <div className="calibration-content px-5">
//       <Card default className="my-5">
//         <CardBody className="px-5 py-4">
//           <Form onSubmit={onSubmit}>
//             <Row>
//               {Object.keys(formik.values).map((key, index) => {
//                 const element = formik.values[key];
//                 const {
//                   type,
//                   name,
//                   label,
//                   min,
//                   max,
//                   value,
//                   isInvalid,
//                   isInvalidMsg,
//                 } = element;

//                 return (
//                   <Col md={6}>
//                     <FormGroup>
//                       <Label for="username">{label}</Label>
//                       <Input
//                         type={type}
//                         name={name}
//                         id={name}
//                         placeholder={
//                           type === "number" ? `${min} - ${max}` : "Type here"
//                         }
//                         value={value}
//                         onChange={(event) => handleOnChange(event, name)}
//                         onBlur={(event) =>
//                           handleBlurChange(name, event.target.value)
//                         }
//                         onFocus={() => handleOnFocus(name)}
//                       />
//                       {(isInvalid || value == null) && (
//                         <div className="flex-70">
//                           <Text Tag="p" size={14} className="text-danger">
//                             {`${isInvalidMsg}`}
//                           </Text>
//                         </div>
//                       )}
//                     </FormGroup>
//                   </Col>
//                 );
//               })}
//             </Row>
//             <div className="text-right pt-4 pb-1 mb-3">
//               <Button
//                 color="primary"
//                 disabled={!validateAllFields(formik.values)}
//               >
//                 Save
//               </Button>
//             </div>
//           </Form>
//         </CardBody>
//       </Card>
//     </div>
//   );
// };

// export default React.memo(CalibrationComponent);
