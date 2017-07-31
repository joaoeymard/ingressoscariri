angular.module("ingressosCariri").controller('homeCtrl', funcCtrl)

function funcCtrl($scope, $rootScope, Evento){
  $rootScope.pos_titulo = 'Ingressos Cariri'

  Evento.get_all().then(function success (res){
    $scope.eventos = res.data
  },function error(res){
    console.log("error: "+res.data);
  })
}
