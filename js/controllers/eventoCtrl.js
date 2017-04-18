angular.module("ingressosCariri").controller("eventoCtrl", function($scope, eventoAPI, $routeParams, $window){
  $scope.evento = {};
  $scope.outerEventos = {};
  $scope.exibe_anuncio = true;

  eventoAPI.getEvento($routeParams.id).then(function(response){
    $scope.evento = response.data[0];
  }, function(response){
    console.log("Aconteceu um problema, error:" + response.data);
  });

  eventoAPI.getEventos().then(function(response){
    $scope.outerEventos = response.data;
  }, function(response){
    console.log("Aconteceu um problema, error:" + response.data);
  });

  $window.onload = function (){
    alert("Iniciou");
  }
}).filter('trustAsResourceUrl', ['$sce', function($sce) {
  return function(val) {
    return $sce.trustAsResourceUrl(val);
  };
}]);
