angular.module("ingressosCariri").controller('homeCtrl', funcCtrl)

function funcCtrl($scope){
  $scope.evento = {
    titulo: "Uma festa qualquer",
    img: "evento.png",
    data: "02/07/2017",
    hora: "22:00",
    cidade: "Juazeiro do Norte",
    estado: "CE",
    local: "Parque da Cidade"
  }
}
