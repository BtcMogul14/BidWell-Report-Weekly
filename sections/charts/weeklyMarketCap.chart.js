{
  type: 'bar',
  data: {
    labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
    datasets: [{
    yAxisID: 'A',
    data: [164, 164, 164, 164, 164, 164, 164]
    },
    {
      yAxisID: 'B',
      label: 'Filled',
      backgroundColor: '#ff000022',
      borderColor: 'red',
      data: [425, 441, 453, 464, 466, 435, 446],
      fill: false,
      type: 'line'
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
        id: 'A',
        type: 'linear',
        position: 'right',
      },
      {
        id: 'B',
        position: 'left',
        display: true,
        ticks: {
          suggestedMin: 400,
          suggestedMax: 450
        },
        scaleLabel: {
          display: true,
          labelString: 'Value'
        }
      }]
    }
  }
}
