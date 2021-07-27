import html2canvas from "html2canvas";
import { jsPDF } from "jspdf";

/**
 * Using this method we can export dynamic contents into pdf
 */
export const saveToPdf = (showTempGraph) => {
  //access html element we want to export contents from
  let input = window.document.getElementsByClassName("graph-wrapper")[0];

  //hide required elements from html element
  let toggleButton = window.document.getElementsByClassName(
    showTempGraph ? "Amplification" : "Temperature"
  )[0];
  toggleButton.style.display = "none";
  let downloadButton =
    window.document.getElementsByClassName("downloadButton")[0];
  downloadButton.style.display = "none";

  //convert html element into canvas
  html2canvas(input).then((canvas) => {
    const img = canvas.toDataURL("image/png");

    //create custom pdf instance with its properties [width, height]
    const doc = new jsPDF("l", "pt", [1024, 700]);

    //add contents into pdf
    doc.addImage(
      img,
      "png",
      input.offsetLeft,
      input.offsetTop,
      input.clientWidth,
      input.clientHeight
    );

    //save with default filename
    doc.save("chart.pdf");

    //unhide elements which got hidden
    toggleButton.style.display = "block";
    downloadButton.style.display = "block";
  });
};
