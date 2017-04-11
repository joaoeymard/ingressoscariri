angular.module("ingressosCariri").controller("homeCtrl", function($scope, eventoAPI){
  $scope.app = {
    'titulo': 'Ingressos Cariri - Portal de eventos do cariri'
  };
  $scope.evento = {}

  eventoAPI.getEventos().then(function(response){
    $scope.eventos = response.data;
  }, function(response){
    console.log("Aconteceu um problema, error:" + response.data);
  });
});
