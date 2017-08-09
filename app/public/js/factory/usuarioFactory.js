angular.module("ingressosCariri").factory("Usuario", funcFactory);

function funcFactory($http, Config){
  var factoryObj = {}

  factoryObj.get_by_titulo = function(titulo){
    // return $http({
    //   method : "GET",
    //   url : Config.baseUrl+'/evento/'+titulo
    // })
  }

  factoryObj.save = function(evento){
    // return $http({
    //   method : "POST",
    //   url : Config.baseUrl+'/evento',
    //   data: JSON.stringify(evento),
    //   headers: {
    //     'Content-Type': 'application/json',
    //   }
    // })
  }

  return factoryObj
}
