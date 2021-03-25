import React from "react";
import { Spinner } from "reactstrap";

const Loader = () => {
  return (
    <div className="position-absolute w-100 h-100 py-5 d-flex justify-content-center overlay">
      <Spinner color="info" size="md" />
    </div>
  );
};

export default Loader;
