angular.module("ingressosCariri").config(function ($routeProvider, $locationProvider) {

  $routeProvider.when("/", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  })
  .when("/home", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  })
  .when("/evento/:id", {
    templateUrl: "view/evento.html",
    controller: "eventoCtrl"
  })
  .when("/login", {
    templateUrl: "view/login.html",
    controller: "loginCtrl"
  })
  .otherwise({redirectTo: "/"});

  $locationProvider.hashPrefix('');
});
