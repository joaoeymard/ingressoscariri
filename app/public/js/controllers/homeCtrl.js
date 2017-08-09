angular.module("ingressosCariri").controller('homeCtrl', funcCtrl)

function funcCtrl($scope, $rootScope, Evento){
  $rootScope.pos_titulo = 'Ingressos Cariri'
  $scope.wait = true;

  Evento.get_all().then(function success (res){
    $scope.wait = false
    $scope.eventos = res.data
  },function error (res){
    $scope.wait = false
    $scope.alerts = [{
      'titulo': 'Ops!',
      'msg': 'Alguma coisa deu errado, por favor recarrege a página em alguns minutos.',
      'tipo': 'alert-warning'
    }]
    console.log("error: "+res.data);
  })
}
