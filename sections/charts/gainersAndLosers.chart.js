{
  type: 'line',
  data: {
    labels: ['Test', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
    datasets: [{
    data: [164, 164, 164, 164, 164, 164, 164],
    filled: false
    },
    {
      data: [425, 441, 453, 464, 466, 435, 446],
      filled: false
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
      display: false
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
      }
      }]
    }
  }
}
