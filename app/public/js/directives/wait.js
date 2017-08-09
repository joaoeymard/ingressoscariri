angular.module("ingressosCariri").directive('wait', function() {
  return{
    restrict: 'E',
    templateUrl: 'diretivas/wait.html',
    scope: {
      text: "@",
    }
  }
})
