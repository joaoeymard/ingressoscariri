angular.module("ingressosCariri").factory("Evento", funcFactory);

function funcFactory($http, Config){
  var listEventos
  var factoryObj = {}

  function factoryEventos_all (){
    return $http({
      method : "GET",
      url : Config.baseUrl+'/evento'
    })
  }

  factoryObj.get_all = function(){
    if(listEventos){
      return listEventos
    }else{
      listEventos = factoryEventos_all()
      return listEventos
    }
  }
  factoryObj.get_by_titulo = function(titulo){
    return $http({
      method : "GET",
      url : Config.baseUrl+'/evento/'+titulo
    })
  }

  factoryObj.save = function(evento){
    return $http({
      method : "POST",
      url : Config.baseUrl+'/evento',
      data: JSON.stringify(evento),
      headers: {
        'Content-Type': 'application/json',
      }
    })
  }

  return factoryObj
}
