angular.module("ingressosCariri").controller('eventoCtrl', funcCtrl)

function funcCtrl($scope,$routeParams,$rootScope,$location,$anchorScroll,$cookieStore,Evento){
  $rootScope.pos_titulo = $routeParams.titulo + ' - Ingressos Cariri'
  $scope.ingressoShow = 0

  $scope.mudouIngresso = function(id){
    $scope.ingressoShow = id
  }

  Evento.get_by_titulo($routeParams.titulo).then(function success (res){
    $scope.evento = res.data
  },function error(res){
    console.log("error: "+res.data);
  })

  $scope.comprar_ingressos = function () {
    voucher = {
      'titulo': $scope.evento.titulo,
      'taxa': $scope.evento.taxa,
      'ingressos': []
    }

    $scope.evento.periodo.map(function(periodo){
      periodo.categoria.map(function(categoria){
        if(categoria.quantidade > 0){
          pedido = {
            'data': periodo.data,
            'atracao': periodo.atracao,
            'nome': categoria.nome,
            'valor': categoria.valor,
            'quantidade': categoria.quantidade
          }
          voucher.ingressos.push(pedido)
        }
      })
    })

    $cookieStore.put('carrinho', voucher)
    $location.path('carrinho')
  }
}
