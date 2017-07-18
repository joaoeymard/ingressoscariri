angular.module("ingressosCariri").controller('carrinhoCtrl', funcCtrl)

function funcCtrl($scope,$rootScope,$cookieStore){
  $rootScope.pos_titulo = 'Ingressos Cariri'
  $scope.total_voucher = 0

  if($cookieStore.get('carrinho')){
    $scope.carrinho = $cookieStore.get('carrinho')

    $scope.carrinho.ingressos.map(function(ingresso){
      $scope.total_voucher += ingresso.valor * ingresso.quantidade
    })
  }
}
