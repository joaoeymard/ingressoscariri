angular.module("ingressosCariri").factory("eventoAPI", function($http, config){
  var _getEventos = function () {
    return $http.get(config.baseUrl + "/getEventos.php");
  };
  var _getEvento = function (id) {
    return $http.get(config.baseUrl + "/getEventos.php?id=" + id);
  };
  var _setEventos = function (evento) {
    return $http.post(config.baseUrl + "/setEventos.php", evento);
  };

  return {
    getEventos: _getEventos,
    getEvento: _getEvento
  };
});
