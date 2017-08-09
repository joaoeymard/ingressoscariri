angular.module("ingressosCariri").directive('cardEvento', function() {
  return{
    restrict: 'E',
    templateUrl: 'diretivas/card_evento.html',
    scope: {
      evento: "=",
    }
  }
})
