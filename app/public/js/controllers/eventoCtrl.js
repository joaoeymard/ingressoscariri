angular.module("ingressosCariri").controller('eventoCtrl', funcCtrl)

function funcCtrl($scope,$routeParams,$rootScope){
  $rootScope.pos_titulo = ' - '+$routeParams.titulo
  $scope.teste = $routeParams.titulo

  $scope.evento = {
    titulo: "Uma festa qualquer",
    img: "evento.png",
    descricao: "A Expocrato 2017, Exposição Agropecuária do Crato, começa a divulgar as primeiras informações sobre a festa. A primeira novidade da edição da Expocrato 2017 envolve a data. A direção da Associação dos Criadores de Caprinos e Ovinos da Região do Araripe resolveu adiar a exposição para a terceira semana de julho este ano. Até o momento, não há informações sobre o local, as atrações e os ingressos de uma das maiores festividades da cidade cearense. Localizado no interior do Ceará, o Crato faz parte da Microrregião do Cariri e é conhecido também como “Terra de Alencar” e “Princesa do Cariri”. Com uma população é de aproximadamente 123 963 habitantes (IBGE 2012) e distante cerca de 567 km até Fortaleza, o Crato apresenta o 9ª maior PIB do Ceará e é a considerada a 3ª cidade com o melhor desenvolvimento. Última edição Em 2016, o evento ocorreu de 10 a 17 de julho na cidade cearense. Com 26 atrações, a programação incluiu nomes de peso como Wesley Safadão, Luan Santana e a banda de reggae Planta e Raiz. Além deles, Simone e Simaria, Aviões do Forró, Marília Mendonça, Jorge e Mateus, Gabriel Diniz e Bruno e Marrone completaram a agenda de shows.",
    cidade: "Juazeiro do Norte",
    estado: "CE",
    local: "Parque da Cidade",
    taxa: 0.15,
    periodo: [{
      data: "2017-07-07T00:00",
      hora: "22:00",
      atracao: "Jorge e Matheus, Safadão",
      lote: 1,
      categoria: [{
        nome: "VIP - Inteira",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "VIP - Meia",
        valor: 30.00,
        quantidade: 50
      }],
    }],
  }
}
