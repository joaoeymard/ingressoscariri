angular.module("ingressosCariri").config(function ($routeProvider) {
  $routeProvider.when("/home", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  });
  $routeProvider.when("/login", {
    templateUrl: "view/login.html",
    controller: "loginCtrl"
  });

  $routeProvider.otherwise({redirectTo: "/home"});
});
