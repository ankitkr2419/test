import html2canvas from "html2canvas";
import { jsPDF } from "jspdf";

/**
 * Using this method we can export dynamic contents into pdf
 * This is used for one page
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

//multiple pages
export const saveToPdf2 = () => {
  let filename = `chart.pdf`;

  //access html element we want to export contents from
  let page1 = document.getElementById("page-1");
  let page2 = document.getElementById("page-2");
  let page3 = document.getElementById("page-3");

  //convert html element into canvas
  html2canvas(page1).then((canvas1) => {
    html2canvas(page2).then((canvas2) => {
      html2canvas(page3).then((canvas3) => {
        const img1 = canvas1.toDataURL("image/png");
        const img2 = canvas2.toDataURL("image/png");
        const img3 = canvas3.toDataURL("image/png");

        //create custom pdf instance with its properties [width, height] in pt units 
        //NOTE: (1 pt = 1.3281472327365 px)
        const doc = new jsPDF("l", "pt", [1024, 700]);

        //add contents into pdf
        doc.addImage(
          img1,
          "png",
          30, //page1.offsetLeft,
          100, //page1.offsetTop,
          page1.clientWidth,
          page1.clientHeight
        );
        doc.addPage();
        doc.addImage(
          img2,
          "png",
          30, //page2.offsetLeft,//left
          100, //page2.offsetTop//top
          page2.clientWidth,
          page2.clientHeight
        );

        //TODO remove page 3 if not needed anymore
        // doc.addPage();
        //add contents into pdf
        // doc.addImage(
        //   img3,
        //   "png",
        //   30, //page2.offsetLeft,//left
        //   100, //page2.offsetTop//top
        //   page3.clientWidth,
        //   page3.clientHeight
        // );

        //save with default filename
        doc.save(filename);
      }); //page 3 finish
    }); //page 2 finish
  }); //page 1 finish
};
