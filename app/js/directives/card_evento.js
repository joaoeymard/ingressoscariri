angular.module("ingressosCariri").directive('cardEvento', function() {
  return{
    restrict: 'E',
    scope: {
      evento: "=",
    },
    templateUrl: 'diretivas/card_evento.html'
  }
})
