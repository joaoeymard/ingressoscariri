angular.module("ingressosCariri").directive("painelevento", function(){
  return {
    templateUrl: "view/painelEvento.html",
    replace: true,
    restrict: "E",
    transclude: true
  };
}).directive("paineleventoMenu", function(){
  return {
    templateUrl: "view/painelEventomenu.html",
    restrict: "A"
  };
});