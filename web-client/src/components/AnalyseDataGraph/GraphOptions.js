export const options = {
  legend: {
    display: false,
  },
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Cycles",
          fontSize: 15,
          fontStyle: "bold",
          padding: 5,
        },
        offset: true,
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "F-value",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
        },
      },
    ],
  },
};
