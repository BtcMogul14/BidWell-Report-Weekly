{
  type: 'bar',
  data: {
  labels: ['5/21', '5/22', '5/23', '5/24', '5/25', '5/26', '5/27'],
    datasets: [{
    yAxisID: 'B',
    type: 'line',
    fill: false,
    backgroundColor: '#ff000022',
    borderColor: 'red',
    data: [383, 368, 339, 334, 332, 333, 325]
    },
    {
      yAxisID: 'A',
      label: 'Filled',
    data: [16, 15, 19, 20, 16, 13, 12],
      fill: false,
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
          labelString: 'Date'
        }
      }],
      yAxes: [{
        id: 'A',
        type: 'linear',
        position: 'right',
        ticks: {
          suggestedMin: 10,
          suggestedMax: 25 
        },
        scaleLabel: {
          display: true,
          labelString: 'Volume in Billions'
        }
      },
      {
        id: 'B',
        position: 'left',
        display: true,
        ticks: {
          suggestedMin: 300,
          suggestedMax: 450
        },
        scaleLabel: {
          display: true,
          labelString: 'Market Cap in Billions'
        }
      }]
    }
  }
}
