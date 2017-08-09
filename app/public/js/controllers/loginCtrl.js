angular.module("ingressosCariri").controller('loginCtrl', funcCtrl)

function funcCtrl($scope, $rootScope, $timeout){
  $rootScope.pos_titulo = 'Login - Ingressos Cariri'

  $scope.logar = function () {
    $scope.alerts = []

    // O codigo da consulta

    $scope.alerts = [{
      'msg': 'Nenhum usuario encontrado',
      'tipo': 'alert-warning'
    }]
  }
}
