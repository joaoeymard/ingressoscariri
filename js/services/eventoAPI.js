angular.module("ingressosCariri").factory("eventoAPI", function($http, config){
  var _getEventos = function () {
    return $http.get(config.baseUrl + "/eventos.php");
  };
  var _setEventos = function (evento) {
    return $http.post(config.baseUrl + "/eventos.php", evento);
  };

  return {
    getEventos: _getEventos
  };
});
