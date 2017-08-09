angular.module("ingressosCariri").filter('cpf', function(){
  return function(cpf){
    if(cpf){
      cpf = cpf.replace(/\D/g, '')
      cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2")
      cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2")
      cpf = cpf.replace(/(\d{3})(\d{1,2})$/, "$1-$2")
    }
    return cpf
  };
});
