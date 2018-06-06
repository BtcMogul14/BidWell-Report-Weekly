{
  type: 'line',
  data: {
  labels: ['5/28', '5/29', '5/30', '5/31', '6/1', '6/2', '6/3'],
    datasets: [{
    data: [5.40, 7.10, 25.85, 36.65, 34.09, 51.42, 57.39],
    fill: false,
    label: 'Huobi Token',
    borderColor: "#483d8b",
    },
    {
    data: [20.19, 78.64, 51.33, 62.46, 56.67, 66.65, 57.25],
    fill: false,
    label: 'Theta Token',
    borderColor: "rgba(247,70,74,1)",
    },
    {
    data: [5.72, 5.72, 19.34, 19.76, 23.47, 22.54, 19.93],
    fill: false,
    label: 'Binance Coin',
    borderColor: "rgba(70,191,189,1)",
    },
    {
    data: [11.70, 0.37, -5.32, -9.10, -10.04, -9.72, -5.83],
    fill: false,
    label: 'CyberMiles',
    borderColor: "rgba(77,83,96,1)",
    },
    {
    data: [7.92, -0.54, -7.34, -7.70, -6.23, -2.99, -6.06],
    fill: false,
    label: 'Mixin',
    borderColor: "#800080",
    },
    {
    data: [6.62, 3.62, 2.48, -2.92, -1.59, 0.37, -6.19],
      fill: false,
    label: 'Polymath',
    borderColor: "rgba(253,180,92,1)",
    }]
  },
  options: {
    elements: {
      point: {
        radius: 0
      }
    },
    devicePixelRatio: 2,
    responsive: true,
    animation: {
      duration: 0
    },
    legend: {
      display: true
    },
    scales: {
      xAxes: [{
        display: true,
        scaleLabel: {
          display: true,
          labelString: 'Day'
        }
      }],
      yAxes: [{
        type: 'linear',
        position: 'right',
      }]
    }
  }
}
