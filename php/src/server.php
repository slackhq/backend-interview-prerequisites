<?php

use BackendInterviewPrerequisites\EchoServer;

function init_and_listen() {
    register_autoloader();
    $server = new EchoServer();
    $server->listen();
}

function register_autoloader() {
    spl_autoload_register(function ($class) {
        $file = str_replace('\\', \DIRECTORY_SEPARATOR, str_replace('BackendInterviewPrerequisites\\', '', $class)).'.php';
        if (file_exists(__DIR__.'/'.$file)) {
            require(__DIR__.'/'.$file);
            return true;
        }
        return false;
    });
}

init_and_listen();