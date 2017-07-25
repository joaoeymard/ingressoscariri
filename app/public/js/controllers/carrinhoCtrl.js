angular.module("ingressosCariri").controller('carrinhoCtrl', funcCtrl)

function funcCtrl($scope,$rootScope,$cookieStore,$timeout){
  $rootScope.pos_titulo = 'Ingressos Cariri'
  $scope.total_voucher = 0
  $scope.logado = false
  $scope.login = true

  if($cookieStore.get('carrinho')){
    $scope.carrinho = $cookieStore.get('carrinho')

    $scope.carrinho.ingressos.map(function(ingresso){
      $scope.total_voucher += ingresso.valor * ingresso.quantidade
    })
  }

  $scope.getLogin = function (){
    if($scope.logado){
      $scope.logado = !$scope.logado
    }else{
      $timeout(function () {
        $scope.logado = !$scope.logado
      }, 300)
    }
  }
}
