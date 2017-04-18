<?php

require_once "conexao.php";

$bd = conexao();

$id = $_GET["id"];

if (isset($id)){
  $sql = $bd->prepare("SELECT * FROM `eventos` WHERE id = ? LIMIT 1");
  $sql->bindValue(1, $id);
  $sql->execute();
}else{
  $sql = $bd->prepare("SELECT * FROM `eventos`");
  $sql->execute();
}

print json_encode($sql->fetchAll());
