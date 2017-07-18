angular.module("ingressosCariri").controller('eventoCtrl', funcCtrl)

function funcCtrl($scope,$routeParams,$rootScope,$location,$anchorScroll,$cookieStore){
  $rootScope.pos_titulo = $routeParams.titulo + ' - Ingressos Cariri'
  $scope.ingressoShow = 0

  $scope.mudouIngresso = function(id){
    $scope.ingressoShow = id
  }

  $scope.evento = {
    titulo: "Uma festa qualquer",
    img: "evento.png",
    descricao: "A Expocrato 2017, Exposição Agropecuária do Crato, começa a divulgar as primeiras informações sobre a festa. A primeira novidade da edição da Expocrato 2017 envolve a data. A direção da Associação dos Criadores de Caprinos e Ovinos da Região do Araripe resolveu adiar a exposição para a terceira semana de julho este ano. Até o momento, não há informações sobre o local, as atrações e os ingressos de uma das maiores festividades da cidade cearense. Localizado no interior do Ceará, o Crato faz parte da Microrregião do Cariri e é conhecido também como “Terra de Alencar” e “Princesa do Cariri”. Com uma população é de aproximadamente 123 963 habitantes (IBGE 2012) e distante cerca de 567 km até Fortaleza, o Crato apresenta o 9ª maior PIB do Ceará e é a considerada a 3ª cidade com o melhor desenvolvimento. Última edição Em 2016, o evento ocorreu de 10 a 17 de julho na cidade cearense. Com 26 atrações, a programação incluiu nomes de peso como Wesley Safadão, Luan Santana e a banda de reggae Planta e Raiz. Além deles, Simone e Simaria, Aviões do Forró, Marília Mendonça, Jorge e Mateus, Gabriel Diniz e Bruno e Marrone completaram a agenda de shows.",
    cidade: "Juazeiro do Norte",
    mapa: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3957.40824117053!2d-39.304664085679555!3d-7.307949594725646!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x7a177cb1dc3881f%3A0xc2fdccdcc5e44682!2sR.+Pero+Coelho%2C+196+-+Centro%2C+Barbalha+-+CE%2C+63180-000!5e0!3m2!1spt-BR!2sbr!4v1499230950761",
    estado: "CE",
    local: "Parque da Cidade",
    taxa: 0.08,
    periodo: [{
      data: "2017-07-07T00:00",
      hora: "22:00",
      atracao: "Safadão",
      lote: 1,
      categoria: [{
        nome: "VIP - Inteira",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "VIP - Meia",
        valor: 30.00,
        quantidade: 50
      },{
        nome: "Pista - Inteira",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "Pista - Meia",
        valor: 30.00,
        quantidade: 50
      }],
    },{
      data: "2017-07-09T00:00",
      hora: "22:00",
      atracao: "Aviões do forró",
      lote: 1,
      categoria: [{
        nome: "VIP",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "Pista",
        valor: 30.00,
        quantidade: 50
      }],
    },{
      data: "2017-07-10T00:00",
      hora: "22:00",
      atracao: "Jorge e Matheus",
      lote: 1,
      categoria: [{
        nome: "VIP",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "Pista",
        valor: 30.00,
        quantidade: 50
      }],
    }]
  }

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
