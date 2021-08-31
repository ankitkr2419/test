/**
 * returns file created from blob formated file
 * input: file in blob format | filename: for output file
 */
export const createFileFromBlob = (pdfInBlobFormat, fileName) => {
  return new File([pdfInBlobFormat], fileName, {
    lastModified: new Date().getTime(),
  });
};

/**
 * returns formData generated from blob formated file
 * input: file in blob format | temperary file name | formDataLabelForFile 
 */
export const createFormDataFromBlob = (
  pdfInBlobFormat,
  fileName,
  formDataLabelForFile
) => {
  var file = createFileFromBlob(pdfInBlobFormat, fileName);

  var data = new FormData();
  data.append(formDataLabelForFile, file);
  return data;
};
