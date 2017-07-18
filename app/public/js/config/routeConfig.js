angular.module("ingressosCariri").config(function ($routeProvider, $locationProvider) {
  $routeProvider

  .when("/", {
    templateUrl: "view/home.html",
    controller: "homeCtrl"
  })
  .when("/evento/:titulo", {
    templateUrl: "view/evento.html",
    controller: "eventoCtrl"
  })
  .when("/ajuda", {
    templateUrl: "view/ajuda.html"
  })
  .when("/carrinho", {
    templateUrl: "view/carrinho.html",
    controller: "carrinhoCtrl"
  })

  .when("/erro", {
    templateUrl: "view/erro.html"
  })
  .otherwise({redirectTo: "/erro"})

  $locationProvider.hashPrefix('')
});
