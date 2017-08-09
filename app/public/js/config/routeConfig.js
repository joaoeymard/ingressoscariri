angular.module("ingressosCariri").config(function ($routeProvider, $httpProvider, $locationProvider) {
  $routeProvider

  .when("/", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  })
  .when("/evento/:titulo", {
    templateUrl: "view/evento.html",
    controller: "eventoCtrl"
  })
  .when("/carrinho", {
    templateUrl: "view/carrinho.html",
    controller: "carrinhoCtrl"
  })
  .when("/ajuda", {
    templateUrl: "view/ajuda.html"
  })

  .when("/login", {
    templateUrl: "view/login.html",
    controller: "loginCtrl"
  })

  .when("/erro", {
    templateUrl: "view/erro.html"
  })
  .otherwise({redirectTo: "/erro"})

  $locationProvider.hashPrefix('')
});
