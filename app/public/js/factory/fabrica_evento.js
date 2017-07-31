angular.module("ingressosCariri").factory("Evento", funcFactory);

function funcFactory($http, Config){
  var factoryObj = {}

  factoryObj.get_all = function(){
    return $http.get(Config.baseUrl+'/eventos')
  }
  factoryObj.get_by_titulo = function(titulo){
    return $http.get(Config.baseUrl+'/evento/'+titulo)
  }

  return factoryObj
}
