<?php

function conexao(){
    $db_host = "mysql796.umbler.com";
    $db_nome = "bd_ingressos";
    $db_usuario = "admin_ingressos";
    $db_senha = "28(]#-aFqBiM";

    try {
        $db = new PDO("mysql:host=$db_host;dbname=$db_nome", $db_usuario, $db_senha);
        $db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        $db->exec('SET NAMES utf8');
        return $db;
    } catch (PDOException $e) {
        echo 'ERROR: ' . $e->getMessage();
    }
}
