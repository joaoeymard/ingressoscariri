angular.module("ingressosCariri").controller('homeCtrl', funcCtrl)

function funcCtrl($scope, $rootScope){
  $rootScope.pos_titulo = 'Ingressos Cariri'
  $scope.eventos = [
    {
      titulo: "Uma festa qualquer",
      img: "evento.png",
      data: "02/07/2017",
      hora: "22:00",
      cidade: "Juazeiro do Norte",
      estado: "CE",
      local: "Parque da Cidade"
    },{
      titulo: "Uma festa qualquer",
      img: "evento.png",
      data: "02/07/2017",
      hora: "22:00",
      cidade: "Juazeiro do Norte",
      estado: "CE",
      local: "Parque da Cidade"
    },{
      titulo: "Uma festa qualquer",
      img: "evento.png",
      data: "02/07/2017",
      hora: "22:00",
      cidade: "Juazeiro do Norte",
      estado: "CE",
      local: "Parque da Cidade"
    },{
      titulo: "Uma festa qualquer",
      img: "evento.png",
      data: "02/07/2017",
      hora: "22:00",
      cidade: "Juazeiro do Norte",
      estado: "CE",
      local: "Parque da Cidade"
    }
  ]
}
