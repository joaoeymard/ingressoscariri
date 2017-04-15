angular.module("ingressosCariri").config(function ($routeProvider) {
  $routeProvider.when("/home", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  });

  $routeProvider.otherwise({redirectTo: "/home"});
});
