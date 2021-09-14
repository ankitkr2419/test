import React from "react";
import PropTypes from "prop-types";

export default function ImageIcon({ src, alt, style, className }) {
  return <img src={src} alt={alt} style={style} className={className} />;
}

ImageIcon.propTypes = {
  src: PropTypes.string.isRequired,
  alt: PropTypes.string.isRequired,
  className: PropTypes.string,
};
