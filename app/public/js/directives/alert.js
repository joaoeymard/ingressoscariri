angular.module("ingressosCariri").directive('alert', function() {
  return{
    restrict: 'E',
    templateUrl: 'diretivas/alert.html',
    scope: {
      alert: "=",
    },
  }
})
